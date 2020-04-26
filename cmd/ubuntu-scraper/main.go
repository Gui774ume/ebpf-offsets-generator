package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"regexp"

	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"

	"github.com/Gui774ume/ebpf-offsets-generator/pkg/model"
	"github.com/Gui774ume/ebpf-offsets-generator/pkg/utils"
)

var (
	ubuntuMainlineVersionsURL       = "https://kernel.ubuntu.com/~kernel-ppa/mainline/"
	ubuntuMainlineVersionPattern    = "v[0-9]\\.[0-9]*\\.*[0-9]*\\/"
	ubuntuGenericLinuxHeaderPattern = "(linux-headers-%d\\.%d\\.%d).*(generic).*(amd64\\.deb)"
	ubuntuAllLinuxHeaderPattern     = "(linux-headers-%d\\.%d\\.%d).*(all\\.deb)"
	outputTemplate                  = `
package ubuntu

import "github.com/Gui774ume/ebpf-offsets-generator/pkg/model"

var db = model.UbuntuDatabase{
	Packages: map[model.KernelVersion]*model.UbuntuPackages{ {{ range $key, $value := .Packages }}
		model.KernelVersion{
			Major: {{ $key.Major }},
			Minor: {{ $key.Minor }},
			Patch: {{ $key.Patch }},
		}: {
			LinuxHeaderAll:     "{{ $value.LinuxHeaderAll }}",
			LinuxHeaderGeneric: "{{ $value.LinuxHeaderGeneric }}",
		}, {{ end }}
	},
}
`
)

func main() {
	// Parse command line arguments
	outputFile := flag.String("output", "./pkg/headersproviders/ubuntu/db.go", "output file")
	flag.Parse()

	// Setup logger
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		DisableLevelTruncation: true,
	})

	// Scrape all ubuntu linux header packages
	logrus.Printf("fetching linux headers recursively from %s ...\n", ubuntuMainlineVersionsURL)
	ubuntuDB, err := scrapeUbuntuLinuxHeadersPackages()
	if err != nil {
		logrus.Errorf("error while scraping ubuntu linux headers packages: %v\n", err)
		return
	}

	// Export Ubuntu database
	logrus.Printf("exporting Ubuntu database to %s ...\n", *outputFile)
	if err := exportUbuntuDatabase(ubuntuDB, *outputFile); err != nil {
		logrus.Errorf("error while exporting ubuntu database: %v", err)
		return
	}
	logrus.Println("done!")
}

func exportUbuntuDatabase(db *model.UbuntuDatabase, outputPath string) error {
	// Parse template
	t, err := template.New("main.c").Funcs(template.FuncMap{"unescape": utils.Unescape}).Parse(outputTemplate)
	if err != nil {
		return err
	}
	// Prepare output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	// Execute template
	if err := t.Execute(outputFile, *db); err != nil {
		return err
	}
	return nil
}

func scrapeUbuntuLinuxHeadersPackages() (*model.UbuntuDatabase, error) {
	db := model.NewUbuntuDatabase()
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Look for the mainline version pattern
		if matched, err := regexp.Match(ubuntuMainlineVersionPattern, []byte(link)); err == nil && matched {
			// Visit recursively to fetch the linux header packages
			c.Visit(e.Request.AbsoluteURL(link))
		}
		// Look for a linux header pattern
		re := regexp.MustCompile(ubuntuMainlineVersionPattern)
		if match := re.Find([]byte(e.Request.URL.String())); len(match) > 0 {
			// Parse kernel version
			kv, err := model.NewKernelVersionFromString(string(match))
			if err != nil {
				return
			}
			// Generate package patterns
			allPattern := fmt.Sprintf(ubuntuAllLinuxHeaderPattern, kv.Major, kv.Minor, kv.Patch)
			genericPattern := fmt.Sprintf(ubuntuGenericLinuxHeaderPattern, kv.Major, kv.Minor, kv.Patch)
			// Look for the a linux header pattern
			if matched, err := regexp.Match(allPattern, []byte(link)); err == nil && matched {
				db.Lock()
				p, ok := db.Packages[*kv]
				if ok {
					if p.LinuxHeaderAll == "" {
						p.LinuxHeaderAll = e.Request.AbsoluteURL(link)
					}
				} else {
					db.Packages[*kv] = &model.UbuntuPackages{
						LinuxHeaderAll: e.Request.AbsoluteURL(link),
					}
				}
				db.Unlock()
			}
			if matched, err := regexp.Match(genericPattern, []byte(link)); err == nil && matched {
				db.Lock()
				p, ok := db.Packages[*kv]
				if ok {
					if p.LinuxHeaderGeneric == "" {
						p.LinuxHeaderGeneric = e.Request.AbsoluteURL(link)
					}
				} else {
					db.Packages[*kv] = &model.UbuntuPackages{
						LinuxHeaderGeneric: e.Request.AbsoluteURL(link),
					}
				}
				db.Unlock()
			}
		}
	})

	// Start scraping
	if err := c.Visit(ubuntuMainlineVersionsURL); err != nil {
		return &db, err
	}
	return &db, nil
}
