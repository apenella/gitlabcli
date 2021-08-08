package ports

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
)

type GitlabRepository interface {
	GitlabProjectRepository
	GitlabGroupRepository
}

type GitlabProjectRepository interface {
	FindProject(string) ([]domain.Project, error)
	ListProjects() ([]domain.Project, error)
	ListProjectsFromGroup(string) ([]domain.Project, error)
}

type GitlabGroupRepository interface {
	FindGroup(string) ([]domain.Group, error)
	ListGroups() ([]domain.Group, error)
}
