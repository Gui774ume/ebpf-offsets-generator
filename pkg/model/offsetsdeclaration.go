package model

import (
	"os"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

// OffsetsDeclaration - Offsets declaration structure
type OffsetsDeclaration struct {
	Version          string `yaml:"version"`
	MinKernel        string `yaml:"min_kernel"`
	MinKernelVersion *KernelVersion
	MaxKernel        string `yaml:"max_kernel"`
	MaxKernelVersion *KernelVersion
	Offsets          []OffsetDeclaration `yaml:"offsets"`
}

// GetKernelVersionRange - Returns the kernel version range of the offsets declaration
func (od OffsetsDeclaration) GetKernelVersionRange() KernelVersionRange {
	return KernelVersionRange{
		MaxVersion: od.MaxKernelVersion,
		MinVersion: od.MinKernelVersion,
	}
}

// GetKernelVersionIter - Returns an iterator to go over all the requested kernel versions
func (od OffsetsDeclaration) GetKernelVersionIter() KernelVersionIter {
	return od.GetKernelVersionRange().Iter()
}

// OffsetDeclaration - Offset declaration structure
type OffsetDeclaration struct {
	OffsetSymbol  string   `yaml:"offset_symbol"`
	StructureName string   `yaml:"structure_name"`
	MemberName    string   `yaml:"member_name"`
	IncludePaths  []string `yaml:"include_paths"`
}

// TemplateData - Template Data provided to the template to generate the c programs
type TemplateData struct {
	Offset OffsetDeclaration
}

// ReadDeclarationFile - Read the provided offsets declaration file and populates the provided OffsetsDeclaration instance
func ReadDeclarationFile(path string, od *OffsetsDeclaration) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			logrus.Errorf("couldn't close offsets declaration file: %v", err)
		}
	}()
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(od); err != nil {
		return err
	}
	// Parse kernel versions
	od.MinKernelVersion, err = NewKernelVersionFromString(od.MinKernel)
	if err != nil {
		return err
	}
	od.MaxKernelVersion, err = NewKernelVersionFromString(od.MaxKernel)
	if err != err {
		return err
	}
	if od.MinKernelVersion.After(*od.MaxKernelVersion) {
		return errors.New("MinKernel must be lower than MaxKernel")
	}
	return nil
}
