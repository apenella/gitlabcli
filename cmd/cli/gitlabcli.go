package gitlabcli

import (
	"fmt"
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
	gitlabrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab"
	gitlabgrouprepo "github.com/apenella/gitlabcli/internal/repositories/gitlab/group"
	gitlabprojectrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab/project"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const DefaultConfigFile = "config"

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
	//var gitRepo gitrepo.GitRepository
	var gitlab gitlabrepo.GitlabRepository
	//var glSrv ports.Service
	//var glStorage storagerepo.ProjectStorage
	//var cliHandler handler.CliHandler

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

			// gitRepo, err = gitrepo.NewGitRepository()
			// if err != nil {
			// 	return errors.New(errContext, "Git repository could not be created", err)
			// }

			gitlab, err = gitlabrepo.NewGitlabRepository(conf.Token, conf.BaseURL, 100)
			if err != nil {
				return errors.New(errContext, "Gitlab repository could not be created", err)
			}

			// glStorage = storagerepo.New(afero.NewOsFs())

			// glSrv, err = service.New(
			// 	glRepo,
			// 	gitRepo,
			// 	glStorage,
			// 	service.WithUseNamespacePath(),
			// 	service.WithBasePath(conf.WorkingDir),
			// )

			// if err != nil {
			// 	return errors.New(errContext, "Gitlab service could not be created", err)
			// }

			// cliHandler, err = handler.NewCliHandler(glSrv, os.Stdout)
			// if err != nil {
			// 	return errors.New(errContext, "Handler cli could not be created", err)
			// }

			listGroupsSubcommand.Options(command.WithRunE(listgroup.RunEHandler(
				gitlabgrouprepo.GitlabGroupRepository{gitlab},
			)))

			listProjectsSubcommand.Options(command.WithRunE(listproject.RunEHandler(
				gitlabprojectrepo.GitlabProjectRepository{gitlab},
			)))

			getGroupSubcommand.Options(command.WithRunE(getgroup.RunEHandler(
				gitlabgrouprepo.GitlabGroupRepository{gitlab},
			)))

			getProjectSubcommand.Options(command.WithRunE(getproject.RunEHandler(
				gitlabprojectrepo.GitlabProjectRepository{gitlab},
			)))

			cloneSubcommand.Options(
				command.WithRunE(clone.RunEHandler(
					gitlabprojectrepo.GitlabProjectRepository{gitlab},
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
