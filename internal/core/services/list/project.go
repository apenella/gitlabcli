package service

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
)

type ListProjectService struct {
	gitlab ports.GitlabProjectRepository
}

func NewListProjectService(gitlab ports.GitlabProjectRepository) (ListProjectService, error) {
	s := &ListProjectService{
		gitlab: gitlab,
	}

	return *s, nil
}

func (s ListProjectService) List() ([]domain.Project, error) {
	return s.gitlab.List()
}

func (s ListProjectService) ListFromGroup(group string) ([]domain.Project, error) {
	return s.gitlab.ListFromGroup(group)
}
