package list

import (
	"github.com/apenella/gitlabcli/pkg/command"
	"github.com/spf13/cobra"
)

// NewCommand creates a list command
func NewCommand() *command.AppCommand {
	getCmd := &cobra.Command{
		Use:   "list",
		Short: "List Gitlab contents",
		Long:  `Set of utils to manage Gitlab repositories`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, args)

			return nil
		},
	}

	return command.NewCommand(getCmd)
}
