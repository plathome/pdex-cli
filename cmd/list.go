package cmd

import (
	"github.com/urfave/cli"
//	"github.com/plathome/pdex-cli/subcmd"
	"../subcmd"
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
		subCmdDgTagKey(),
		subCmdDevicesTagKey(),
		subCmdAppsTagKey(),
	}
	return command
}

func subCmdDeviceGroups() cli.Command {
	return cli.Command{
		Name:        	"devicegroups",
		Aliases: 		[]string{"dg"},
		Description: 	"get devicegroups list.",
		Usage:       	"list devicegroups",
		Action:      	subcmd.ListDeviceGroups,
	}
}

func subCmdDevices() cli.Command {
	return cli.Command{
		Name:        	"devices",
		Aliases:		[]string{"de"},
		Description: 	"list devices",
		Usage:       	"list devices --deid-prefix DEVICE-ID-PREFIX",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid-prefix",
				Value:       "",
				Usage:       "list devices --deid-prefix DEVICE-ID-PREFIX",
				Destination: &subcmd.FlagDeviceGroup,
			},
		},
		Action: subcmd.ListDevices,
	}
}

func subCmdChannels() cli.Command {
	return cli.Command{
		Name:        "channels",
		Aliases: 	 []string{"ch"},
		Description: "list channels",
		Usage:       "list channels --deid DEVICE-ID",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid",
				Value:       "",
				Usage:       "list channels --deid DEVICE-ID",
				Destination: &subcmd.FlagDeviceId,
			},
		},
		Action: subcmd.ListChannels,
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

func subCmdDgTagKey() cli.Command {
	return cli.Command{
		Name:        "dg-tags",
		Description: "get tag list of devicegroup",
		Usage:       "list dg-tags --deid-prefix DEVICE-ID-PREFIX",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid-prefix",
				Value:       "",
				Usage:       "list dg-tags --deid-prefix DEVICE-ID-PREFIX",
				Destination: &subcmd.FlagDeviceGroup,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "list dg-tags --deid-prefix DEVICE-ID-PREFIX --key KEY",
				Destination: &subcmd.FlagKey,
			},
		},
		Action:      subcmd.ListDgTagKey,
	}
}

func subCmdDevicesTagKey() cli.Command {
	return cli.Command{
		Name:        "de-tags",
		Description: "get tag list of device",
		Usage:       "list de-tags --deid DEVICE-ID",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid",
				Value:       "",
				Usage:       "list de-tags --deid DEVICE-ID",
				Destination: &subcmd.FlagDeviceId,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "list de-tags --deid DEVICE-ID --key KEY",
				Destination: &subcmd.FlagKey,
			},
		},
		Action:      subcmd.ListDeviceTagKey,
	}
}

func subCmdAppsTagKey() cli.Command {
	return cli.Command{
		Name:        "app-tags",
		Description: "get tag list of application",
		Usage:       "list app-tags --app-id APP-ID",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "app-id",
				Value:       "",
				Usage:       "list app-tags --app-id APP-ID",
				Destination: &subcmd.FlagAppId,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "list app-tags --app-id APP-ID --key KEY",
				Destination: &subcmd.FlagKey,
			},
		},
		Action:      subcmd.ListAppTagKey,
	}
}
