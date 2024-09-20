package version

import (
	"os"
	"runtime"

	"github.com/apenella/gitlabcli/internal/release"
	"github.com/apenella/gitlabcli/pkg/command"
	"github.com/spf13/cobra"
)

// NewCommand creates a version command
func NewCommand() *command.AppCommand {
	getCmd := &cobra.Command{
		Use:               "version",
		Short:             "gitlabcli version",
		Long:              "gitlabcli version",
		PersistentPreRunE: nil,
		RunE: func(cmd *cobra.Command, args []string) error {

			r := release.NewRelease(runtime.GOOS, runtime.GOARCH)
			o := release.NewOutput(os.Stdout)
			err := o.Text(r)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return command.NewCommand(getCmd)
}
