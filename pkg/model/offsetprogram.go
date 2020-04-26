package model

// OffsetProgram - This structure is used to register a program that needs to be compiled against the different kernel versions
type OffsetProgram struct {
	Offset      OffsetDeclaration
	ProgramPath string
	BinaryPath  string
}
