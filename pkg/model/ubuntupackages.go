package model

import (
	"fmt"
	"sync"
)

// UbuntuPackages - This structure is used to hold the all and generic linux header for the Ubuntu provider
type UbuntuPackages struct {
	LinuxHeaderAll     string
	LinuxHeaderGeneric string
}

func (up *UbuntuPackages) String() string {
	return fmt.Sprintf("%v", *up)
}

// UbuntuDatabase - This structure holds all the ubuntu linux headers required to compile the programs of the offsets-generator
type UbuntuDatabase struct {
	sync.Mutex
	Packages map[KernelVersion]*UbuntuPackages
}

// NewUbuntuDatabase - Returns a new ubuntu database
func NewUbuntuDatabase() UbuntuDatabase {
	return UbuntuDatabase{
		Packages: map[KernelVersion]*UbuntuPackages{},
	}
}
