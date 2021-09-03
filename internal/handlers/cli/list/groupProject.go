package clihandler

import (
	"fmt"

	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

type ListGroupProjectCliHandler struct {
	writer  ports.GitlabProjectOutputRepository
	service ports.ListGroupService
}

func NewListGroupProjectCliHandler(s ports.ListGroupService, w ports.GitlabProjectOutputRepository) (ListGroupProjectCliHandler, error) {
	return ListGroupProjectCliHandler{
		service: s,
		writer:  w,
	}, nil
}

func (h ListGroupProjectCliHandler) ListProjectsFromGroup(group string) error {

	errContext := "clihandler::ListProjectsFromGroup"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	data, err := h.service.ListProjects(group)
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Could not list projects from group '%s'", group), err)
	}

	h.writer.Table(data)

	return nil
}
