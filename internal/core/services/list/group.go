package service

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
)

type ListGroupService struct {
	gitlab ports.GitlabGroupRepository
}

func NewListGroupService(gitlab ports.GitlabGroupRepository) (ListGroupService, error) {
	s := &ListGroupService{
		gitlab: gitlab,
	}

	return *s, nil
}

func (s ListGroupService) List() ([]domain.Group, error) {
	return s.gitlab.List()
}
