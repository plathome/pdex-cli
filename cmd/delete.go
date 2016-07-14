package cmd

import (
	"github.com/urfave/cli"
//	"github.com/plathome/pdex-cli/subcmd"
	"../subcmd"
)

func DeleteCmd() cli.Command {
	command := cli.Command{
		Name:  "delete",
		Usage: "delete list up desired resource",
	}
	command.Subcommands = []cli.Command{
		subCmdDeleteChannel(),
		subCmdDeleteAccount(),
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
