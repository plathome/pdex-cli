package cmd

import (
	"github.com/urfave/cli"
	"github.com/plathome/pdex-cli/cmd"
//	"../cmd"
)

func ListCmd() cli.Command {
	command := cli.Command{
		Name:  "list",
		Aliases: []string{"ls"},
		Usage: "list up desired resource",
	}
	command.Subcommands = []cli.Command{
		subCmdDeviceGroups(),
		subCmdDevices(),
		subCmdChannels(),
		subCmdApps(),
	}
	return command
}

func subCmdDeviceGroups() cli.Command {
	return cli.Command{
		Name:        "devicegroups",
		Aliases: []string{"dg"},
		Description: "get devicegroups list.",
		Usage:       "list devicegroups",
		Action:      subcmd.ListDeviceGroups,
	}
}

func subCmdDevices() cli.Command {
	return cli.Command{
		Name:        "devices",
		Aliases: []string{"de"},
		Description: "get devices list.",
		Usage:       "list devices --deid-prefix",
		Action:      subcmd.ListDevices,
	}
}

func subCmdChannels() cli.Command {
	return cli.Command{
		Name:        "channels",
		Aliases: []string{"ch"},
		Description: "get channels list.",
		Usage:       "list channels",
		Action:      subcmd.ListChannels,
	}
}

func subCmdApps() cli.Command {
	return cli.Command{
		Name:        "apps",
		Description: "get applications list.",
		Usage:       "list apps",
		Action:      subcmd.ListApps,
	}
}
