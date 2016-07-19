package cmd

import (
	"github.com/urfave/cli"
	"github.com/plathome/pdex-cli/subcmd"
//	"../subcmd"
)

func DeleteCmd() cli.Command {
	command := cli.Command{
		Name:  "delete",
		Usage: "delete list up desired resource",
	}
	command.Subcommands = []cli.Command{
		subCmdDeleteChannel(),
		subCmdDeleteAccount(),
		subCmdDeleteDgTagKey(),
		subCmdDeleteDeviceTagKey(),
		subCmdDeleteAppTagKey(),
	}
	return command
}

func subCmdDeleteChannel() cli.Command {
	return cli.Command{
		Name:        	"channels",
		Aliases: 		[]string{"ch"},
		Description: 	"delete channels",
		Usage:       	"delete ch --deid DEID --channel-id CHANNEL-ID --confirm true/false",
		Flags: []cli.Flag {
			cli.StringFlag {
				Name:        "deid",
				Value:       "",
				Usage:       "delete ch --deid DEID --channel-id CHANNEL-ID --confirm true/false",
				Destination: &subcmd.FlagDeviceId,
			},
			cli.StringFlag {
				Name:        "channel-id",
				Value:       "",
				Usage:       "delete ch --deid DEID --channel-id CHANNEL-ID --confirm true/false",
				Destination: &subcmd.FlagChannelId,
			},
			cli.StringFlag {
				Name:        "confirm",
				Value:       "",
				Usage:       "delete ch --deid DEID --channel-id CHANNEL-ID --confirm true/false",
				Destination: &subcmd.FlagConfirmation,
			},
		},
		Action:      	subcmd.DeleteChannel,
	}
}

func subCmdDeleteAccount() cli.Command {
	return cli.Command{
		Name:        	"account",
		Aliases: 		[]string{"ac"},
		Description: 	"delete account",
		Usage:       	"delete ac --confirm true/false",
		Flags: []cli.Flag {
			cli.StringFlag {
				Name:        "confirm",
				Value:       "",
				Usage:       "delete ac --confirm true/false",
				Destination: &subcmd.FlagConfirmation,
			},
		},
		Action:      	subcmd.DeleteAccount,
	}
}

func subCmdDeleteDgTagKey() cli.Command {
	return cli.Command{
		Name:        "dg-tags",
		Description: "delete tag key of devicegroup",
		Usage:       "delete dg-tags --deid-prefix DEVICE-ID-PREFIX --key KEY",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid-prefix",
				Value:       "",
				Usage:       "delete dg-tags --deid-prefix DEVICE-ID-PREFIX --key KEY",
				Destination: &subcmd.FlagDeviceGroup,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "delete dg-tags --deid-prefix DEVICE-ID-PREFIX --key KEY",
				Destination: &subcmd.FlagKey,
			},
		},
		Action:      subcmd.DeleteDgTagKey,
	}
}

func subCmdDeleteDeviceTagKey() cli.Command {
	return cli.Command{
		Name:        "de-tags",
		Description: "delete tag key of device",
		Usage:       "delete de-tags --deid DEVICE-ID --key KEY",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid",
				Value:       "",
				Usage:       "delete de-tags --deid DEVICE-ID --key KEY",
				Destination: &subcmd.FlagDeviceId,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "delete de-tags --deid DEVICE-ID --key KEY",
				Destination: &subcmd.FlagKey,
			},
		},
		Action:      subcmd.DeleteDeviceTagKey,
	}
}

func subCmdDeleteAppTagKey() cli.Command {
	return cli.Command{
		Name:        "ap-tags",
		Description: "delete tag key of app",
		Usage:       "delete ap-tags --app-id APP-ID --key KEY",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "app-id",
				Value:       "",
				Usage:       "delete app-tags --app-id APP-ID --key KEY",
				Destination: &subcmd.FlagAppId,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "delete app-tags --app-id APP-ID --key KEY",
				Destination: &subcmd.FlagKey,
			},
		},
		Action:      subcmd.DeleteApppTagKey,
	}
}
