package get

import (
	"github.com/apenella/gitlabcli/pkg/command"
	"github.com/spf13/cobra"
)

// NewCommand creates a get command
func NewCommand() *command.AppCommand {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get information from Gitlab",
		Long: `Set of utils to manage Gitlab repositories
 get subcommand to get information from Gitlab
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, args)
			return nil
		},
	}

	return command.NewCommand(getCmd)
}
