package gitlabcli

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/apenella/gitlabcli/internal/cli/clone"
	"github.com/apenella/gitlabcli/internal/cli/get"
	getgroup "github.com/apenella/gitlabcli/internal/cli/get/group"
	getproject "github.com/apenella/gitlabcli/internal/cli/get/project"
	"github.com/apenella/gitlabcli/internal/cli/initialize"
	"github.com/apenella/gitlabcli/internal/cli/list"
	listgroup "github.com/apenella/gitlabcli/internal/cli/list/group"
	listproject "github.com/apenella/gitlabcli/internal/cli/list/project"
	"github.com/apenella/gitlabcli/internal/cli/version"
	"github.com/apenella/gitlabcli/internal/configuration"
	loadconfiguration "github.com/apenella/gitlabcli/internal/configuration/load"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/go-playground/validator"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// DefaultConfigFile is the default configuration file name
const DefaultConfigFile = "config"

var viperconfig *viper.Viper

func init() {
	user, err := user.Current()
	if err != nil {
		panic(fmt.Sprintf("current user information can not be achieved. %s\n", err.Error()))
	}

	configPath := filepath.Join(user.HomeDir, ".config", "gitlabcli")

	err = os.MkdirAll(configPath, 0755)
	if err != nil {
		panic(fmt.Sprintf("configuration directory can not be created. %s\n", err.Error()))
	}

	viperconfig = viper.New()
	viperconfig.AutomaticEnv()
	viperconfig.SetEnvPrefix("gitlabcli")
	viperconfig.SetConfigName(DefaultConfigFile)
	viperconfig.SetConfigType("yaml")
	viperconfig.AddConfigPath(configPath)
}

var configFile string

// NewCommand creates a gitlabcli command
func NewCommand() (*command.AppCommand, error) {
	var err error
	var conf configuration.Configuration
	var cloneSubcommand,
		getGroupSubcommand,
		getProjectSubcommand,
		getSubcommand,
		initSubcommand,
		listGroupsSubcommand,
		listProjectsSubcommand,
		listSubcommand,
		versionSubcommand *command.AppCommand

	gitlabCmd := &cobra.Command{
		Use:           "gitlabcli",
		Short:         "Gitlab command line interface",
		Long:          `Gitlab command line interface offers a set of utils to manage Gitlab repositories`,
		SilenceErrors: true,
		SilenceUsage:  true,
		Run:           func(cmd *cobra.Command, args []string) {},

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			errContext := "gitlabcli::NewCommand::PersistentPreRunE"

			err = loadconfiguration.Load(viperconfig, afero.NewOsFs(), configFile)
			if err != nil {
				return errors.New(errContext, "Configuration could not be loaded", err)
			}

			conf, err = loadconfiguration.Unmarshal(viperconfig)
			if err != nil {
				return errors.New(errContext, "Configuration could not be unmarshaled", err)
			}

			conf.FixCompatibility()

			v := validator.New()
			err = conf.Validate(v)
			if err != nil {
				return errors.New(errContext, "Invalid configuration", err)
			}

			return nil
		},
	}

	gitlabCmd.PersistentFlags().StringVar(&configFile, "config", "", "Configuration file")

	cloneSubcommand = clone.NewCommand(&conf)
	getGroupSubcommand = getgroup.NewCommand(&conf)
	getProjectSubcommand = getproject.NewCommand(&conf)
	getSubcommand = get.NewCommand()
	initSubcommand = initialize.NewCommand(viperconfig)
	listGroupsSubcommand = listgroup.NewCommand(&conf)
	listProjectsSubcommand = listproject.NewCommand(&conf)
	listSubcommand = list.NewCommand()
	versionSubcommand = version.NewCommand()

	gitlabCommand := &command.AppCommand{
		Command: gitlabCmd,
	}

	getSubcommand.AddCommands(
		getGroupSubcommand,
		getProjectSubcommand)

	listSubcommand.AddCommands(
		listGroupsSubcommand,
		listProjectsSubcommand)

	gitlabCommand.AddCommands(
		getSubcommand,
		initSubcommand,
		listSubcommand,
		cloneSubcommand,
		versionSubcommand)

	return gitlabCommand, nil
}
