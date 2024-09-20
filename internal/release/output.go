package release

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"text/template"

	errors "github.com/apenella/go-common-utils/error"
)

const (
	versionTmpl = `gitlabcli {{ .Version }} Commit: {{ .Commit }} {{ .OsArch }} BuildDate: {{ .BuildDate }}`
)

// Output represents the output for a release
type Output struct {
	Writer io.Writer
}

// NewOutput returns a new Output
func NewOutput(w io.Writer) *Output {
	return &Output{
		Writer: w,
	}
}

// Text writes a release information in text format
func (v *Output) Text(r *Release) error {
	errContext := "release::Text"

	var w bytes.Buffer

	if v.Writer == nil {
		v.Writer = os.Stdout
	}

	tmpl, err := template.New("version").Parse(versionTmpl)
	if err != nil {
		return errors.New(errContext, "Error parsing version template", err)
	}

	err = tmpl.Execute(io.Writer(&w), r)
	if err != nil {
		return errors.New(errContext, "Error appling version parsed template", err)
	}

	fmt.Fprintln(v.Writer, w.String())

	return nil
}
