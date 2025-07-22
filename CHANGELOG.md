# CHANGELOG

## [unreleased] (yyyy-mm-dd)

### Changed

- Bump golang.org/x/oauth2 from 0.18.0 to 0.27.0
- Bump github.com/cloudflare/circl from 1.3.7 to 1.6.1
- Bump github.com/go-git/go-git/v5 from 5.12.0 to 5.13.0
- Bump golang.org/x/crypto from 0.27.0 to 0.31.0
- Bump golang.org/x/net from 0.29.0 to 0.38.0

## [v0.3.2]

### Added

- Include the staticcheck tool

### Changed

- Adapt souce code to make it complaiant to golint, staticcheck, errcheck and gosec
- Bump go-git to v5.12.0 that fixes the security vulnerabilities CVE-2023-49569
- Upgrade to Golang 1.23

## [v0.3.1]

### Fixed

- Ensure configuration folder is created to avoid a non existing directory error on initialize command

## [v0.3.0]

### Added

- On clone subcommand, added a flag to set the working directory
- `base_url` configuration attribute has been renamed to `gitlab_api_url`
- on `initialize` command `base-url` flag has been renamed to `gitlab-api-url`

## [v0.2.0]

### Added

- output group and projects is managed by output repository

### Changed

- Move repositories creation from PersistentPreRun to each command
- PresistentPreRun on root command only creates the configuration
- handler clone allow to clone multiple projectes
- handler get group allow to get details for multiple groups
- handler get project allow to get details for multiple projects

## [v0.1.0]

### Added

- Command to list Gitlab projects
- Command to get Gitlab project details 
- Command to list Gitlab groups
- Command to get Gitlab groups details
- Command to clone projects from Gitlab
