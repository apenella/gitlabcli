package clihandler

import (
	"fmt"
	"io"

	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

type ListProjectCliHandler struct {
	writer  io.Writer
	service ports.ListProjectService
}

func NewListProjectCliHandler(s ports.ListProjectService, w io.Writer) (ListProjectCliHandler, error) {
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

	fmt.Fprintln(h.writer, "List projects")
	for _, item := range data {
		fmt.Fprintf(h.writer, "%s\n", item)
	}

	return nil
}
