package clihandler

import (
	"fmt"

	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

type GetGroupCliHandler struct {
	writer  ports.GitlabGroupOutputRepository
	service ports.GetGroupService
}

func NewGetGroupCliHandler(s ports.GetGroupService, w ports.GitlabGroupOutputRepository) (GetGroupCliHandler, error) {
	return GetGroupCliHandler{
		service: s,
		writer:  w,
	}, nil
}

func (h GetGroupCliHandler) GetGroup(group string) error {

	errContext := "clihandler::GetGroup"

	if h.service == nil {
		return errors.New(errContext, "Handler service is not defined")
	}

	data, err := h.service.Get(group)
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Could not get group '%s'", group), err)
	}

	h.writer.Table(data)

	return nil
}
