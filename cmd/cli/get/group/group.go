package getgroup

import (
	"fmt"

	"github.com/apenella/gitlabcli/internal/core/ports"
	getservice "github.com/apenella/gitlabcli/internal/core/services/get"
	handler "github.com/apenella/gitlabcli/internal/handlers/cli/get"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/cobra"
)

//var group string

func NewCommand() *command.AppCommand {

	getGroupCmd := &cobra.Command{
		Use:   "group [<group_name>]+",
		Short: "Get group information from Gitlab",
		Long:  `Set of utils to manage Gitlab repositories`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("unregitered handler for RunE on 'group' subcommand")
		},
		Args: cobra.MinimumNArgs(1),
	}
	// getGroupCmd.Flags().StringVarP(&group, "group", "g", "", "Gitlab group name to be consulted")

	return command.NewCommand(getGroupCmd)
}

func RunEHandler(gitlab ports.GitlabGroupRepository, output ports.GitlabGroupOutputRepository) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		var service getservice.GetGroupService
		var h handler.GetGroupCliHandler

		errContext := "getgroup::RunEHandler"

		service, err = getservice.NewGetGroupService(gitlab)
		if err != nil {
			return errors.New(errContext, "Gitlab service could not be created", err)
		}

		h, err = handler.NewGetGroupCliHandler(service, output)
		if err != nil {
			return errors.New(errContext, "Handler cli could not be created", err)
		}

		for _, group := range args {
			err = h.GetGroup(group)

			if err != nil {
				return errors.New(errContext, "Group detail could not be achieved", err)
			}
		}

		return nil
	}
}
