package getgroup

import (
	"fmt"

	handler "github.com/apenella/gitlabcli/internal/handlers"
	"github.com/apenella/gitlabcli/pkg/command"
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

func RunEHandler(h handler.CliHandler) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {

		for _, group := range args {
			err := h.GetGroup(group)

			if err != nil {
				return err
			}
		}

		return nil
	}
}
