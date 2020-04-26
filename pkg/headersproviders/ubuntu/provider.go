package ubuntu

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"sort"

	"github.com/sirupsen/logrus"

	"github.com/Gui774ume/ebpf-offsets-generator/pkg/utils"
	"github.com/pkg/errors"

	"github.com/Gui774ume/ebpf-offsets-generator/pkg/model"
)

// Provider - Header provider
type Provider struct {
	tmp string
}

// SetTmpPath - Creates a new provider and sets the temporary folder path
func (p *Provider) SetTmpPath(path string) {
	p.tmp = path
}

// HasVersion - Returns true if the current provider has headers for the provided kernel version
func (p *Provider) HasVersion(kv model.KernelVersion) bool {
	_, ok := db.Packages[kv]
	return ok
}

// PullHeaders - Pulls headers for the provided kernel version
func (p *Provider) PullHeaders(kv model.KernelVersion) ([]string, error) {
	var headers []string
	headersRoot := path.Join(p.tmp, kv.String())
	ubuntuPackage, ok := db.Packages[kv]
	if !ok {
		return headers, errors.Errorf("ubuntu provider: couldn't find kernel version %v", kv)
	}
	// Download header packages
	if ubuntuPackage.LinuxHeaderAll != "" {
		if _, err := os.Stat(path.Join(headersRoot, "all.deb")); err != nil {
			if err = p.downloadAndExtract(ubuntuPackage.LinuxHeaderAll, headersRoot, "all.deb"); err != nil {
				return headers, err
			}
		}
	}
	if ubuntuPackage.LinuxHeaderGeneric != "" {
		if _, err := os.Stat(path.Join(headersRoot, "generic.deb")); err != nil {
			if err = p.downloadAndExtract(ubuntuPackage.LinuxHeaderGeneric, headersRoot, "generic.deb"); err != nil {
				return headers, err
			}
		}
	}
	// generate headers path
	return p.generateHeaders(headersRoot)
}

func (p *Provider) downloadAndExtract(url string, outputDir string, filename string) error {
	outputPath := path.Join(outputDir, filename)
	// Download header packages
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}
	logrus.Printf("downloading %s ...\n", url)
	if err := utils.DownloadFile(url, outputPath); err != nil {
		return errors.Wrapf(err, "ubuntu provider: failed to download header %v", url)
	}
	// Extract the headers
	logrus.Printf("extracting %s ...\n", url)
	_, err := exec.Command("dpkg-deb", "-x", outputPath, outputDir).Output()
	if err != nil {
		return err
	}
	return nil
}

// KernelVersionsList - Returns the list of all available kernel versions
func (p *Provider) KernelVersionsList() []model.KernelVersion {
	keys := make([]model.KernelVersion, len(db.Packages))
	for k := range db.Packages {
		keys = append(keys, k)
	}
	sort.Sort(model.ByVersion(keys))
	return keys
}

// generateHeaders - Returns the list of headers to include from the headers root
func (p *Provider) generateHeaders(root string) ([]string, error) {
	// Find the linux-headers package suffix
	var headers []string
	var suffix string
	srcRoot := path.Join(root, "usr/src/")
	files, err := ioutil.ReadDir(srcRoot)
	if err != nil {
		return headers, err
	}
	for _, file := range files {
		if matched, err := regexp.Match("(-generic)", []byte(file.Name())); err == nil && matched {
			suffix = file.Name()
		}
	}
	if suffix == "" {
		switch len(files) {
		case 1:
			suffix = files[0].Name()
			break
		case 2:
			suffix = files[1].Name()
			break
		default:
			return headers, errors.New("couldn't find headers root")
		}
	}
	return []string{
		path.Join(srcRoot, suffix, "include"),
		path.Join(srcRoot, suffix, "include/uapi"),
		path.Join(srcRoot, suffix, "include/generated/uapi"),
		path.Join(srcRoot, suffix, "arch/x86/include"),
		path.Join(srcRoot, suffix, "arch/x86/include/uapi"),
		path.Join(srcRoot, suffix, "arch/x86/include/generated"),
	}, nil
}
