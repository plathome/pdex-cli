package cmd

import (
	"github.com/urfave/cli"
	"github.com/plathome/pdex-cli/subcmd"
//	"../subcmd"
)

func ReadCmd() cli.Command {
	command := cli.Command {
		Name:  		"read",
		Aliases: 	[]string{"r"},
		Usage: 		"read messages",
	}
	command.Subcommands = []cli.Command{
		subCmdReadCommands(),
		subCmdReadMessages(),
		subCmdReadLatestCommands(),
		subCmdReadLatestMessages(),
	}
	return command
}

func subCmdReadCommands() cli.Command {
	return cli.Command {
		Name:        	"commands",
		Aliases: 		[]string{"cmd"},
		Description: 	"read commands from channels",
		Usage:       	"read commands",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid",
				Value:       "",
				Usage:       "read commands --deid=DEVICE_ID",
				Destination: &subcmd.FlagDeviceId,
			},
			cli.StringFlag{
				Name:        "cmdid",
				Value:       "",
				Usage:       "read commands --deid=DEVICE_ID --cmdid=COMMAND_ID",
				Destination: &subcmd.FlagCmdId,
			},
		},
		Action:      subcmd.ReadCommands,
	}
}

func subCmdReadLatestCommands() cli.Command {
	return cli.Command {
		Name:        	"commands-latest",
		Aliases: 		[]string{"cmd-latest"},
		Description: 	"read latest commands from channels",
		Usage:       	"read commands-latest",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid",
				Value:       "",
				Usage:       "read cmd-latest --deid=DEVICE_ID",
				Destination: &subcmd.FlagDeviceId,
			},
		},
		Action:      subcmd.ReadLatestCommand,
	}
}

func subCmdReadMessages() cli.Command {
	return cli.Command{
		Name:        	"messages",
		Aliases: 		[]string{"msg"},
		Description: 	"read messages from devices",
		Usage:       	"read messages",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "app-id",
				Value:       "",
				Usage:       "read messages --app-id=APP_ID",
				Destination: &subcmd.FlagAppId,
			},
			cli.StringFlag{
				Name:        "msgid",
				Value:       "",
				Usage:       "read messages --app-id=APP_ID --msgid=MSG_ID",
				Destination: &subcmd.FlagMsgId,
			},
			cli.StringFlag{
				Name:        "channel-id",
				Value:       "",
				Usage:       "read messages --channel-id=CHANNEL_ID --msgid=MSG_ID",
				Destination: &subcmd.FlagChannelId,
			},
		},
		Action:      subcmd.ReadMessages,
	}
}

func subCmdReadLatestMessages() cli.Command {
	return cli.Command{
		Name:        	"messages-latest",
		Aliases: 		[]string{"msg-latest"},
		Description: 	"read latest message from devices",
		Usage:       	"read msg-latest",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "app-id",
				Value:       "",
				Usage:       "read msg-latest --app-id=APP_ID",
				Destination: &subcmd.FlagAppId,
			},
			cli.StringFlag{
				Name:        "channel-id",
				Value:       "",
				Usage:       "read msg-latest --channel-id=CHANNEL_ID --msgid=MSG_ID",
				Destination: &subcmd.FlagChannelId,
			},
		},
		Action:      subcmd.ReadLatestMessages,
	}
}
