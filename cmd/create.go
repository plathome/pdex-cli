package cmd

import (
	"github.com/urfave/cli"
//	"github.com/plathome/pdex-cli/subcmd"
	"../subcmd"
)

func CreateCmd() cli.Command {
	command := cli.Command {
		Name:  		"create",
		Aliases:	[]string{"cr"},
		Usage: 		"create new resource",
	}
	command.Subcommands = []cli.Command{
		subCmdCreateDG(),
		subCmdCreateApp(),
		subCmdCreateDevice(),
		subCmdCreateChannel(),
		subCmdCreateSession(),
		subCmdCreateUser(),
		subCmdCreateDgTags(),
		subCmdCreateDeviceTags(),
		subCmdCreateAppTags(),
	}
	return command
}

func subCmdCreateDG() cli.Command {
	return cli.Command {
		Name:        	"devicegroups",
		Aliases: 		[]string{"dg"},
		Description: 	"create new devicegroup resource",
		Usage:       	"create devicegroups",
		Action:      	subcmd.CreateDG,
	}
}

func subCmdCreateApp() cli.Command {
	return cli.Command{
		Name:        "apps",
		Description: "create new app",
		Usage:       "create apps --app-name-suffix APP-NAME-SUFFIX",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "app-name-suffix",
				Value:       "",
				Usage:       "create apps --app-name-suffix APP-NAME-SUFFIX",
				Destination: &subcmd.FlagAppNameSuffix,
			},
		},
		Action: subcmd.CreateApp,
	}
}

func subCmdCreateDevice() cli.Command {
	return cli.Command{
		Name:        	"devices",
		Aliases: 		[]string{"de"},
		Description: 	"create new device",
		Usage:       	"create devices --deid-prefix DEVICE-ID-SUFFIX",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid-prefix",
				Value:       "",
				Usage:       "create devices --deid-prefix DEVICE-ID-PREFIX",
				Destination: &subcmd.FlagDeviceGroup,
			},
		},
		Action: subcmd.CreateDevice,
	}
}

func subCmdCreateChannel() cli.Command {
	return cli.Command{
		Name:        	"channels",
		Aliases:		[]string{"ch"},
		Description: 	"create new channel",
		Usage:       	"create channels --deid DEVICE-ID --app-id APP-ID",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid",
				Value:       "",
				Usage:       "create devices --deid DEVICE-ID --app-id APP-ID",
				Destination: &subcmd.FlagDeviceId,
			},
			cli.StringFlag{
				Name:        "app-id",
				Value:       "",
				Usage:       "create devices --deid DEVICE-ID --app-id APP-ID",
				Destination: &subcmd.FlagAppId,
			},
		},
		Action: subcmd.CreateChannel,
	}
}

func subCmdCreateSession() cli.Command {
	return cli.Command {
		Name:        	"sessions",
		Aliases: 		[]string{"ses"},
		Description: 	"create new session",
		Usage:       	"create session  --username USERNAME --password PASSWORD",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "username",
				Value:       "",
				Usage:       "create session --username USERNAME --password PASSWORD",
				Destination: &subcmd.FlagUsername,
			},
			cli.StringFlag{
				Name:        "password",
				Value:       "",
				Usage:       "create session --username USERNAME --password PASSWORD",
				Destination: &subcmd.FlagPassword,
			},
		},
		Action:      	subcmd.CreateSession,
	}
}

func subCmdCreateUser() cli.Command {
	return cli.Command {
		Name:			"users",
		Aliases:		[]string{"usr"},
		Description:	"create new user",
		Usage:			"create user --username USERNAME --password PASSWORD",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "username",
				Value:       "",
				Usage:       "create user --username USERNAME --password PASSWORD",
				Destination: &subcmd.FlagUsername,
			},
			cli.StringFlag{
				Name:        "password",
				Value:       "",
				Usage:       "create user --username USERNAME --password PASSWORD",
				Destination: &subcmd.FlagPassword,
			},
		},
		Action:			subcmd.CreateUser,
	}
}

func subCmdCreateDgTags() cli.Command {
	return cli.Command {
		Name:			"dg-tags",
		Description:	"create new devicegroup tags",
		Usage:			"create dg-tags --deid-prefix DEVICE-ID-PREFIX --key KEY --value VALUE",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid-prefix",
				Value:       "",
				Usage:       "create dg-tags --deid-prefix DEVICE-ID-PREFIX --key KEY --value VALUE",
				Destination: &subcmd.FlagDeviceGroup,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "create dg-tags --deid-prefix DEVICE-ID-PREFIX --key KEY --value VALUE",
				Destination: &subcmd.FlagKey,
			},
			cli.StringFlag{
				Name:        "value",
				Value:       "",
				Usage:       "create dg-tags --deid-prefix DEVICE-ID-PREFIX --key KEY --value VALUE",
				Destination: &subcmd.FlagValue,
			},
		},
		Action:			subcmd.CreateDgTags,
	}
}

func subCmdCreateDeviceTags() cli.Command {
	return cli.Command {
		Name:			"de-tags",
		Description:	"create new device tags",
		Usage:			"create de-tags --deid DEVICE-ID --key KEY --value VALUE",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "deid",
				Value:       "",
				Usage:       "create device-tags --deid DEVICE-ID --key KEY --value VALUE",
				Destination: &subcmd.FlagDeviceId,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "create device-tags --deid DEVICE-ID --key KEY --value VALUE",
				Destination: &subcmd.FlagKey,
			},
			cli.StringFlag{
				Name:        "value",
				Value:       "",
				Usage:       "create device-tags --deid DEVICE-ID --key KEY --value VALUE",
				Destination: &subcmd.FlagValue,
			},
		},
		Action:			subcmd.CreateDeviceTags,
	}
}

func subCmdCreateAppTags() cli.Command {
	return cli.Command {
		Name:			"app-tags",
		Description:	"create new app tags",
		Usage:			"create app-tags --app-id APP-ID --key KEY --value VALUE",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "app-id",
				Value:       "",
				Usage:       "create ap-tags --app-id APP-ID --key KEY --value VALUE",
				Destination: &subcmd.FlagAppId,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "create app-tags --app-id APP-ID --key KEY --value VALUE",
				Destination: &subcmd.FlagKey,
			},
			cli.StringFlag{
				Name:        "value",
				Value:       "",
				Usage:       "create app-tags --app-id APP-ID --key KEY --value VALUE",
				Destination: &subcmd.FlagValue,
			},
		},
		Action:			subcmd.CreateAppTags,
	}
}
