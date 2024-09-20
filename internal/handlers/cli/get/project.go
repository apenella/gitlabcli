package clihandler

import (
	"fmt"

	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

// GetProjectCliHandler handles get project command
type GetProjectCliHandler struct {
	writer  ports.GitlabProjectOutputRepository
	service ports.GetProjectService
}

// NewGetProjectCliHandler returns a new GetProjectCliHandler
func NewGetProjectCliHandler(s ports.GetProjectService, w ports.GitlabProjectOutputRepository) (GetProjectCliHandler, error) {
	return GetProjectCliHandler{
		service: s,
		writer:  w,
	}, nil
}

// GetProject handles command to get a project
func (h GetProjectCliHandler) GetProject(projects ...string) error {

	errContext := "clihandler::GetProject"

	result := []domain.Project{}

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	for _, project := range projects {
		data, err := h.service.Get(project)
		if err != nil {
			return errors.New(errContext, fmt.Sprintf("Could not get project '%s'", project), err)
		}

		result = append(result, data...)
	}

	h.writer.Table(result)

	return nil
}
