package service

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
)

// GetGroupService struct
type GetGroupService struct {
	gitlab ports.GitlabGroupRepository
}

// NewGetGroupService returns a new GetGroupService
func NewGetGroupService(gitlab ports.GitlabGroupRepository) (GetGroupService, error) {
	s := GetGroupService{
		gitlab: gitlab,
	}

	return s, nil
}

// Get returns a list of groups
func (s GetGroupService) Get(group string) ([]domain.Group, error) {
	return s.gitlab.Find(group)
}
