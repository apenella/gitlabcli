# Gitlabcli

Gitlabcli is a command line tool to interactuate with Gitlab repository

- [Gitlabcli](#gitlabcli)
  - [Getting started](#getting-started)
    - [Installation](#installation)
      - [Installation from tarball](#installation-from-tarball)
    - [Configuration](#configuration)
      - [Environment variables](#environment-variables)
      - [Configuration parameteres](#configuration-parameteres)
    - [Commands](#commands)
    - [Authentication](#authentication)
  - [Comming updates, new features or ideas](#comming-updates-new-features-or-ideas)
  - [License](#license)

> **Disclaimer**: gitlabcli has only been tested on Linux systems

## Getting started

### Installation

#### Installation from tarball
- Download `gitlabcli` tarball from github releases
```sh
$ curl -sLO https://github.com/apenella/gitlabcli/releases/download/v0.3.0/gitlabcli_0.3.0_Linux-x86_64.tar.gz
```

- Untar `gitlabcli` package
```sh
$ tar xzfv gitlabcli_0.2.0_Linux-x86_64.tar.gz
```

- Start using `gitlabcli`
```sh
$ gitlabcli -help
Set of utils to manage Gitlab repositories

Usage:
  gitlabcli [flags]
  gitlabcli [command]

Available Commands:
  clone       Clone repositories from Gitlab to localhost
  completion  generate the autocompletion script for the specified shell
  get         Get information from Gitlab
  help        Help about any command
  initialize  Initializes gitlabcli
  list        List Gitlab contents
  version     gitlabcli version

Flags:
      --config string   Configuration file
  -h, --help            help for gitlabcli

Use "gitlabcli [command] --help" for more information about a command.
```

### Configuration
Before start using *gitlabcli* you must create its configuration file.
By default, configuration file location is `~/.config/gitlabcli/config.yml` but you could store it to any location. In that case, `--config` flag must be provided on the command call.

You could run the `initialize` subcommand to initialize `gitlabcli`. That command takes care to initialize the configuration parameters properly.
```sh
$ gitlabcli initialize --gitlab-api-url https://mygitlab.com/api/v4 --working-dir /projects
```

#### Environment variables
`gitlabcli` supports environment variables configuration. In that case, environment variables must be named as uppercased parameter name and prefixed by `GITLABCLI_`. For instance, `working_dir` parameter could be configured by `GITLABCLI_WORKING_DIR` environment variable.

#### Configuration parameteres

| Parameter  | Type  | Description |
|---|---|---|
| **gitlab_api_url** | string | Gitlab API URL base. Check it on [Gitlab documentation](https://docs.gitlab.com/ee/api/#how-to-use-the-api) |
| **gitlab_token** | string | Token to authenticate to Gitlab API |
| **working_dir** | string | Location to store cloned projects |


Example:
```yaml
gitlab_api_url: https://mygitlab.com/api/v4
gitlab_token: ThatIsAGitlabToken
working_dir: /projects
```

### Commands
- **Clone**: Clone one or multiple projects from Gitlab. It also supports to clone all Gitlab projects or those projects that belong to a group.
- **List**
    - Achieve a Gitlab *projects* list
    - Achieve a Gitlab *groups* list
- **Get**
    - Get *project* details
    - Get *group* details
- **Initialize**: Initialize gitlabcli configuration

### Authentication
*list* and *get* operations uses Gitlab API and requires a Gitlab token.

*Clone* operations only support to clone over ssh and the only supported authentication method is ssh-agent

## Comming updates, new features or ideas
- Private key file authentication when cloning projects
- User/pass authentication when cloning projects over HTTP/S
  
## License
gitlabcli is available under [MIT](https://github.com/apenella/gitlabcli/blob/master/LICENSE) license.
