package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// KernelVersion - Kernel version representation
type KernelVersion struct {
	Major int
	Minor int
}

func (kv KernelVersion) String() string {
	return fmt.Sprintf("v%d.%d", kv.Major, kv.Minor)
}

// ReferenceName - Returns the reference name of the tag in the kernel repository
func (kv KernelVersion) ReferenceName() string {
	return fmt.Sprintf("refs/tags/%s", kv.String())
}

// Before - Returns true if the current kernel version is before the input kernel version
func (kv KernelVersion) Before(version KernelVersion) bool {
	return kv.Int() < version.Int()
}

// After - Returns true if the current kernel version is after the input kernel version
func (kv KernelVersion) After(version KernelVersion) bool {
	return kv.Int() > version.Int()
}

// Equal - Returns true if the two kernel versions are equal
func (kv KernelVersion) Equal(version KernelVersion) bool {
	return kv.Int() == version.Int()
}

// Int - Returns an integer representation of the kernel version
func (kv KernelVersion) Int() int {
	return kv.Major*1000 + kv.Minor
}

// NextMinor - Increments the current kernel version to the next minor version
func (kv *KernelVersion) NextMinor() {
	kv.Minor += 1
}

// NextMajor - Increments the current kernel version to the next major version
func (kv *KernelVersion) NextMajor() {
	kv.Major += 1
	kv.Minor = 0
}

// NewKernelVersionFromString - Creates a new kernel version instance from the string representation of a kernel version
func NewKernelVersionFromString(kv string) (*KernelVersion, error) {
	split := strings.Split(kv, ".")
	if len(split) != 2 {
		return nil, errors.Errorf("the following kernel version doesn't follow the pattern x.y: %v", kv)
	}
	major, err := strconv.Atoi(split[0])
	if err != nil {
		return nil, errors.Wrapf(err, "invalid major number in %v", kv)
	}
	minor, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, errors.Wrapf(err, "invalid minor number in %v", kv)
	}
	return &KernelVersion{
		Major: major,
		Minor: minor,
	}, nil
}

// KernelVersionRange - Kernel version range
type KernelVersionRange struct {
	MinVersion *KernelVersion
	MaxVersion *KernelVersion
}

func (kvr KernelVersionRange) String() string {
	return fmt.Sprintf("[%s to %s]", kvr.MinVersion, kvr.MaxVersion)
}

// ExtendMaxKernelVersion - Extends the max kernel version to the provided kernel version (if applicable)
func (kvr *KernelVersionRange) ExtendMaxKernelVersion(kv KernelVersion) {
	if kvr.MaxVersion == nil {
		kvr.MaxVersion = &KernelVersion{
			Major: kv.Major,
			Minor: kv.Minor,
		}
	}
	if kv.After(*kvr.MaxVersion) {
		kvr.MaxVersion.Major = kv.Major
		kvr.MaxVersion.Minor = kv.Minor
	}
}

// ExtendMinKernelVersion - Extends the min kernel version to the provided kernel version (if applicable)
func (kvr *KernelVersionRange) ExtendMinKernelVersion(kv KernelVersion) {
	if kvr.MinVersion == nil {
		kvr.MinVersion = &KernelVersion{
			Major: kv.Major,
			Minor: kv.Minor,
		}
	}
	if kv.Before(*kvr.MinVersion) {
		kvr.MinVersion.Major = kv.Major
		kvr.MinVersion.Minor = kv.Minor
	}
}

// Contains - Returns true if the KernelVersionRange contains the input kernel version
func (kvr KernelVersionRange) Contains(kv KernelVersion) bool {
	return kvr.MinVersion.Before(kv) && kvr.MaxVersion.After(kv)
}

// Iter - Returns an iterator to loop through the possible kernel versions
func (kvr KernelVersionRange) Iter() KernelVersionIter {
	return KernelVersionIter{
		versionRange: &kvr,
		cursor:       nil,
	}
}

// KernelVersionIter - Iterator structure to go over a KernelVersionRange
type KernelVersionIter struct {
	versionRange *KernelVersionRange
	cursor       *KernelVersion
}

// NextMinor - Returns the next minor version after the current cursor
func (kvi *KernelVersionIter) NextMinor() *KernelVersion {
	if kvi.cursor == nil {
		kvi.cursor = kvi.versionRange.MinVersion
		return kvi.cursor
	}
	kvi.cursor.NextMinor()
	if kvi.cursor.After(*kvi.versionRange.MaxVersion) {
		return nil
	}
	return kvi.cursor
}

// NextMajor - Returns the next minor version after the current cursor
func (kvi *KernelVersionIter) NextMajor() *KernelVersion {
	if kvi.cursor == nil {
		kvi.cursor = kvi.versionRange.MinVersion
		return kvi.cursor
	}
	kvi.cursor.NextMajor()
	if kvi.cursor.After(*kvi.versionRange.MaxVersion) {
		return nil
	}
	return kvi.cursor
}
