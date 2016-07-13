package cmd

import (
	"github.com/urfave/cli"
//	"github.com/plathome/pdex-cli/subcmd"
	"../subcmd"
)

func CreateCmd() cli.Command {
	command := cli.Command {
		Name:  		"create",
		Aliases:	[]string{"cr"},
		Usage: 		"create new resource",
	}
	command.Subcommands = []cli.Command{
		subCmdCreateDG(),
		subCmdCreateApp(),
		subCmdCreateDevice(),
		subCmdCreateChannel(),
	}
	return command
}

func subCmdCreateDG() cli.Command {
	return cli.Command {
		Name:        	"devicegroups",
		Aliases: 		[]string{"dg"},
		Description: 	"create new devicegroup resource",
		Usage:       	"create devicegroups",
		Action:      	subcmd.CreateDG,
	}
}

func subCmdCreateApp() cli.Command {
	return cli.Command{
		Name:        "apps",
		Description: "create new app",
		Usage:       "create apps --app-name-suffix APP-NAME-SUFFIX",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "app-name-suffix",
				Value:       "",
				Usage:       "create apps --app-name-suffix APP-NAME-SUFFIX",
				Destination: &subcmd.FlagAppNameSuffix,
			},
		},
		Action: subcmd.CreateApp,
	}
}

func subCmdCreateDevice() cli.Command {
	return cli.Command{
		Name:        	"devices",
		Aliases: 		[]string{"de"},
		Description: 	"create new device",
		Usage:       	"create devices --deid-prefix DEVICE-ID-SUFFIX",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid-prefix",
				Value:       "",
				Usage:       "create devices --deid-prefix DEVICE-ID-PREFIX",
				Destination: &subcmd.FlagDeviceGroup,
			},
		},
		Action: subcmd.CreateDevice,
	}
}

func subCmdCreateChannel() cli.Command {
	return cli.Command{
		Name:        	"channels",
		Aliases:		[]string{"ch"},
		Description: 	"create new channel",
		Usage:       	"create channels --deid DEVICE-ID --app-id APP-ID",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid",
				Value:       "",
				Usage:       "create devices --deid DEVICE-ID --app-id APP-ID",
				Destination: &subcmd.FlagDeviceId,
			},
			cli.StringFlag{
				Name:        "app-id",
				Value:       "",
				Usage:       "create devices --deid DEVICE-ID --app-id APP-ID",
				Destination: &subcmd.FlagAppId,
			},
		},
		Action: subcmd.CreateChannel,
	}
}
