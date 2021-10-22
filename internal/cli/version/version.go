package version

import (
	"os"
	"runtime"

	"github.com/apenella/gitlabcli/internal/release"
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

			r := release.NewRelease(runtime.GOOS, runtime.GOARCH)
			o := release.NewReleaseOutput(os.Stdout)
			err := o.Text(r)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return command.NewCommand(getCmd)
}
