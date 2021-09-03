package group

import (
	"fmt"
	"io"

	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/ryanuber/columnize"
)

type GroupOutputRepository struct {
	w io.Writer
}

func NewGroupOutputRepository(w io.Writer) GroupOutputRepository {
	return GroupOutputRepository{w}
}

func (o GroupOutputRepository) Text(g domain.Group) {
	fmt.Fprintln(o.w, g)
}

func (o GroupOutputRepository) Table(groups []domain.Group) {
	output := []string{"Id|Group|Description|Path|Web Url"}

	for _, g := range groups {
		output = append(output, fmt.Sprintf("%d|%s|%s|%s|%s", g.Id, g.Name, g.Description, g.Path, g.WebUrl))
	}

	result := columnize.SimpleFormat(output)
	fmt.Fprintln(o.w, result)
}
