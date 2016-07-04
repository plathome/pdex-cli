package cmd

import (
	"github.com/urfave/cli"
	"github.com/plathome/pdex-cli/cmd"
)

func UtilsCmd() cli.Command {
	command := cli.Command {
		Name:  		"util",
		Aliases: 	[]string{"u"},
		Usage: 		"util ping,version,changelog",
	}
	command.Subcommands = []cli.Command{
		subCmdPing(),
		subCmdVersion(),
		subCmdChangelog(),
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
