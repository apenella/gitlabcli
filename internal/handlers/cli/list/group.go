package clihandler

import (
	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

// ListGroupCliHandler handles list group command
type ListGroupCliHandler struct {
	writer  ports.GitlabGroupOutputRepository
	service ports.ListGroupService
}

// NewListGroupCliHandler returns a new ListGroupCliHandler
func NewListGroupCliHandler(s ports.ListGroupService, w ports.GitlabGroupOutputRepository) (ListGroupCliHandler, error) {
	return ListGroupCliHandler{
		service: s,
		writer:  w,
	}, nil
}

// ListGroups handles command to list groups
func (h ListGroupCliHandler) ListGroups() error {

	errContext := "clihandler::ListGroups"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	data, err := h.service.List()
	if err != nil {
		return errors.New(errContext, "Could not list groups")
	}

	h.writer.Table(data)

	return nil
}
