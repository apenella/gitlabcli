package clihandler

import (
	"fmt"

	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

// ListGroupProjectCliHandler handles list group project command
type ListGroupProjectCliHandler struct {
	writer  ports.GitlabProjectOutputRepository
	service ports.ListGroupService
}

// NewListGroupProjectCliHandler returns a new ListGroupProjectCliHandler
func NewListGroupProjectCliHandler(s ports.ListGroupService, w ports.GitlabProjectOutputRepository) (ListGroupProjectCliHandler, error) {
	return ListGroupProjectCliHandler{
		service: s,
		writer:  w,
	}, nil
}

// ListProjectsFromGroup handles command to list projects from a group
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
