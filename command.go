package cli

import (
	"flag"
	"os"
)

type Command struct {
	Name        string
	Short       string
	HelpFunc    func()
	HandlerFunc func(args []string)
	Flags       *flag.FlagSet
	LocalFlags  *flag.FlagSet
	Subcommands map[string]*Command
}

func NewCommand(name string) *Command {
	command := &Command{
		Name: name,
	}
	command.Flags = flag.NewFlagSet(name, flag.ContinueOnError)
	command.LocalFlags = flag.NewFlagSet(name, flag.ContinueOnError)

	command.HelpFunc = func() {
		command.Usage()
	}

	command.HandlerFunc = func(args []string) {
		command.HelpFunc()
	}

	return command
}

func (c *Command) Usage() {
	c.Flags.PrintDefaults()
}

func (c *Command) AddCommands(cmds ...*Command) {
	if c.Subcommands == nil {
		c.Subcommands = make(map[string]*Command)
	}

	for _, cmd := range cmds {
		c.Subcommands[cmd.Name] = cmd
		if cmd.Short != "" {
			c.Subcommands[cmd.Short] = cmd
		}
	}
}

func (c *Command) Execute(args []string) {
	if args == nil {
		args = os.Args[1:]
	}

	// Parse global flags
	c.Flags.Parse(args)
	args = c.Flags.Args()

	// Parse local flags
	c.LocalFlags.Parse(args)
	args = c.LocalFlags.Args()

	if len(args) <= 0 {
		c.HandlerFunc(nil)
		return
	}

	subCommand, ok := c.Subcommands[args[0]]
	if ok {
		subCommand.Execute(args[1:])
	} else {
		c.HandlerFunc(args)
	}
}
