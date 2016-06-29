package cmd

import (
	"github.com/urfave/cli"
	"github.com/plathome/pdex-cli/cmd"
)

func ReadCmd() cli.Command {
	command := cli.Command{
		Name:  "read",
		Aliases: []string{"r"},
		Usage: "read messages",
	}
	command.Subcommands = []cli.Command{
		subCmdReadCommands(),
		subCmdReadMessages(),
	}
	return command
}

func subCmdReadCommands() cli.Command {
	return cli.Command{
		Name:        "commands",
		Aliases: []string{"cmd"},
		Description: "read commands from channels",
		Usage:       "read commands",
		Action:      subcmd.ReadCommands,
	}
}

func subCmdReadMessages() cli.Command {
	return cli.Command{
		Name:        "messages",
		Aliases: []string{"msg"},
		Description: "read messages from devices",
		Usage:       "read messages",
		Action:      subcmd.ReadMessages,
	}
}
