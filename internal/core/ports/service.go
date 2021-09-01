package ports

import "github.com/apenella/gitlabcli/internal/core/domain"

type GetGroupService interface {
	Get(group string) ([]domain.Group, error)
}

type GetProjectService interface {
	Get(project string) ([]domain.Project, error)
}

type ListGroupService interface {
	List() ([]domain.Group, error)
	ListProjects(string) ([]domain.Project, error)
}

type ListProjectService interface {
	List() ([]domain.Project, error)
}

type GitCloneService interface {
	CloneProject(project string) error
	CloneProjectsFromGroup(group string) error
	CloneAll() error
}
