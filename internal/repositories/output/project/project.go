package project

import (
	"fmt"
	"io"

	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/ryanuber/columnize"
)

// OutputRepository is responsible to output project information
type OutputRepository struct {
	w io.Writer
}

// NewOutputRepository returns a new OutputRepository
func NewOutputRepository(w io.Writer) *OutputRepository {
	return &OutputRepository{w}
}

// Text writes a project information in text format
func (o *OutputRepository) Text(g domain.Project) {
	fmt.Fprintln(o.w, g)
}

// Table writes a project information in table format
func (o *OutputRepository) Table(projects []domain.Project) {
	output := []string{"Id|Name|Description|Path|HTTP URL"}

	for _, p := range projects {
		output = append(output, fmt.Sprintf("%d|%s|%s|%s|%s", p.ID, p.Name, p.Description, p.Path, p.Httpurl))
	}

	result := columnize.SimpleFormat(output)
	fmt.Fprintln(o.w, result)
}
