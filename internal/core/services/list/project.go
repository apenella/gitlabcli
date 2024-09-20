package service

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
)

// ListProjectService struct
type ListProjectService struct {
	gitlab ports.GitlabProjectRepository
}

// NewListProjectService returns a new ListProjectService
func NewListProjectService(gitlab ports.GitlabProjectRepository) (ListProjectService, error) {
	s := ListProjectService{
		gitlab: gitlab,
	}

	return s, nil
}

// List returns a list of projects
func (s ListProjectService) List() ([]domain.Project, error) {
	return s.gitlab.List()
}
