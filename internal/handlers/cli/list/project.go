package clihandler

import (
	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

type ListProjectCliHandler struct {
	writer  ports.GitlabProjectOutputRepository
	service ports.ListProjectService
}

func NewListProjectCliHandler(s ports.ListProjectService, w ports.GitlabProjectOutputRepository) (ListProjectCliHandler, error) {
	return ListProjectCliHandler{
		service: s,
		writer:  w,
	}, nil
}

func (h ListProjectCliHandler) ListProjects() error {

	errContext := "clihandler::ListProjects"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	data, err := h.service.List()
	if err != nil {
		return errors.New(errContext, "Could not list projects", err)
	}

	h.writer.Table(data)

	return nil
}
