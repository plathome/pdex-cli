package subcmd

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func SendMessages(context *cli.Context) error {
	if FlagDeviceId == "" {
		fmt.Println("send messages --deid DEVICE_ID")
		return nil
	} else {
		SetActingProfile()
		conf, err := ReadConfigs()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
			os.Exit(1)
		}
		DeviceSendMessage(conf.PdexUrl, conf.AccessKey, FlagDeviceId, context.Args().First())
	}
	return nil
}

func ReadMessages(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagAppId == "" && FlagMsgId == "" && FlagChannelId == "" {
		fmt.Println("pdex read messages --app-id APP_ID")
		fmt.Println("pdex read messages --app-id APP_ID --msgid MSG_ID")
		fmt.Println("pdex read messages --channel-id CHANNEL_ID --app-id APP_ID")
		fmt.Println("pdex read messages --channel-id CHANNEL_ID --app-id APP_ID --msgid MSG_ID")
	}
	if FlagAppId != "" && FlagMsgId == "" && FlagChannelId == "" {
		apptoken := GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		ReadMessage("apps", conf.PdexUrl, FlagAppId, apptoken)
	}
	if FlagAppId != "" && FlagMsgId != "" && FlagChannelId == "" {
		apptoken := GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		ReadSingleMessage("apps", conf.PdexUrl, FlagAppId, apptoken, FlagMsgId)
	}

	if FlagAppId != "" && FlagMsgId == "" && FlagChannelId != "" {
		apptoken := GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		ReadMessage("channels", conf.PdexUrl, FlagChannelId, apptoken)
	}
	if FlagAppId != "" && FlagMsgId != "" && FlagChannelId != "" {
		apptoken := GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		ReadSingleMessage("channels", conf.PdexUrl, FlagChannelId, apptoken, FlagMsgId)
	}

	return nil
}
