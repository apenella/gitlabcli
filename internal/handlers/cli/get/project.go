package clihandler

import (
	"fmt"
	"io"

	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

type GetProjectCliHandler struct {
	writer  io.Writer
	service ports.GetProjectService
}

func NewGetProjectCliHandler(s ports.GetProjectService, w io.Writer) (GetProjectCliHandler, error) {
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

	fmt.Fprintf(h.writer, "Get project: '%s'\n", project)
	for _, item := range data {
		fmt.Fprintf(h.writer, "%s\n", item)
	}

	return nil
}
