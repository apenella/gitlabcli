package clihandler

import (
	"fmt"
	"io"

	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

type ListGroupCliHandler struct {
	writer  io.Writer
	service ports.ListGroupService
}

func NewListGroupCliHandler(s ports.ListGroupService, w io.Writer) (ListGroupCliHandler, error) {
	return ListGroupCliHandler{
		service: s,
		writer:  w,
	}, nil
}

func (h ListGroupCliHandler) ListGroups() error {

	errContext := "clihandler::ListGroups"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	data, err := h.service.List()
	if err != nil {
		return errors.New(errContext, "Could not list groups")
	}

	fmt.Fprintln(h.writer, "List groups")
	for _, item := range data {
		fmt.Fprintf(h.writer, "%s\n", item)

	}

	return nil
}
