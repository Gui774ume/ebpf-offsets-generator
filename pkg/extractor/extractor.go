package extractor

import (
	"os"
	"path"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/Gui774ume/ebpf-offsets-generator/pkg/model"
)

// ExportOffsetsDatabaseToFile - Exports the provided offsets database to a file
func ExportOffsetsDatabaseToFile(db *model.OffsetsDatabase, output string) error {
	// Make all parent directories if necessary
	if err := os.MkdirAll(path.Dir(output), os.ModePerm); err != nil {
		return errors.Wrapf(err, "couldn't create directory %s", path.Dir(output))
	}
	// Create the output file
	outputFile, err := os.Create(output)
	if err != nil {
		return errors.Wrapf(err, "couldn't create file %s", output)
	}
	// Export the database
	encoder := yaml.NewEncoder(outputFile)
	return errors.Wrap(encoder.Encode(db), "couldn't encore yaml data")
}
