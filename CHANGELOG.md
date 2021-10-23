# CHANGELOG

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
