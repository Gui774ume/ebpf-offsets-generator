package headersproviders

import (
	"github.com/Gui774ume/ebpf-offsets-generator/pkg/headersproviders/ubuntu"
	"github.com/Gui774ume/ebpf-offsets-generator/pkg/model"
)

// Providers - Map of available headers provider
var Providers = map[string]model.HeadersProvider{
	"ubuntu": &ubuntu.Provider{},
}
