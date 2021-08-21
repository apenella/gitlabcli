package main

import (
	"fmt"
	"os"

	gitlabcli "github.com/apenella/gitlabcli/cmd/cli"
)

func main() {
	var err error

	cmd := gitlabcli.NewCommand()
	err = cmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
