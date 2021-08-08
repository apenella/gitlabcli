package command

import (
	"github.com/spf13/cobra"
)

// CobraRunFunc is a cobra handler function
type CobraRunFunc func(cmd *cobra.Command, args []string)

// CobraRunEFunc is a cobra handler function which returns an error
type CobraRunEFunc func(cmd *cobra.Command, args []string) error

type CommandOptionsFunc func(c *AppCommand)

// AppCommand defines a stevedore command element
type AppCommand struct {
	Command *cobra.Command
}

func NewCommand(cmd *cobra.Command) *AppCommand {
	return &AppCommand{
		Command: cmd,
	}
}

// AddCommand method add a new subcommand to stevedore command
func (c *AppCommand) AddCommands(cmds ...*AppCommand) {
	for _, cmd := range cmds {
		c.Command.AddCommand(cmd.Command)
	}
}

// Execute executes cobra command
func (c *AppCommand) Execute() error {
	return c.Command.Execute()
}

func (c *AppCommand) Options(opts ...CommandOptionsFunc) {

	for _, opt := range opts {
		opt(c)
	}
}

func WithRun(f CobraRunFunc) CommandOptionsFunc {
	return func(c *AppCommand) {
		c.Command.Run = f
	}
}

func WithRunE(f CobraRunEFunc) CommandOptionsFunc {
	return func(c *AppCommand) {
		c.Command.RunE = f
	}
}
