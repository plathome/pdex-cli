package subcmd

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func SendCommands(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagAppId == "" || FlagChannelId == "" {
		fmt.Println("pdex send commands --channel-id CHANNEL_ID --app-id APP_ID")
	}
	if FlagAppId != "" && FlagChannelId != "" {
		ChannelSendCommand(conf.PdexUrl, conf.AccessKey, FlagChannelId, FlagAppId, context.Args().First())
	}
	return nil
}

func ReadCommands(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagDeviceId == "" && FlagCmdId == "" {
		fmt.Println("pdex read commands --deid DEVICE_ID")
		fmt.Println("pdex read commands --deid DEVICE_ID --cmdid COMMAND_ID")
	}
	if FlagDeviceId != "" && FlagCmdId == "" {
		ReadCommandsApi(conf.PdexUrl, FlagDeviceId, conf.AccessKey, "commands")
	}
	if FlagDeviceId != "" && FlagCmdId != "" {
		ReadCommandsApi(conf.PdexUrl, FlagDeviceId, conf.AccessKey, "commands/" + FlagCmdId)
	}
	return nil
}
