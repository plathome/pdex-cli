package cmd

import (
	"github.com/urfave/cli"
	"github.com/plathome/pdex-cli/cmd"
)

func ShowCmd() cli.Command {
	command := cli.Command {
		Name:  		"show",
		Aliases:	[]string{"sh"},
		Usage: 		"show details of resource",
	}
	command.Subcommands = []cli.Command{
		subCmdMe(),
		subCmdDG(),
		subCmdApp(),
		subCmdChannel(),
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
		Name:        	"devicegroups",
		Aliases: 		[]string{"dg"},
		Description: 	"devicegroups --deid-prefix",
		Usage:       	"show devicegroups --deid-prefix DEVICE-ID-PREFIX",
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

func subCmdApp() cli.Command {
	return cli.Command{
		Name:        "apps",
		Description: "show apps details info",
		Usage:       "show apps --app-id APP-ID",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "app-id",
				Value:       "",
				Usage:       "show apps --app-id APP-ID",
				Destination: &subcmd.FlagAppId,
			},
		},
		Action: subcmd.ShowApp,
	}
}

func subCmdChannel() cli.Command {
	return cli.Command{
		Name:        	"channels",
		Aliases:		[]string{"ch"},
		Description: 	"show channels",
		Usage:       	"show channels --deid DEVICE-ID --app-id APP-ID",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid",
				Value:       "",
				Usage:       "show channels --deid DEVICE-ID --channel-id CHANNEL-ID",
				Destination: &subcmd.FlagDeviceId,
			},
			cli.StringFlag{
				Name:        "channel-id",
				Value:       "",
				Usage:       "show channels --deid DEVICE-ID --channel-id CHANNEL-ID",
				Destination: &subcmd.FlagChannelId,
			},
			cli.StringFlag{
				Name:        "app-id",
				Value:       "",
				Usage:       "show channels --app-id APP-ID --channel-id CHANNEL-ID",
				Destination: &subcmd.FlagAppId,
			},
		},
		Action: subcmd.ShowChannel,
	}
}

