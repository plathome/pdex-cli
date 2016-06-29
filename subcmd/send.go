package cmd

import (
	"github.com/urfave/cli"
	"../cmd"
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
				Name:        "channelid",
				Value:       "",
				Usage:       "send commands --channelid=<channelid>.",
				Destination: &subcmd.FlagChannelId,
			},
		},
		Action: subcmd.SendCommands,
	}
}

func subCmdSendMessages() cli.Command {
	return cli.Command{
		Name:        "messages",
		Aliases: []string{"cmd"},
		Description: "sending messages to apps, channels",
		Usage:       "send messages --deviceid=deviceid",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deviceid",
				Value:       "",
				Usage:       "send messages --deviceid=<deviceid>.",
				Destination: &subcmd.FlagDeviceId,
			},
		},
		Action: subcmd.SendMessages,
	}
}
