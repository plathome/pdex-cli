package cmd

import (
	"github.com/urfave/cli"
//	"github.com/plathome/pdex-cli/subcmd"
	"../subcmd"
)

func UpdateCmd() cli.Command {
	command := cli.Command {
		Name:  		"update",
		Aliases:	[]string{"up"},
		Usage: 		"update session",
	}
	command.Subcommands = []cli.Command{
		subCmdUpdateSession(),
		subCmdUpdatePassword(),
		subCmdUpdateApp(),
	}
	return command
}

func subCmdUpdateSession() cli.Command {
	return cli.Command {
		Name:        	"session",
		Aliases: 		[]string{"s"},
		Description: 	"update session",
		Usage:       	"update session",
		Action:      	subcmd.UpdateSession,
	}
}

func subCmdUpdatePassword() cli.Command {
	return cli.Command {
		Name:			"password",
		Aliases:		[]string{"pwd"},
		Description:	"update pwd",
		Usage:			"update pwd --current-password CURRENT-PASSWORD --new-password PASSWORD",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "current-password",
				Value:       "",
				Usage:       "update pwd --current-password CURRENT-PASSWORD --new-password PASSWORD",
				Destination: &subcmd.FlagCurrentPassword,
			},
			cli.StringFlag{
				Name:        "new-password",
				Value:       "",
				Usage:       "update pwd --current-password CURRENT-PASSWORD --new-password PASSWORD",
				Destination: &subcmd.FlagNewPassword,
			},
		},
		Action:			subcmd.UpdatePassword,
	}
}

func subCmdUpdateApp() cli.Command {
	return cli.Command {
		Name:			"apps",
		Description:	"update apps",
		Usage:			"update apps --app-name-suffix APP-NAME-SIFFIX --app-id APPID",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "app-name-suffix",
				Value:       "",
				Usage:       "update apps --app-name-suffix APP-NAME-SIFFIX --app-id APPID",
				Destination: &subcmd.FlagAppNameSuffix,
			},
			cli.StringFlag{
				Name:        "app-id",
				Value:       "",
				Usage:       "update apps --app-name-suffix APP-NAME-SIFFIX --app-id APPID",
				Destination: &subcmd.FlagAppId,
			},
		},
		Action:			subcmd.UpdateApp,
	}
}
