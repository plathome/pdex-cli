package cmd

import (
	"bytes"
	"github.com/urfave/cli"
	"github.com/plathome/pdex-cli/subcmd"
//	"../subcmd"
)

func UtilsCmd() cli.Command {
	command := cli.Command {
		Name:  		"util",
		Aliases: 	[]string{"u"},
		Usage: 		"util ping, version, changelog, hmac",
	}
	command.Subcommands = []cli.Command{
		subCmdPing(),
		subCmdVersion(),
		subCmdChangelog(),
		subCmdHmac(),
	}
	return command
}

func subCmdPing() cli.Command {
	return cli.Command{
		Name:        "ping",
		Description: "send commands to devices",
		Usage:       "util ping",
		Action:      subcmd.PingCommand,
	}
}

func subCmdVersion() cli.Command {
	return cli.Command{
		Name:        "version",
		Description: "PDExchange version",
		Usage:       "util version",
		Action:      subcmd.VersionCommand,
	}
}

func subCmdChangelog() cli.Command {
	return cli.Command{
		Name:        "changelog",
		Description: "PDExchange changelog",
		Usage:       "util changelog",
		Action:      subcmd.ChangelogCommand,
	}
}

func subCmdHmac() cli.Command {
	return cli.Command{
		Name:        "hmac",
		Description: "Digestkey generator",
		Usage:       "util hmac --deid DEVICE-ID",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid",
				Value:       "",
				Usage:       "util hmac --deid DEVICE-ID",
				Destination: &subcmd.FlagDeviceId,
			},
		},
		Action:      subcmd.HmacCommand,
	}
}

func StringConcat(array []string) string {
	var buff bytes.Buffer
	for _, element := range array {
		buff.WriteString(element)
	}
	return buff.String()
}
