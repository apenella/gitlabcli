package service

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
)

type GetProjectService struct {
	gitlab ports.GitlabProjectRepository
}

func NewGetProjectService(gitlab ports.GitlabProjectRepository) (GetProjectService, error) {
	s := GetProjectService{
		gitlab: gitlab,
	}

	return s, nil
}

func (s GetProjectService) Get(project string) ([]domain.Project, error) {
	return s.gitlab.Find(project)
}
