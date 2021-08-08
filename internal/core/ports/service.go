package ports

import "github.com/apenella/gitlabcli/internal/core/domain"

type Service interface {
	GitlabService
	GitService
}

type GitlabService interface {
	GetProject(project string) ([]domain.Project, error)
	ListProjects() ([]domain.Project, error)
	ListProjectsFromGroup(string) ([]domain.Project, error)
	GetGroup(group string) ([]domain.Group, error)
	ListGroups() ([]domain.Group, error)
}

type GitService interface {
	Clone(filter func() ([]domain.Project, error)) error
	CloneProject(project string) error
	CloneProjectsFromGroup(group string) error
	CloneAll() error
}
