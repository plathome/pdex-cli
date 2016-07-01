package cmd

import (
	"github.com/urfave/cli"
	"github.com/plathome/pdex-cli/cmd"
//	"../cmd"
)

func SendCmd() cli.Command {
	command := cli.Command{
		Name:  "send",
		Aliases: []string{"s"},
		Usage: "send messages",
	}
	command.Subcommands = []cli.Command{
		subCmdSendCommands(),
		subCmdSendMessages(),
	}
	return command
}

func subCmdSendCommands() cli.Command {
	return cli.Command{
		Name:        "commands",
		Aliases: []string{"cmd"},
		Description: "sending commands to devices",
		Usage:       "send commands --channelid=channelid",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "channel-id",
				Value:       "",
				Usage:       "send commands --channelid=<channelid>.",
				Destination: &subcmd.FlagChannelId,
			},
			cli.StringFlag{
				Name:        "app-id",
				Value:       "",
				Usage:       "send commands --appid=<appid>.",
				Destination: &subcmd.FlagAppId,
			},
		},
		Action: subcmd.SendCommands,
	}
}

func subCmdSendMessages() cli.Command {
	return cli.Command{
		Name:        "messages",
		Aliases: []string{"msg"},
		Description: "sending messages to apps, channels",
		Usage:       "send messages --deid DEVICE_ID",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid",
				Value:       "",
				Usage:       "send messages --deid DEVICE_ID",
				Destination: &subcmd.FlagDeviceId,
			},
		},
		Action: subcmd.SendMessages,
	}
}
