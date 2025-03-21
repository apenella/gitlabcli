package service

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
)

// ListGroupService struct
type ListGroupService struct {
	gitlab ports.GitlabGroupRepository
}

// NewListGroupService returns a new ListGroupService
func NewListGroupService(gitlab ports.GitlabGroupRepository) (ListGroupService, error) {
	s := ListGroupService{
		gitlab: gitlab,
	}

	return s, nil
}

// List returns a list of groups
func (s ListGroupService) List() ([]domain.Group, error) {
	return s.gitlab.List()
}

// ListProjects returns a list of projects that belongs to a group
func (s ListGroupService) ListProjects(group string) ([]domain.Project, error) {
	return s.gitlab.ListProjects(group)
}
