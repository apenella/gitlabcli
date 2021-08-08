package listgroup

import (
	"fmt"

	handler "github.com/apenella/gitlabcli/internal/handlers"
	"github.com/apenella/gitlabcli/pkg/command"
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

func RunEHandler(h handler.CliHandler) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		err := h.ListGroups()
		if err != nil {
			return err
		}

		return nil
	}
}
