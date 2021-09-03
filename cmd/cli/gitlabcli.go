package gitlabcli

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/apenella/gitlabcli/cmd/cli/clone"
	"github.com/apenella/gitlabcli/cmd/cli/get"
	getgroup "github.com/apenella/gitlabcli/cmd/cli/get/group"
	getproject "github.com/apenella/gitlabcli/cmd/cli/get/project"
	"github.com/apenella/gitlabcli/cmd/cli/initialize"
	"github.com/apenella/gitlabcli/cmd/cli/list"
	listgroup "github.com/apenella/gitlabcli/cmd/cli/list/group"
	listproject "github.com/apenella/gitlabcli/cmd/cli/list/project"
	"github.com/apenella/gitlabcli/cmd/cli/version"
	"github.com/apenella/gitlabcli/cmd/configuration"
	loadconfiguration "github.com/apenella/gitlabcli/cmd/configuration/load"
	"github.com/apenella/gitlabcli/internal/core/ports"
	gitlabrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab"
	gitlabgrouprepo "github.com/apenella/gitlabcli/internal/repositories/gitlab/group"
	gitlabprojectrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab/project"
	groupoutputrepo "github.com/apenella/gitlabcli/internal/repositories/output/group"
	projectoutputrepo "github.com/apenella/gitlabcli/internal/repositories/output/project"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const DefaultConfigFile = "config"

const PerPage = 100

var viperconfig *viper.Viper

func init() {
	user, err := user.Current()
	if err != nil {
		panic(fmt.Sprintf("current user information can not be achieved. %s\n", err.Error()))
	}

	viperconfig = viper.New()
	viperconfig.AutomaticEnv()
	viperconfig.SetEnvPrefix("gitlabcli")
	viperconfig.SetConfigName(DefaultConfigFile)
	viperconfig.SetConfigType("yaml")
	viperconfig.AddConfigPath(filepath.Join(user.HomeDir, ".config", "gitlabcli"))
}

var configFile string

func NewCommand() *command.AppCommand {
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
	var gitlab gitlabrepo.GitlabRepository
	var outputGroup ports.GitlabGroupOutputRepository
	var outputProject ports.GitlabProjectOutputRepository

	gitlabCmd := &cobra.Command{
		Use:   "gitlabcli",
		Short: "gitlab cli",
		Long:  `Set of utils to manage Gitlab repositories`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, args)
			return nil
		},
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

			v := validator.New()
			err = conf.Validate(v)
			if err != nil {
				return errors.New(errContext, "Invalid configuration", err)
			}

			gitlab, err = gitlabrepo.NewGitlabRepository(conf.Token, conf.BaseURL, PerPage)
			if err != nil {
				return errors.New(errContext, "Gitlab repository could not be created", err)
			}

			outputGroup = groupoutputrepo.NewGroupOutputRepository(os.Stdout)
			outputProject = projectoutputrepo.NewProjectOutputRepository(os.Stdout)

			listGroupsSubcommand.Options(command.WithRunE(listgroup.RunEHandler(
				gitlabgrouprepo.NewGitlabGroupRepository(gitlab.Client.Groups, PerPage),
				outputGroup,
			)))

			listProjectsSubcommand.Options(command.WithRunE(listproject.RunEHandler(
				gitlabprojectrepo.NewGitlabProjectRepository(gitlab.Client.Projects, PerPage),
				gitlabgrouprepo.NewGitlabGroupRepository(gitlab.Client.Groups, PerPage),
				outputProject,
			)))

			getGroupSubcommand.Options(command.WithRunE(getgroup.RunEHandler(
				gitlabgrouprepo.NewGitlabGroupRepository(gitlab.Client.Groups, PerPage),
				outputGroup,
			)))

			getProjectSubcommand.Options(command.WithRunE(getproject.RunEHandler(
				gitlabprojectrepo.NewGitlabProjectRepository(gitlab.Client.Projects, PerPage),
				outputProject,
			)))

			cloneSubcommand.Options(
				command.WithRunE(clone.RunEHandler(
					gitlabprojectrepo.NewGitlabProjectRepository(gitlab.Client.Projects, PerPage),
					gitlabgrouprepo.NewGitlabGroupRepository(gitlab.Client.Groups, PerPage),
					conf.WorkingDir)))

			return nil
		},
	}
	gitlabCmd.PersistentFlags().StringVar(&configFile, "config", "", "Configuration file")

	cloneSubcommand = clone.NewCommand()
	getGroupSubcommand = getgroup.NewCommand()
	getProjectSubcommand = getproject.NewCommand()
	getSubcommand = get.NewCommand()
	initSubcommand = initialize.NewCommand(viperconfig)
	listGroupsSubcommand = listgroup.NewCommand()
	listProjectsSubcommand = listproject.NewCommand()
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

	return gitlabCommand
}
