package cmd

import (
	"github.com/urfave/cli"
	"github.com/plathome/pdex-cli/cmd"
//	"../cmd"
)

func CreateCmd() cli.Command {
	command := cli.Command{
		Name:  "create",
		Usage: "create new resource",
	}
	command.Subcommands = []cli.Command{
		subCmdCreateDG(),
	}
	return command
}

func subCmdCreateDG() cli.Command {
	return cli.Command{
		Name:        	"devicegroups",
		Aliases: 		[]string{"dg"},
		Description: 	"create new devicegroup resources",
		Usage:       	"create devicegroups",
		Action:      	subcmd.CreateDG,
	}
}
