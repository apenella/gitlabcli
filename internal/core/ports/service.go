package ports

import "github.com/apenella/gitlabcli/internal/core/domain"

// GetGroupService interface
type GetGroupService interface {
	Get(group string) ([]domain.Group, error)
}

// GetProjectService interface
type GetProjectService interface {
	Get(project string) ([]domain.Project, error)
}

// ListGroupService interface
type ListGroupService interface {
	List() ([]domain.Group, error)
	ListProjects(string) ([]domain.Project, error)
}

// ListProjectService interface
type ListProjectService interface {
	List() ([]domain.Project, error)
}

// GitCloneService interface
type GitCloneService interface {
	CloneProject(project string) error
	CloneProjectsFromGroup(group string) error
	CloneAll() error
}
