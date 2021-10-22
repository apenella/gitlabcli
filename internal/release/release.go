package release

import (
	"fmt"
	"strings"
)

var (
	Version, Commit, BuildDate string
)

// Release
type Release struct {
	BuildDate string
	Commit    string
	Header    string
	OsArch    string
	Version   string
}

// NewRelease
func NewRelease(os, arch string) *Release {
	return &Release{
		BuildDate: strings.TrimSpace(BuildDate),
		Version:   strings.TrimSpace(Version),
		Commit:    strings.TrimSpace(Commit),
		OsArch:    fmt.Sprintf("%s/%s", os, arch),
	}
}
