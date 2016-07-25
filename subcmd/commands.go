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
	if FlagDeviceId == "" && FlagCmdId == "" && FlagQuery == "" {
		fmt.Println("pdex read commands --deid DEVICE_ID")
		fmt.Println("pdex read commands --deid DEVICE_ID --query FILTER_QUERY_STRING")
		fmt.Println("pdex read commands --deid DEVICE_ID --cmdid COMMAND_ID")
	}
	if FlagDeviceId != "" && FlagCmdId == "" && FlagQuery == "" {
		ReadCommandsApi(conf.PdexUrl, FlagDeviceId, conf.AccessKey, "commands")
	}
	if FlagDeviceId != "" && FlagCmdId != "" && FlagQuery == "" {
		ReadCommandsApi(conf.PdexUrl, FlagDeviceId, conf.AccessKey, "commands/" + FlagCmdId)
	}
	if FlagDeviceId != "" && FlagCmdId == "" && FlagQuery != "" {
		ReadCommandsApiFilter(conf.PdexUrl, FlagDeviceId, conf.AccessKey, "commands", FlagQuery)
	}
	return nil
}

func ReadLatestCommand(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagDeviceId == "" {
		fmt.Println("pdex read command-latest --deid DEVICE_ID")
	} else {
		ReadCommandsApi(conf.PdexUrl, FlagDeviceId, conf.AccessKey, "commands/latest")
	}
	return nil
}
