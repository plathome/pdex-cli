package cmd

import (
	"github.com/urfave/cli"
//	"github.com/plathome/pdex-cli/cmd"
	"../cmd"
)

func ReadCmd() cli.Command {
	command := cli.Command{
		Name:  "read",
		Aliases: []string{"r"},
		Usage: "read messages",
	}
	command.Subcommands = []cli.Command{
		subCmdReadCommands(),
		subCmdReadMessages(),
	}
	return command
}

func subCmdReadCommands() cli.Command {
	return cli.Command{
		Name:        "commands",
		Aliases: []string{"cmd"},
		Description: "read commands from channels",
		Usage:       "read commands",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deviceid",
				Value:       "",
				Usage:       "read commands --deviceid=<deviceid>.",
				Destination: &subcmd.FlagDeviceId,
			},
			cli.StringFlag{
				Name:        "cmdid",
				Value:       "",
				Usage:       "read commands --deviceid=<deviceid> --cmdif=<cmdid>.",
				Destination: &subcmd.FlagCmdId,
			},
		},
		Action:      subcmd.ReadCommands,
	}
}

func subCmdReadMessages() cli.Command {
	return cli.Command{
		Name:        "messages",
		Aliases: []string{"msg"},
		Description: "read messages from devices",
		Usage:       "read messages",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "appid",
				Value:       "",
				Usage:       "read messages --appid=<appid>.",
				Destination: &subcmd.FlagAppId,
			},
			cli.StringFlag{
				Name:        "msgid",
				Value:       "",
				Usage:       "read messages --appid=<appid> --msgid=<mgsid>.",
				Destination: &subcmd.FlagMsgId,
			},
		},
		Action:      subcmd.ReadMessages,
	}
}
