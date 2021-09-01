package clihandler

import (
	"fmt"
	"io"

	"github.com/apenella/gitlabcli/internal/core/ports"
	errors "github.com/apenella/go-common-utils/error"
)

type GetGroupCliHandler struct {
	writer  io.Writer
	service ports.GetGroupService
}

func NewGetGroupCliHandler(s ports.GetGroupService, w io.Writer) (GetGroupCliHandler, error) {
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

	fmt.Fprintf(h.writer, "Get group: '%s'\n", group)
	for _, item := range data {
		fmt.Fprintf(h.writer, "%s\n", item)
	}
	return nil
}
