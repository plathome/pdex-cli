package cmd

import (
	"github.com/urfave/cli"
//	"github.com/plathome/pdex-cli/subcmd"
	"../subcmd"
)

func UpdateCmd() cli.Command {
	command := cli.Command {
		Name:  		"update",
		Aliases:	[]string{"up"},
		Usage: 		"update session",
	}
	command.Subcommands = []cli.Command{
		subCmdUpdateSession(),
	}
	return command
}

func subCmdUpdateSession() cli.Command {
	return cli.Command {
		Name:        	"session",
		Aliases: 		[]string{"s"},
		Description: 	"update session",
		Usage:       	"update session",
		Action:      	subcmd.UpdateSession,
	}
}
