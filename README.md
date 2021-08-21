# Gitlabcli

Gitlabcli is a command line tool to interactuate with Gitlab repository

- [Gitlabcli](#gitlabcli)
  - [Getting started](#getting-started)
    - [Configuration](#configuration)
    - [Commands](#commands)
  - [Authentication](#authentication)
  - [Features for later updates](#features-for-later-updates)

> **Disclaimer**: Gitlabcli has only been tested on Linux systems

## Getting started

### Configuration
Before start using *gitlabcli* you must create its configuration file.
By default, configuration file location is `~/.config/gitlabcli/config.yml` but you could store it to any location. In that case, `--config` flag must be provided on the command call.

| Parameter  | Type  | Description |
|---|---|---|
| **base_url** | string | Base URL for API requests |
| **gitlab_token** | string | Token to authenticate to Gitlab API |
| **working_dir** | string | Location to store cloned projects |


Example:
```yaml
base_url: https://mygitlab.com/api/v4
gitlab_token: ThatIsAGitlabToken
working_dir: /projects
```

### Commands
- **Clone**: Clone one or multiple projects from Gitlab. It also supports to clone all Gitlab projects or those projects that belong to a group.
- **List**
    - Achieve a list of *projects*
    - Achieve a list of *groups*
- **Get**
    - Get *project* details
    - Get *group* details

## Authentication
*list* and *get* operations uses Gitlab API and requires a Gitlab token.

*Clone* operations only support to clone over ssh and the only supported authentication method is ssh-agent

## Features for later updates
- Clone authentication using a key file
- Clone over http/s and authenticating with user/password
