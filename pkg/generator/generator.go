package generator

import (
	"html/template"
	"os"
	"path"
	"path/filepath"

	"github.com/Gui774ume/ebpf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Gui774ume/ebpf-offsets-generator/pkg/headersproviders"
	"github.com/Gui774ume/ebpf-offsets-generator/pkg/model"
	"github.com/Gui774ume/ebpf-offsets-generator/pkg/utils"
)

const (
	generatedPrograms = "generated_programs"
	linuxKernelPath   = "linux_kernel"
	linuxKernelGithub = "https://github.com/torvalds/linux"
)

// OffsetsGenerator - Offsets generator main structure
type OffsetsGenerator struct {
	OffsetsDeclaration model.OffsetsDeclaration
	options            *OffsetsGeneratorOptions
	cTemplate          *template.Template
	provider           model.HeadersProvider
}

// OffsetsGeneratorOptions - Offsets generator option structure
type OffsetsGeneratorOptions struct {
	OffsetsDeclarationPath string
	TemplatePath           string
	TemplateStatic         string
	TemporaryFolder        string
	PurgeKernelRepository  bool
	HeaderProvider         string
}

// NewOffsetsGeneratorWithOptions - Creates a new Offsets Generator with the provided options
func NewOffsetsGeneratorWithOptions(options *OffsetsGeneratorOptions) (*OffsetsGenerator, error) {
	og := OffsetsGenerator{
		options: options,
	}
	// Parse offsets declaration
	if err := model.ReadDeclarationFile(options.OffsetsDeclarationPath, &og.OffsetsDeclaration); err != nil {
		return nil, err
	}
	// Select the header provider
	var ok bool
	og.provider, ok = headersproviders.Providers[options.HeaderProvider]
	if !ok {
		return nil, errors.Errorf("unknown header provider: %s", options.HeaderProvider)
	}
	og.provider.SetTmpPath(path.Join(options.TemporaryFolder, "headers"))
	return &og, nil
}

// ComputeOffsets - Computes offsets for the provided kernel versions and kernel configurations
func (og *OffsetsGenerator) ComputeOffsets() (*model.OffsetsDatabase, error) {
	// 1) Generate the offsets programs used to compute the provided offsets.
	programs, err := og.GenerateOffsetsPrograms()
	if err != nil {
		return nil, err
	}
	// 2) Check out each requested kernel version, compile the c programs generated above with each kernel and compute offsets.
	return og.CompileAndExtractOffsets(programs)
}

// GenerateOffsetsPrograms - Generate the offsets programs used to compute the provided offsets
func (og *OffsetsGenerator) GenerateOffsetsPrograms() ([]*model.OffsetProgram, error) {
	var programs []*model.OffsetProgram
	logrus.Printf("generating offsets programs ...\n")
	// Parse c template
	var err error
	og.cTemplate = template.New("main.c").Funcs(template.FuncMap{"unescape": utils.Unescape})
	og.cTemplate, err = og.cTemplate.ParseFiles(og.options.TemplatePath)
	if err != nil {
		return programs, errors.Wrap(err, "couldn't parse c template")
	}

	// Generate output folder
	generatedProgramsPath := path.Join(og.options.TemporaryFolder, generatedPrograms)
	if err := os.MkdirAll(generatedProgramsPath, os.ModePerm); err != nil {
		return programs, errors.Wrap(err, "couldn't create temporary path")
	}

	// Loop through all the provided offsets and generate the c programs
	for _, offset := range og.OffsetsDeclaration.Offsets {
		if progs, err := og.generateOffsetPrograms(generatedProgramsPath, offset); err != nil {
			logrus.Warnf("ignoring symbol %s\n", offset.OffsetSymbol)
		} else {
			programs = append(programs, progs...)
		}
	}
	logrus.Printf("offsets programs generation done!\n")
	return programs, nil
}

