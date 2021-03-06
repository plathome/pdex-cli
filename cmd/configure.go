package cmd

import (
	"github.com/urfave/cli"
	"github.com/plathome/pdex-cli/subcmd"
//	"../subcmd"
)

func ConfigureCmd() cli.Command {
	command := cli.Command {
		Name:  		"configure",
		Aliases: 	[]string{"c"},
		Usage: 		"configure profiles",
	}
	command.Subcommands = []cli.Command{
		subCmdConfigure(),
		subCmdConfigureProfile(),
		subCmdConfigureList(),
	}
	return command
}

func subCmdConfigure() cli.Command {
	return cli.Command {
		Name:        	"set",
		Description: 	"configuration profiles setup",
		Usage:       	"configure set --url API_END_POINT --access-key ACCESS_KEY",
		Flags: []cli.Flag {
			cli.StringFlag {
				Name:        "url",
				Value:       "",
				Usage:       "configure set --url API_END_POINT",
				Destination: &subcmd.FlagUrl,
			},
			cli.StringFlag {
				Name:        "access-key",
				Value:       "",
				Usage:       "configure set --access-key ACCESS_KEY",
				Destination: &subcmd.FlagAccessKey,
			},
			cli.StringFlag {
				Name:        "username",
				Value:       "",
				Usage:       "configure set --url API_END_POINT --username USER_NAME --password PASS_WORD",
				Destination: &subcmd.FlagUsername,
			},
			cli.StringFlag {
				Name:        "password",
				Value:       "",
				Usage:       "configure set --url API_END_POINT --username USER_NAME --password PASS_WORD",
				Destination: &subcmd.FlagPassword,
			},
		},
		Action:      subcmd.ConfigureCommands,
	}
}

func subCmdConfigureProfile() cli.Command {
	return cli.Command {
		Name:        "profile",
		Description: "configuration profiles setup for profile",
		Usage:       "configure profile --name PROFILE_NAME --url API_END_POINT --access-key ACCESS_KEY",
		Flags: []cli.Flag {
			cli.StringFlag {
				Name:        "name",
				Value:       "",
				Usage:       "configure profile --name PROFILE_NAME --url API_END_POINT --access-key ACCESS_KEY",
				Destination: &subcmd.FlagProfileName,
			},
			cli.StringFlag{
				Name:        "url",
				Value:       "",
				Usage:       "configure profile --name PROFILE_NAME --url API_END_POINT",
				Destination: &subcmd.FlagUrl,
			},
			cli.StringFlag{
				Name:        "access-key",
				Value:       "",
				Usage:       "configure profile --name PROFILE_NAME --url API_END_POINT --access-key ACCESS_KEY",
				Destination: &subcmd.FlagAccessKey,
			},
			cli.StringFlag {
				Name:        "username",
				Value:       "",
				Usage:       "configure profile --name PROFILE_NAME --url API_END_POINT --username USER_NAME --password PASS_WORD",
				Destination: &subcmd.FlagUsername,
			},
			cli.StringFlag {
				Name:        "password",
				Value:       "",
				Usage:       "configure profile --name PROFILE_NAME --url API_END_POINT --username USER_NAME --password PASS_WORD",
				Destination: &subcmd.FlagPassword,
			},
		},
		Action:      subcmd.ConfigureCommandsProfile,
	}
}

func subCmdConfigureList() cli.Command {
	return cli.Command {
		Name:        	"list",
		Aliases: 		[]string{"ls"},
		Description: 	"list configuration profiles",
		Usage:       	"configure list",
		Action:      	subcmd.ListConfigureCommands,
	}
}
