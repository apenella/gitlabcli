package release

import (
	"fmt"
	"strings"
)

var (
	// Version is the version of the release
	Version string
	// Commit is the commit hash of the release
	Commit string
	// BuildDate is the date when the release was built
	BuildDate string
)

// Release contains information about the release
type Release struct {
	// BuildDate is the date when the release was built
	BuildDate string
	// Commit is the commit hash of the release
	Commit string
	// Header is the header of the release
	Header string
	// OsArch is the OS and architecture of the release
	OsArch string
	// Version is the version of the release
	Version string
}

// NewRelease creates a new Release instance
func NewRelease(os, arch string) *Release {
	return &Release{
		BuildDate: strings.TrimSpace(BuildDate),
		Version:   strings.TrimSpace(Version),
		Commit:    strings.TrimSpace(Commit),
		OsArch:    fmt.Sprintf("%s/%s", os, arch),
	}
}
