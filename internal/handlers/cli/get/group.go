package clihandler

import (
	"fmt"

	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

// GetGroupCliHandler handles get group command
type GetGroupCliHandler struct {
	writer  ports.GitlabGroupOutputRepository
	service ports.GetGroupService
}

// NewGetGroupCliHandler returns a new GetGroupCliHandler
func NewGetGroupCliHandler(s ports.GetGroupService, w ports.GitlabGroupOutputRepository) (GetGroupCliHandler, error) {
	return GetGroupCliHandler{
		service: s,
		writer:  w,
	}, nil
}

// GetGroup handles command to get a group
func (h GetGroupCliHandler) GetGroup(groups ...string) error {

	errContext := "clihandler::GetGroup"

	result := []domain.Group{}

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	for _, group := range groups {
		data, err := h.service.Get(group)
		if err != nil {
			return errors.New(errContext, fmt.Sprintf("Could not get group '%s'", group), err)
		}

		result = append(result, data...)
	}

	h.writer.Table(result)

	return nil
}
