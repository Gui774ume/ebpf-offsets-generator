package generator

import (
	"bytes"
	"github.com/Gui774ume/ebpf-offsets-generator/pkg/model"
	"io"
	"os/exec"
)

// compile - This function compiles an eBPF program with the provided headers and returns the path to the elf bytecode
func compile(program *model.OffsetProgram, headers []string) error {
	// Clang and llc commands
	clangArgs := []string{
		"-D__KERNEL__",
		"-D__ASM_SYSREG_H",
		"-Wno-unused-value",
		"-Wno-pointer-sign",
		"-Wno-compare-distinct-pointer-types",
		"-Wunused",
		"-Wall",
		"-O2",
		"-emit-llvm",
		"-c",
		"-o",
		"-",
	}
	// Add headers
	clangArgs = append(clangArgs, includeList(headers)...)
	// Add program source
	clangArgs = append(clangArgs, program.ProgramPath)
	llcArgs := []string{
		"-march=bpf",
		"-filetype=obj",
		"-o",
	}
	// Add output file
	llcArgs = append(llcArgs, program.BinaryPath)
	// Pipe commands
	clang := exec.Command("clang", clangArgs...)
	llc := exec.Command("llc", llcArgs...)
	r, w := io.Pipe()
	clang.Stdout = w
	llc.Stdin = r
	var b bytes.Buffer
	llc.Stdout = &b
	if err := clang.Start(); err != nil {
		return err
	}
	if err := llc.Start(); err != nil {
		return err
	}
	if err := clang.Wait(); err != nil {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}
	return llc.Wait()
}

func includeList(headers []string) []string {
	var out []string
	for _, header := range headers {
		out = append(out, "-I"+header)
	}
	return out
}
