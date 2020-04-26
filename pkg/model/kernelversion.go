package model

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// KernelVersion - Kernel version representation
type KernelVersion struct {
	Major int
	Minor int
	Patch int
}

func (kv KernelVersion) String() string {
	if kv.Patch > 0 {
		return fmt.Sprintf("v%d.%d.%d", kv.Major, kv.Minor, kv.Patch)
	}
	return fmt.Sprintf("v%d.%d", kv.Major, kv.Minor)
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

// IsNeighborOf - Returns true if the input version differs from the current version from 1 patch
func (kv KernelVersion) IsNeighborOf(version KernelVersion) bool {
	return math.Abs(float64(kv.Int()-version.Int())) == 1
}

// Copy - Returns a copy of the current kernel version
func (kv KernelVersion) Copy() *KernelVersion {
	return &KernelVersion{
		Major: kv.Major,
		Minor: kv.Minor,
		Patch: kv.Patch,
	}
}

// Int - Returns an integer representation of the kernel version
func (kv KernelVersion) Int() int {
	return kv.Major*1000000 + kv.Minor*1000 + kv.Patch
}

// NextPatch - Increments the current kernel version to the next patch version
func (kv *KernelVersion) NextPatch() {
	kv.Patch += 1
}

// NextMinor - Increments the current kernel version to the next minor version
func (kv *KernelVersion) NextMinor() {
	kv.Minor += 1
	kv.Patch = 0
}

// NextMajor - Increments the current kernel version to the next major version
func (kv *KernelVersion) NextMajor() {
	kv.Major += 1
	kv.Minor = 0
	kv.Patch = 0
}

// NewKernelVersionFromString - Creates a new kernel version instance from the string representation of a kernel version
func NewKernelVersionFromString(kv string) (*KernelVersion, error) {
	// Remove any trailing `/` or `v` character
	kv = strings.ReplaceAll(kv, "v", "")
	kv = strings.ReplaceAll(kv, "/", "")
	split := strings.Split(kv, ".")
	if len(split) == 2 {
		split = append(split, "0")
	}
	if len(split) != 3 {
		return nil, errors.Errorf("the following kernel version doesn't follow the pattern x.y.z: %v", kv)
	}
	major, err := strconv.Atoi(split[0])
	if err != nil {
		return nil, errors.Wrapf(err, "invalid major number in %v", kv)
	}
	minor, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, errors.Wrapf(err, "invalid minor number in %v", kv)
	}
	patch, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, errors.Wrapf(err, "invalid patch number in %v", kv)
	}
	return &KernelVersion{
		Major: major,
		Minor: minor,
		Patch: patch,
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
			Patch: kv.Patch,
		}
	}
	if kv.After(*kvr.MaxVersion) {
		kvr.MaxVersion.Major = kv.Major
		kvr.MaxVersion.Minor = kv.Minor
		kvr.MaxVersion.Patch = kv.Patch
	}
}

// ExtendMinKernelVersion - Extends the min kernel version to the provided kernel version (if applicable)
func (kvr *KernelVersionRange) ExtendMinKernelVersion(kv KernelVersion) {
	if kvr.MinVersion == nil {
		kvr.MinVersion = &KernelVersion{
			Major: kv.Major,
			Minor: kv.Minor,
			Patch: kv.Patch,
		}
	}
	if kv.Before(*kvr.MinVersion) {
		kvr.MinVersion.Major = kv.Major
		kvr.MinVersion.Minor = kv.Minor
		kvr.MinVersion.Patch = kv.Patch
	}
}

// Contains - Returns true if the KernelVersionRange contains the input kernel version
func (kvr KernelVersionRange) Contains(kv KernelVersion) bool {
	return kvr.MinVersion.Before(kv) && kvr.MaxVersion.After(kv) || kvr.MinVersion.Equal(kv) || kvr.MaxVersion.Equal(kv)
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

// NextPatch - Returns the next patch version after the current cursor
func (kvi *KernelVersionIter) NextPatch() *KernelVersion {
	if kvi.cursor == nil {
		kvi.cursor = kvi.versionRange.MinVersion.Copy()
		return kvi.cursor
	}
	kvi.cursor.NextPatch()
	if kvi.cursor.After(*kvi.versionRange.MaxVersion) {
		return nil
	}
	return kvi.cursor
}

// NextMinor - Returns the next minor version after the current cursor
func (kvi *KernelVersionIter) NextMinor() *KernelVersion {
	if kvi.cursor == nil {
		kvi.cursor = kvi.versionRange.MinVersion.Copy()
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
		kvi.cursor = kvi.versionRange.MinVersion.Copy()
		return kvi.cursor
	}
	kvi.cursor.NextMajor()
	if kvi.cursor.After(*kvi.versionRange.MaxVersion) {
		return nil
	}
	return kvi.cursor
}

// ByVersion - Implements sort.Interface for []KernelVersion based on version
type ByVersion []KernelVersion

func (a ByVersion) Len() int           { return len(a) }
func (a ByVersion) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVersion) Less(i, j int) bool { return a[i].Before(a[j]) }
