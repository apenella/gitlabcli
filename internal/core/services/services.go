package service

import (
	"fmt"
	"path/filepath"

	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

type OptionsFunc func(*Service)

const (
	SSHMode uint8 = iota
	HTTPMode
)

type Service struct {
	gitlab           ports.GitlabRepository
	git              ports.GitRepository
	mode             uint8
	useNamespacePath bool
	basePath         string
	dir              string
}

func WithMode(m uint8) OptionsFunc {
	return func(s *Service) {
		s.mode = m
	}
}

func WithUseNamespacePath() OptionsFunc {
	return func(s *Service) {
		s.useNamespacePath = true
	}
}

func WithBasePath(p string) OptionsFunc {
	return func(s *Service) {
		s.basePath = p
	}
}

func WithDir(d string) OptionsFunc {
	return func(s *Service) {
		s.dir = d
	}
}

func New(gitlab ports.GitlabRepository, git ports.GitRepository, opts ...OptionsFunc) (Service, error) {
	s := &Service{
		gitlab: gitlab,
		git:    git,
	}

	for _, opt := range opts {
		opt(s)
	}

	return *s, nil
}

func (s Service) GetProject(project string) ([]domain.Project, error) {
	return s.gitlab.FindProject(project)
}

func (s Service) ListProjects() ([]domain.Project, error) {
	return s.gitlab.ListProjects()
}

func (s Service) ListProjectsFromGroup(group string) ([]domain.Project, error) {
	return s.gitlab.ListProjectsFromGroup(group)
}

func (s Service) GetGroup(group string) ([]domain.Group, error) {
	return s.gitlab.FindGroup(group)
}

func (s Service) ListGroups() ([]domain.Group, error) {
	return s.gitlab.ListGroups()
}

func (s Service) Clone(filter func() ([]domain.Project, error)) error {

	errContext := "service::Clone"
	errs := []error{}

	projects, err := filter()
	if err != nil {
		return errors.New(errContext, "Error filtering projects list to be cloned", err)
	}

	for _, p := range projects {
		err := s.clone(p)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.New(errContext, "One ore more projects could not be cloned", errs...)
	}

	return nil
}

func (s Service) CloneProject(project string) error {

	errContext := "service::CloneProject"
	errs := []error{}

	projects, err := s.gitlab.FindProject(project)
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Project '%s' could not be found", project), err)
	}

	for _, p := range projects {
		err := s.clone(p)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.New(errContext, "One ore more projects could not be cloned", errs...)
	}

	return nil
}

func (s Service) CloneProjectsFromGroup(group string) error {
	errContext := "service::CloneProjectsFromGroup"
	errs := []error{}

	projects, err := s.gitlab.ListProjectsFromGroup(group)
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Project from group '%s' could not be achieved", group), err)
	}

	for _, p := range projects {
		err := s.clone(p)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.New(errContext, "One ore more projects could not be cloned", errs...)
	}

	return nil
}
func (s Service) CloneAll() error {

	errContext := "service::CloneProject"
	errs := []error{}

	projects, err := s.gitlab.ListProjects()
	if err != nil {
		return errors.New(errContext, "Project could not be achieved", err)
	}

	for _, p := range projects {
		err := s.clone(p)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.New(errContext, "One ore more projects could not be cloned", errs...)
	}

	return errors.New(errContext, fmt.Sprintf("'%s' not implemented", errContext))
}

func (s Service) clone(p domain.Project) error {

	errContext := "service::clone"

	dir := s.directory(p)
	url := s.url(p)
	fmt.Printf("Cloning '%s' into '%s'... ", p.Name, dir)
	err := s.git.Clone(dir, url)
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Project '%s' could not be cloned", p.Name), err)
	}
	fmt.Println("DONE")

	return nil
}

func (s Service) directory(project domain.Project) string {

	if s.dir != "" {
		return s.dir
	}

	directory := s.basePath

	if s.useNamespacePath {
		directory = filepath.Join(directory, project.Path)
	}

	return directory
}

func (s Service) url(project domain.Project) string {
	var url string

	switch s.mode {
	case SSHMode:
		url = project.Sshurl
	case HTTPMode:
		url = project.Httpurl
	}

	return url
}
