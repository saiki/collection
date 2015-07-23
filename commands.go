package main

import (
	"fmt"
	"os"

	"github.com/a_saiki/collection/command"
	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "add",
		Usage:  "",
		Action: command.CmdAdd,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "list",
		Usage:  "",
		Action: command.CmdList,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "delete",
		Usage:  "",
		Action: command.CmdDelete,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
