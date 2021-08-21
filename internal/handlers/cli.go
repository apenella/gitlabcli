package handler

import (
	"fmt"
	"io"

	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

type CliHandler struct {
	writer  io.Writer
	service ports.Service
}

func NewCliHandler(s ports.Service, w io.Writer) (CliHandler, error) {
	return CliHandler{
		service: s,
		writer:  w,
	}, nil
}

func (h CliHandler) GetProject(project string) error {

	errContext := "handler::GetProject"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	data, err := h.service.GetProject(project)
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Could not get project '%s'", project), err)
	}

	fmt.Fprintf(h.writer, "Get project: '%s'\n", project)
	for _, item := range data {
		fmt.Fprintf(h.writer, "%s\n", item)
	}

	return nil
}

func (h CliHandler) ListProjects() error {

	errContext := "handler::ListProjects"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	data, err := h.service.ListProjects()
	if err != nil {
		return errors.New(errContext, "Could not list projects", err)
	}

	fmt.Fprintln(h.writer, "List projects")
	for _, item := range data {
		fmt.Fprintf(h.writer, "%s\n", item)
	}

	return nil
}

func (h CliHandler) ListProjectsFromGroup(group string) error {

	errContext := "handler::ListProjectsFromGroup"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	data, err := h.service.ListProjectsFromGroup(group)
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Could not list projects from group '%s'", group), err)
	}

	fmt.Fprintf(h.writer, "List projects from group: '%s'\n", group)
	for _, item := range data {
		fmt.Fprintf(h.writer, "%s\n", item)
	}

	return nil
}

func (h CliHandler) GetGroup(group string) error {

	errContext := "handler::GetGroup"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	data, err := h.service.GetGroup(group)
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Could not get group '%s'", group), err)
	}

	fmt.Fprintf(h.writer, "Get group: '%s'\n", group)
	for _, item := range data {
		fmt.Fprintf(h.writer, "%s\n", item)
	}
	return nil
}

func (h CliHandler) ListGroups() error {

	errContext := "handler::ListGroups"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	data, err := h.service.ListGroups()
	if err != nil {
		return errors.New(errContext, "Could not list groups")
	}

	fmt.Fprintln(h.writer, "List groups")
	for _, item := range data {
		fmt.Fprintf(h.writer, "%s\n", item)

	}

	return nil
}

func (h CliHandler) CloneProject(project string) error {

	errContext := "handler::CloneProject"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	filter := func(project string) func() ([]domain.Project, error) {
		return func() ([]domain.Project, error) {
			return h.service.GetProject(project)
		}
	}
	err := h.service.Clone(filter(project))
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Project '%s' could not be cloned", project), err)
	}

	return nil
}

func (h CliHandler) CloneProjectFromGroup(group string) error {

	errContext := "handler::CloneProjectFromGroup"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	filter := func(group string) func() ([]domain.Project, error) {
		return func() ([]domain.Project, error) {
			return h.service.ListProjectsFromGroup(group)
		}
	}
	err := h.service.Clone(filter(group))
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("One or more projects from group '%s' could not be cloned", group), err)
	}

	return nil
}

func (h CliHandler) CloneAll() error {

	errContext := "handler::CloneAll"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	filter := func() func() ([]domain.Project, error) {
		return func() ([]domain.Project, error) {
			return h.service.ListProjects()
		}
	}
	err := h.service.Clone(filter())

	if err != nil {
		return errors.New(errContext, "One or more projects could not be cloned", err)
	}

	return nil
}
