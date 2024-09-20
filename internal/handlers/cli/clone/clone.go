package clihandler

import (
	"fmt"
	"io"

	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

// CloneCliHandler handles clone project command
type CloneCliHandler struct {
	writer  io.Writer
	service ports.GitCloneService
}

// NewCloneCliHandler returns a new CloneCliHandler
func NewCloneCliHandler(s ports.GitCloneService, w io.Writer) (CloneCliHandler, error) {
	return CloneCliHandler{
		service: s,
		writer:  w,
	}, nil
}

// CloneProject handles command to clone of a project
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

// CloneProjectFromGroup handles command to clone of all projects from a group
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

// CloneAll handles command to clone of all projects
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
