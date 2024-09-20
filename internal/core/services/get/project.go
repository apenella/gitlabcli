package service

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
)

// GetProjectService struct
type GetProjectService struct {
	gitlab ports.GitlabProjectRepository
}

// NewGetProjectService returns a new GetProjectService
func NewGetProjectService(gitlab ports.GitlabProjectRepository) (GetProjectService, error) {
	s := GetProjectService{
		gitlab: gitlab,
	}

	return s, nil
}

// Get returns a list of projects
func (s GetProjectService) Get(project string) ([]domain.Project, error) {
	return s.gitlab.Find(project)
}