func (og *OffsetsGenerator) generateOffsetPrograms(generatedProgramsPath string, offset model.OffsetDeclaration) ([]*model.OffsetProgram, error) {
	var programs []*model.OffsetProgram
	// Create folder for symbol
	tmpPath := path.Join(generatedProgramsPath, offset.OffsetSymbol)
	if err := os.MkdirAll(tmpPath, os.ModePerm); err != nil {
		logrus.Warnf("couldn't generate temporary path %s: %v\n", tmpPath, err)
		return programs, err
	}

	// Copy bpf helpers
	staticPath := path.Join(tmpPath, filepath.Base(og.options.TemplateStatic))
	if err := utils.CopyDir(og.options.TemplateStatic, staticPath); err != nil {
		return programs, err
	}

	// Generate the c programs from the template
	outputPath := path.Join(tmpPath, offset.OffsetSymbol) + ".c"
	outputFile, err := os.Create(outputPath)
	if err != nil {
		logrus.Warnf("couldn't create %s: %v\n", outputPath, err)
		return programs, err
	}
	data := model.TemplateData{
		Offset: offset,
	}
	if err := og.cTemplate.Execute(outputFile, data); err != nil {
		logrus.Warnf("couldn't execute template for %s: %v\n", outputPath, err)
		return programs, err
	}
	prog := model.OffsetProgram{
		Offset:      offset,
		ProgramPath: outputPath,
	}
	programs = append(programs, &prog)
	return programs, nil
}

// CompileAndExtractOffsets - Compiles the provided programs with the right headers and extract the generated offsets
func (og *OffsetsGenerator) CompileAndExtractOffsets(programs []*model.OffsetProgram) (*model.OffsetsDatabase, error) {
	offsetsDatabase := model.NewOffsetsDatabase(og.OffsetsDeclaration.Version, og.options.HeaderProvider)
	// Loop through all the kernel versions, compile and run the programs
	for _, kv := range og.provider.KernelVersionsList() {
		// Check if the version is in the input range
		if !og.OffsetsDeclaration.GetKernelVersionRange().Contains(kv) {
			continue
		}
		logrus.Printf("working on version %v ...\n", kv)
		// Pulls the headers for the current version
		headers, err := og.provider.PullHeaders(kv)
		if err != nil {
			logrus.Errorf("couldn't pull headers: %v", err)
			continue
		}
		for _, program := range programs {
			// Compile and extract the offset
			offset, err := og.compileAndExtractOffset(program, headers, kv)
			if err != nil {
				logrus.Errorf("couldn't compute offset for program %s and version %s: %s\n", program.ProgramPath, kv, err)
				continue
			}
			logrus.Printf("generated offset for %s and kernel version %s: %v\n", program.Offset.OffsetSymbol, kv, offset.Offset)
			// Add offset to the database
			offsetsDatabase.Insert(*offset)
		}
	}
	return offsetsDatabase, nil
}

// compileAndExtractOffset - Compiles the input program with the input headers, returns the generated offset
func (og *OffsetsGenerator) compileAndExtractOffset(program *model.OffsetProgram, headers []string, kv model.KernelVersion) (*model.GeneratedOffset, error) {
	// Generate output path
	outputRoot, outputFilename := path.Split(program.ProgramPath)
	outputRoot = path.Join(outputRoot, kv.String())
	if err := os.MkdirAll(outputRoot, os.ModePerm); err != nil {
		return nil, err
	}
	ext := path.Ext(outputFilename)
	if len(ext) > 0 {
		outputFilename = outputFilename[:len(outputFilename)-len(ext)]
	}
	outputFilename = outputFilename + ".o"
	program.BinaryPath = path.Join(outputRoot, outputFilename)

	// Compile eBPF program
	if err := compile(program, headers); err != nil {
		return nil, errors.Wrapf(err, "couldn't compile program %s with kernel version %v", program.Offset.OffsetSymbol, kv)
	}

	// Extract generated offset
	spec, err := ebpf.LoadCollectionSpec(program.BinaryPath)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't load eBPF program %s build for kernel version %v", program.Offset.OffsetSymbol, kv)
	}
	prog := spec.Programs["tracepoint/syscalls/sys_enter_execve"]
	if prog == nil {
		return nil, errors.Errorf("couldn't find eBPF program in %s", program.BinaryPath)
	}
	// The eBPF program was designed so that the first constant of the first instruction is the offset we are looking for
	if len(prog.Instructions) == 0 {
		return nil, errors.Errorf("coulnd't find the first instruction of the eBPF program in %s", program.BinaryPath)
	}
	return &model.GeneratedOffset{
		MinKernelVersion: kv,
		MaxKernelVersion: kv,
		Offset:           prog.Instructions[0].Constant,
		OffsetSymbol:     program.Offset.OffsetSymbol,
	}, nil
}
