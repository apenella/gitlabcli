package group

import (
	"fmt"
	"io"

	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/ryanuber/columnize"
)

// OutputRepository is responsible to output group information
type OutputRepository struct {
	w io.Writer
}

// NewOutputRepository returns a new OutputRepository
func NewOutputRepository(w io.Writer) *OutputRepository {
	return &OutputRepository{w}
}

// Text writes a group information in text format
func (o *OutputRepository) Text(g domain.Group) {
	fmt.Fprintln(o.w, g)
}

// Table writes a group information in table format
func (o *OutputRepository) Table(groups []domain.Group) {
	output := []string{"Id|Group|Description|Path|Web Url"}

	for _, g := range groups {
		output = append(output, fmt.Sprintf("%d|%s|%s|%s|%s", g.ID, g.Name, g.Description, g.Path, g.WebURL))
	}

	result := columnize.SimpleFormat(output)
	fmt.Fprintln(o.w, result)
}
