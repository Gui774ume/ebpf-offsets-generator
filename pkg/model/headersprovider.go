package model

// HeadersProvider - Header provider
type HeadersProvider interface {
	HasVersion(kv KernelVersion) bool
	PullHeaders(kv KernelVersion) ([]string, error)
	SetTmpPath(path string)
	KernelVersionsList() []KernelVersion
}
