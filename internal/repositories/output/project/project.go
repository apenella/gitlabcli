package project

import (
	"fmt"
	"io"

	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/ryanuber/columnize"
)

type ProjectOutputRepository struct {
	w io.Writer
}

func NewProjectOutputRepository(w io.Writer) ProjectOutputRepository {
	return ProjectOutputRepository{w}
}

func (o ProjectOutputRepository) Text(g domain.Project) {
	fmt.Fprintln(o.w, g)
}

func (o ProjectOutputRepository) Table(projects []domain.Project) {
	output := []string{"Id|Name|Description|Path|HTTP URL"}

	for _, p := range projects {
		output = append(output, fmt.Sprintf("%d|%s|%s|%s|%s", p.Id, p.Name, p.Description, p.Path, p.Httpurl))
	}

	result := columnize.SimpleFormat(output)
	fmt.Fprintln(o.w, result)
}
