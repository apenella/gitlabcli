# Gitlabcli

Gitlabcli is a command line tool to interactuate with Gitlab repository.

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

## Operations supported*
- **Clone**: Clone one or multiple projects from Gitlab. It also supports to all Gitlab projects or those projects that belong to a group. 
- **List**
    - Achieve a list of *projects*
    - Achieve a list of *groups*
- **Get**
    - Get *projects* details
    - Get *groups* details

## Authentication
*List* and *get* operations requires a Gitlab token.
*Clone* operations only support clones over ssh and authentication through those keys loaded to an ssh agent

## Features for later updates
- Clone authentication using a key file
- Clone over http/s and authenticating with user/password
- Testing
- Automatic building and releases
