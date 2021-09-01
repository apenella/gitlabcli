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

// type GitlabService interface {
// 	GitlabProjectService
// 	GitlabGroupService
// }

// type GitlabProjectService interface {
// 	GetProject(project string) ([]domain.Project, error)
// 	ListProjects() ([]domain.Project, error)
// 	ListProjectsFromGroup(string) ([]domain.Project, error)
// }

// type GitlabGroupService interface {
// 	GetGroup(group string) ([]domain.Group, error)
// 	ListGroups() ([]domain.Group, error)
// }

// type GitService interface {
// 	Clone(filter func() ([]domain.Project, error)) error
// 	CloneProject(project string) error
// 	CloneProjectsFromGroup(group string) error
// 	CloneAll() error
// }
