package main

import (
	"fmt"
	"os"

	gitlabcli "github.com/apenella/gitlabcli/internal/cli"
	"github.com/apenella/gitlabcli/pkg/command"
)

func main() {
	var err error
	var cmd *command.AppCommand

	cmd, err = gitlabcli.NewCommand()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = cmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
