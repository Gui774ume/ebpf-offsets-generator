package model

// OffsetsDatabase - Offsets database structure
type OffsetsDatabase struct {
	Version          string                     `yaml:"version"`
	GeneratedOffsets map[string]GeneratedOffset `yaml:"generated_offsets"`
}

// GeneratedOffset - Generated offset structure
type GeneratedOffset struct {
	MinKernel    string             `yaml:"min_kernel"`
	MaxKernel    string             `yaml:"max_kernel"`
	Offset       int                `yaml:"offset"`
	KernelConfig []KernelConfigDiff `yaml:"kernel_config"`
}

// KernelConfigDiff - Kernel config diff structure
type KernelConfigDiff struct {
	Config []string `yaml:"config"`
	Diff   int      `yaml:"diff"`
}
