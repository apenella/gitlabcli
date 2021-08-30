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

func (h ListProjectCliHandler) ListProjectsFromGroup(group string) error {

	errContext := "clihandler::ListProjectsFromGroup"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	data, err := h.service.ListFromGroup(group)
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Could not list projects from group '%s'", group), err)
	}

	fmt.Fprintf(h.writer, "List projects from group: '%s'\n", group)
	for _, item := range data {
		fmt.Fprintf(h.writer, "%s\n", item)
	}

	return nil
}
