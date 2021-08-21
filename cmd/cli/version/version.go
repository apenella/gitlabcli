package version

import (
	"os"

	"github.com/apenella/gitlabcli/cmd/release"
	"github.com/apenella/gitlabcli/pkg/command"
	"github.com/spf13/cobra"
)

func NewCommand() *command.AppCommand {
	getCmd := &cobra.Command{
		Use:               "version",
		Short:             "gitlabcli version",
		Long:              "gitlabcli version",
		PersistentPreRunE: nil,
		RunE: func(cmd *cobra.Command, args []string) error {

			r := release.NewRelease(os.Stdout)
			err := r.PrintVersion()
			if err != nil {
				return err
			}
			return nil
		},
	}

	return command.NewCommand(getCmd)
}
