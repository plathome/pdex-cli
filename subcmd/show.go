package cmd

import (
	"github.com/urfave/cli"
	"github.com/plathome/pdex-cli/cmd"
//	"../cmd"
)

func ShowCmd() cli.Command {
	command := cli.Command{
		Name:  "show",
		Usage: "show details of resource",
	}
	command.Subcommands = []cli.Command{
		subCmdMe(),
		subCmdDG(),
	}
	return command
}

func subCmdMe() cli.Command {
	return cli.Command{
		Name:        "me",
		Description: "show user resources",
		Usage:       "show me",
		Action:      subcmd.ShowMyself,
	}
}

func subCmdDG() cli.Command {
	return cli.Command{
		Name:        "devicegroups",
		Aliases: []string{"dg"},
		Description: "devicegroups --deid-prefix",
		Usage:       "show devicegroups --deid-prefix DEVICE-ID-PREFIX",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid-prefix",
				Value:       "",
				Usage:       "show devicegroups --deid-prefix DEVICE-ID-PREFIX",
				Destination: &subcmd.FlagDeviceGroup,
			},
		},
		Action: subcmd.ShowDeviceGroup,
	}
}
