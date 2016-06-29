package cmd

import (
	"github.com/urfave/cli"
	"../cmd"
)

func ConfigureCmd() cli.Command {
	command := cli.Command{
		Name:  "configure",
		Aliases: []string{"c"},
		Usage: "configure profiles",
	}
	command.Subcommands = []cli.Command{
		subCmdConfigure(),
		subCmdConfigureList(),
	}
	return command
}

func subCmdConfigure() cli.Command {
	return cli.Command{
		Name:        "set",
		Description: "list configuration profiles setup",
		Usage:       "configure set --url <endpoint> --accesskey <key>",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "url",
				Value:       "",
				Usage:       "list the devicegroups <devicegroupid>.",
				Destination: &subcmd.FlagUrl,
			},
			cli.StringFlag{
				Name:        "accesskey",
				Value:       "",
				Usage:       "list the devicegroups <devicegroupid>.",
				Destination: &subcmd.FlagAccessKey,
			},
		},
		Action:      subcmd.ConfigureCommands,
	}
}

func subCmdConfigureList() cli.Command {
	return cli.Command{
		Name:        "list",
		Aliases: []string{"ls"},
		Description: "list configuration profiles",
		Usage:       "configure list",
		Action:      subcmd.ListConfigureCommands,
	}
}
