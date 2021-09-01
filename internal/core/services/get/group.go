package service

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
)

type GetGroupService struct {
	gitlab ports.GitlabGroupRepository
}

func NewGetGroupService(gitlab ports.GitlabGroupRepository) (GetGroupService, error) {
	s := GetGroupService{
		gitlab: gitlab,
	}

	return s, nil
}

func (s GetGroupService) Get(group string) ([]domain.Group, error) {
	return s.gitlab.Find(group)
}
