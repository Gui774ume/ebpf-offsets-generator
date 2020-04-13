package main

import (
	"flag"

	"github.com/Gui774ume/ebpf-offsets-generator/pkg/generator"

	"github.com/sirupsen/logrus"
)

func main() {
	// Parse command line arguments
	verbose := flag.Uint("verbose", 6, "verbose level (0 to 6)")
	offsets := flag.String("offsets", "", "offsets declaration path")
	template := flag.String("template", "./template/main.c", "c template path")
	templateStatic := flag.String("templateStatic", "./template/bpf", "template static files to copy next to the generated programs")
	tmpFolder := flag.String("temporaryFolder", "./tmp", "temporary folder")
	outputFolder := flag.String("profile", "", "output folder")
	purge := flag.Bool("purge", false, "purges the Linux Kernel repository and clones it again")
	flag.Parse()

	// Set log verbose level
	logrus.SetLevel(logrus.Level(*verbose))
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		DisableLevelTruncation: true,
	})

	// Create offsets generator
	og, err := generator.NewOffsetsGeneratorWithOptions(
		&generator.OffsetsGeneratorOptions{
			OutputFolder:           *outputFolder,
			OffsetsDeclarationPath: *offsets,
			TemplatePath:           *template,
			TemporaryFolder:        *tmpFolder,
			PurgeKernelRepository:  *purge,
			TemplateStatic:         *templateStatic,
		},
	)
	if err != nil {
		logrus.Errorf("couldn't create new offsets generator: %v", err)
		return
	}

	// Compute offsets
	_, err = og.ComputeOffsets()
	if err != nil {
		logrus.Errorf("couldn't start offsets generator: %v", err)
	}
	return
}
