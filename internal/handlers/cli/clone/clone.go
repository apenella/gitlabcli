package clihandler

import (
	"fmt"
	"io"

	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

type CloneCliHandler struct {
	writer  io.Writer
	service ports.GitCloneService
}

func NewCloneCliHandler(s ports.GitCloneService, w io.Writer) (CloneCliHandler, error) {
	return CloneCliHandler{
		service: s,
		writer:  w,
	}, nil
}

func (h CloneCliHandler) CloneProject(projects ...string) error {

	errContext := "clihandler::CloneProject"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	for _, project := range projects {
		err := h.service.CloneProject(project)
		if err != nil {
			return errors.New(errContext, fmt.Sprintf("Project '%s' could not be cloned", project), err)
		}
	}

	return nil
}

func (h CloneCliHandler) CloneProjectFromGroup(group string) error {

	errContext := "clihandler::CloneProjectFromGroup"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	err := h.service.CloneProjectsFromGroup(group)
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("One or more projects from group '%s' could not be cloned", group), err)
	}

	return nil
}

func (h CloneCliHandler) CloneAll() error {

	errContext := "clihandler::CloneAll"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	err := h.service.CloneAll()
	if err != nil {
		return errors.New(errContext, "One or more projects could not be cloned", err)
	}

	return nil
}
