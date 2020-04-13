package generator

import (
	"html/template"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4/plumbing"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"

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
	repository         *git.Repository
	tree               *git.Worktree
}

// OffsetsGeneratorOptions - Offsets generator option structure
type OffsetsGeneratorOptions struct {
	OffsetsDeclarationPath string
	OutputFolder           string
	TemplatePath           string
	TemplateStatic         string
	TemporaryFolder        string
	PurgeKernelRepository  bool
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
	computedOffsets, err := og.CompileAndRunPrograms(programs)
	if err != nil {
		return nil, err
	}
	return computedOffsets, nil
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

	// Generate all the possible combinations of the kernel config parameters
	combinations := [][]string{}
	for _, params := range offset.VariableKernelParameters {
		combinations = append(combinations, utils.GenerateCombinations(params)...)
	}

	// Generate the c programs from the template
	for _, combination := range combinations {
		outputPath := path.Join(tmpPath, strings.Join(combination, "-")) + ".c"
		outputFile, err := os.Create(outputPath)
		if err != nil {
			logrus.Warnf("couldn't create %s: %v\n", outputPath, err)
			return programs, err
		}
		data := model.TemplateData{
			KernelConfig: combination,
			Offset:       offset,
		}
		if err := og.cTemplate.Execute(outputFile, data); err != nil {
			logrus.Warnf("couldn't execute template for %s: %v\n", outputPath, err)
			return programs, err
		}
		prog := model.OffsetProgram{
			Offset:       offset,
			KernelConfig: combination,
			ProgramPath:  outputPath,
		}
		programs = append(programs, &prog)
	}
	return programs, nil
}

// CompileAndRunPrograms - Compiles and runs the provided programs with each kernel version
func (og *OffsetsGenerator) CompileAndRunPrograms(programs []*model.OffsetProgram) (*model.OffsetsDatabase, error) {
	var err error
	// 1) Clone or open the linux kernel repository
	if og.repository, err = og.cloneOrOpenKernelRepository(); err != nil {
		return nil, err
	}
	if og.tree, err = og.repository.Worktree(); err != nil {
		return nil, err
	}
	// 2) Compute the ranges that we need to compile the programs against
	ranges, err := og.computeVersionRanges()
	if err != nil {
		return nil, err
	}
	// 3) Loop through all the kernel versions, compile and run the programs
	for _, kvr := range ranges {
		iter := kvr.Iter()
		for {
			version := iter.NextMinor()
			if version == nil {
				break
			}
			if err := og.CheckoutKernelVersion(version); err != nil {
				logrus.Errorf("couldn't checkout branch %s: %s\n", version.ReferenceName(), err)
				continue
			}
			for _, program := range programs {
				if err := og.CompileAndRunProgram(program); err != nil {
					logrus.Errorf("couldn't compute offset for program %s and version %s: %s\n", err)
					continue
				}
			}
		}
	}
	return nil, nil
}

// cloneOrOpenKernelRepository - Clones or open the kernel repository
func (og *OffsetsGenerator) cloneOrOpenKernelRepository() (*git.Repository, error) {
	kernelPath := path.Join(og.options.TemporaryFolder, linuxKernelPath)
	var err error
	// Check if the repository exists
	if _, err = os.Stat(kernelPath); err == nil {
		// This means that the repository is already present, check if it should be purged
		if og.options.PurgeKernelRepository {
			logrus.Println("purging kernel repository ...")
			if err = os.RemoveAll(kernelPath); err != nil {
				return nil, err
			}
		} else {
			// Try to open the repository
			logrus.Println("opening kernel repository ...")
			return git.PlainOpen(kernelPath)
		}
	}
	if err = os.MkdirAll(kernelPath, os.ModePerm); err != nil {
		logrus.Warnf("couldn't generate temporary path %s: %v\n", kernelPath, err)
		return nil, err
	}
	logrus.Printf("cloning the Linux Kernel repository ...\n")
	repository, err := git.PlainClone(kernelPath, false, &git.CloneOptions{
		URL:      linuxKernelGithub,
		Progress: os.Stdout,
		Depth:    1,
	})
	if err != nil && err != git.ErrRepositoryAlreadyExists {
		return nil, err
	}
	logrus.Printf("cloning done!\n")
	return repository, nil
}

// computeVersionRanges - Computes the ranges of versions that the programs need to be compiled against
func (og *OffsetsGenerator) computeVersionRanges() ([]model.KernelVersionRange, error) {
	inputRange := og.OffsetsDeclaration.GetKernelVersionRange()
	var ranges []model.KernelVersionRange
	rangeTmp := model.KernelVersionRange{
		MinVersion: &model.KernelVersion{
			Major: inputRange.MinVersion.Major,
			Minor: inputRange.MinVersion.Minor,
		},
		MaxVersion: &model.KernelVersion{
			Major: inputRange.MinVersion.Major,
			Minor: inputRange.MinVersion.Minor,
		},
	}
	// Loop through the input kernel versions
	logrus.Printf("checking available kernel versions for input range: %s\n", inputRange)
	iter := inputRange.Iter()
	for {
		next := iter.NextMinor()
		if next == nil {
			// Append the current rangeTmp to the list of computed ranges
			ranges = append(ranges, rangeTmp)
			break
		}
		// Check if the next version exists
		if _, err := og.repository.Tag(next.String()); err != nil {
			// Append the current rangeTmp to the list of computed ranges
			ranges = append(ranges, rangeTmp)
			// check the next Major version
			next = iter.NextMajor()
			if next == nil {
				break
			}
			if _, err := og.repository.Tag(next.String()); err != nil {
				break
			}
			// Prepare the new kernel version range
			rangeTmp = model.KernelVersionRange{
				MinVersion: &model.KernelVersion{
					Major: next.Major,
					Minor: next.Minor,
				},
			}
		}
		// Extend the temporary range to include the current "next" cursor
		rangeTmp.ExtendMaxKernelVersion(*next)
	}
	logrus.Printf("available ranges: %v\n", ranges)
	return ranges, nil
}

// CheckoutKernelVersion - Checkout to the provided kernel version
func (og *OffsetsGenerator) CheckoutKernelVersion(version *model.KernelVersion) error {
	logrus.Tracef("checking out kernel version: %s\n", version)
	return og.tree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(version.ReferenceName()),
		Force:  true,
	})
}

// CompileAndRunProgram - Compile the provided program and run it to compute its offset
func (og *OffsetsGenerator) CompileAndRunProgram(program *model.OffsetProgram) error {
	// Compile program
	logrus.Tracef("compiling: %s %v\n", program.Offset.OffsetSymbol, program.KernelConfig)
	// Run program
	return nil
}
