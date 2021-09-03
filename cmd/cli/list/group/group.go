package listgroup

import (
	"fmt"

	"github.com/apenella/gitlabcli/internal/core/ports"
	listservice "github.com/apenella/gitlabcli/internal/core/services/list"
	handler "github.com/apenella/gitlabcli/internal/handlers/cli/list"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/cobra"
)

func NewCommand() *command.AppCommand {

	getGroupCmd := &cobra.Command{
		Use:     "groups",
		Aliases: []string{"group"},
		Short:   "List gitlab groups",
		Long:    `Set of utils to manage Gitlab repositories`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("unregitered handler for RunE on 'group' subcommand")
		},
	}

	return command.NewCommand(getGroupCmd)
}

func RunEHandler(gitlab ports.GitlabGroupRepository, outputGroup ports.GitlabGroupOutputRepository) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		var service listservice.ListGroupService
		var h handler.ListGroupCliHandler

		errContext := "listgroup::RunEHandler"

		service, err = listservice.NewListGroupService(gitlab)
		if err != nil {
			return errors.New(errContext, "Gitlab service could not be created", err)
		}

		h, err = handler.NewListGroupCliHandler(service, outputGroup)
		if err != nil {
			return errors.New(errContext, "Handler cli could not be created", err)
		}

		err = h.ListGroups()
		if err != nil {
			return err
		}

		return nil
	}
}
