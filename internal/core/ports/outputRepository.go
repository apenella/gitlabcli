package ports

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
)

type GitlabGroupOutputRepository interface {
	Text(g domain.Group)
	Table(groups []domain.Group)
}

type GitlabProjectOutputRepository interface {
	Text(g domain.Project)
	Table(groups []domain.Project)
}
