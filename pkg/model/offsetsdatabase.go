package model

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// OffsetsDatabase - Offsets database structure
type OffsetsDatabase struct {
	Version          string                         `yaml:"version"`
	Provider         string                         `yaml:"provider"`
	GeneratedOffsets map[string]*[]*GeneratedOffset `yaml:"generated_offsets"`
}

// NewOffsetsDatabase - Returns a pointer to a new Offsets Database
func NewOffsetsDatabase(version string, provider string) *OffsetsDatabase {
	return &OffsetsDatabase{
		Version:          version,
		Provider:         provider,
		GeneratedOffsets: map[string]*[]*GeneratedOffset{},
	}
}

// DumpSymbol - Dumps the generated offsets to the terminal
func (od *OffsetsDatabase) DumpSymbol(symbol string) error {
	db, ok := od.GeneratedOffsets[symbol]
	if !ok {
		return errors.Errorf("unknown symbol: %v", symbol)
	}
	logrus.Printf("dumping offsets for %v\n", symbol)
	for _, off := range *db {
		logrus.Printf("%v -> %v : %v\n", off.MinKernelVersion, off.MaxKernelVersion, off.Offset)
	}
	return nil
}

// Insert - Inserts a newly generated offset in the database
func (od *OffsetsDatabase) Insert(offset GeneratedOffset) {
	// Select the right offsets list from the offset symbol
	offsetDB, ok := od.GeneratedOffsets[offset.OffsetSymbol]
	if !ok {
		offsetDB = &[]*GeneratedOffset{}
		od.GeneratedOffsets[offset.OffsetSymbol] = offsetDB
	}
	// Loop through all the offsets ranges and insert the new one at the right place
	inserted := false
	for _, dbOffset := range *offsetDB {
		// Check if the offsets match
		if dbOffset.Offset == offset.Offset {
			// Check if the kernel versions are contiguous
			if dbOffset.MinKernelVersion.IsNeighborOf(offset.MinKernelVersion) && dbOffset.MinKernelVersion.After(offset.MinKernelVersion) {
				dbOffset.MinKernelVersion.Patch = offset.MinKernelVersion.Patch
				dbOffset.MinKernel = dbOffset.MinKernelVersion.String()
				inserted = true
			}
			if dbOffset.MaxKernelVersion.IsNeighborOf(offset.MaxKernelVersion) && dbOffset.MaxKernelVersion.Before(offset.MaxKernelVersion) {
				dbOffset.MaxKernelVersion.Patch = offset.MaxKernelVersion.Patch
				dbOffset.MaxKernel = dbOffset.MaxKernelVersion.String()
				inserted = true
			}
		}
	}
	// If the offset wasn't inserted, then this is a new range of kernel versions
	if !inserted {
		offset.MaxKernel = offset.MaxKernelVersion.String()
		offset.MinKernel = offset.MinKernelVersion.String()
		*offsetDB = append(*offsetDB, &offset)
	}
}

// GeneratedOffset - Generated offset structure
type GeneratedOffset struct {
	OffsetSymbol     string        `yaml:"-"`
	MinKernel        string        `yaml:"min_kernel"`
	MinKernelVersion KernelVersion `yaml:"-"`
	MaxKernel        string        `yaml:"max_kernel"`
	MaxKernelVersion KernelVersion `yaml:"-"`
	Offset           int64         `yaml:"offset"`
}
