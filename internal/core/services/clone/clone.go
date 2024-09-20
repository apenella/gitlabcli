package clone

import (
	"fmt"
	"path/filepath"

	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

// OptionsFunc to set options to Service
type OptionsFunc func(*Service)

const (
	// SSHMode mode
	SSHMode uint8 = iota
	// HTTPMode mode
	HTTPMode
)

// Service struct defines a service to clone projects
type Service struct {
	project          ports.GitlabProjectRepository
	group            ports.GitlabGroupRepository
	git              ports.GitRepository
	storage          ports.StorageRepository
	mode             uint8
	useNamespacePath bool
	basePath         string
	dir              string
}

// WithMode is an OptionsFunc to set mode to Service
func WithMode(m uint8) OptionsFunc {
	return func(s *Service) {
		s.mode = m
	}
}

// WithUseNamespacePath is an OptionsFunc to set useNamespacePath to Service
func WithUseNamespacePath() OptionsFunc {
	return func(s *Service) {
		s.useNamespacePath = true
	}
}

// WithBasePath is an OptionsFunc to set basePath to Service
func WithBasePath(p string) OptionsFunc {
	return func(s *Service) {
		s.basePath = p
	}
}

// WithDir is an OptionsFunc to set dir to Service
func WithDir(d string) OptionsFunc {
	return func(s *Service) {
		s.dir = d
	}
}

// NewService creates a new Service instance
func NewService(project ports.GitlabProjectRepository, group ports.GitlabGroupRepository, git ports.GitRepository, storage ports.StorageRepository, opts ...OptionsFunc) (Service, error) {
	s := &Service{
		project: project,
		group:   group,
		git:     git,
		storage: storage,
	}

	for _, opt := range opts {
		opt(s)
	}

	return *s, nil
}

// CloneProject clones a project
func (s Service) CloneProject(project string) error {

	errContext := "service::CloneProject"
	errs := []error{}

	projects, err := s.project.Find(project)
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
		return errors.New(errContext, "One or more projects could not be cloned", errs...)
	}

	return nil
}

// CloneProjectsFromGroup clones all projects from a GitLab group
func (s Service) CloneProjectsFromGroup(group string) error {
	errContext := "service::CloneProjectsFromGroup"
	errs := []error{}

	projects, err := s.group.ListProjects(group)
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
		return errors.New(errContext, "One or more projects could not be cloned", errs...)
	}

	return nil
}

// CloneAll clones all projects
func (s Service) CloneAll() error {

	errContext := "service::CloneProject"
	errs := []error{}

	projects, err := s.project.List()
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
		return errors.New(errContext, "One or more projects could not be cloned", errs...)
	}

	return errors.New(errContext, fmt.Sprintf("'%s' not implemented", errContext))
}

// clone clones a project
func (s Service) clone(p domain.Project) error {

	var err error
	var existDir bool

	errContext := "service::clone"

	dir := s.directory(p)

	existDir, err = s.storage.DirExists(dir)
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Error validating if '%s' directory exist", dir), err)
	}

	if existDir {
		fmt.Printf("Project '%s' already cloned on '%s'\n", p.Name, dir)
		return nil
	}

	url := s.url(p)
	fmt.Printf("Cloning '%s' into '%s'... ", p.Name, dir)
	err = s.git.Clone(dir, url)
	_ = url
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Project '%s' could not be cloned", p.Name), err)
	}
	fmt.Println("DONE")

	return nil
}

// directory returns the directory to clone a project
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

// url returns the url to clone a project
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
