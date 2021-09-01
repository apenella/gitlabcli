package ports

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
)

type GitlabProjectRepository interface {
	Find(string) ([]domain.Project, error)
	List() ([]domain.Project, error)
}

type GitlabGroupRepository interface {
	Find(string) ([]domain.Group, error)
	List() ([]domain.Group, error)
	ListProjects(string) ([]domain.Project, error)
}
