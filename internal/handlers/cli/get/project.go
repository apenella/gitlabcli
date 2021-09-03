package clihandler

import (
	"fmt"

	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

type GetProjectCliHandler struct {
	writer  ports.GitlabProjectOutputRepository
	service ports.GetProjectService
}

func NewGetProjectCliHandler(s ports.GetProjectService, w ports.GitlabProjectOutputRepository) (GetProjectCliHandler, error) {
	return GetProjectCliHandler{
		service: s,
		writer:  w,
	}, nil
}

func (h GetProjectCliHandler) GetProject(project string) error {

	errContext := "clihandler::GetProject"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	data, err := h.service.Get(project)
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Could not get project '%s'", project), err)
	}

	h.writer.Table(data)

	return nil
}
