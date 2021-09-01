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

type ReleaseOutput struct {
	Writer io.Writer
}

func NewReleaseOutput(w io.Writer) ReleaseOutput {
	return ReleaseOutput{
		Writer: w,
	}
}

func (v ReleaseOutput) Text(r *Release) error {
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
