package ports

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
)

// GitlabGroupOutputRepository interface
type GitlabGroupOutputRepository interface {
	Text(g domain.Group)
	Table(groups []domain.Group)
}

// GitlabProjectOutputRepository interface
type GitlabProjectOutputRepository interface {
	Text(g domain.Project)
	Table(groups []domain.Project)
}
