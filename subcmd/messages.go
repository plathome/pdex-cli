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
	if FlagAppId == "" && FlagMsgId == "" && FlagChannelId == "" && FlagQuery == "" {
		fmt.Println("pdex read messages --app-id APP_ID")
		fmt.Println("pdex read messages --app-id APP_ID --query QUERY_FILTER")
		fmt.Println("pdex read messages --app-id APP_ID --msgid MSG_ID")
		fmt.Println("pdex read messages --channel-id CHANNEL_ID --app-id APP_ID")
		fmt.Println("pdex read messages --channel-id CHANNEL_ID --app-id APP_ID --query QUERY_FILTER")
		fmt.Println("pdex read messages --channel-id CHANNEL_ID --app-id APP_ID --msgid MSG_ID")
	}
	if FlagAppId != "" && FlagMsgId == "" && FlagChannelId == "" && FlagQuery == "" {
		apptoken := GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		ReadLatestMessage(fmt.Sprintf("%s/%s/%s/%s",conf.PdexUrl,"apps",FlagAppId,"messages"),apptoken)
	}
	if FlagAppId != "" && FlagMsgId == "" && FlagChannelId == "" && FlagQuery != "" {
		apptoken := GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		ReadFilteredMessage(fmt.Sprintf("%s/%s/%s/%s",conf.PdexUrl,"apps",FlagAppId,"messages"),apptoken, FlagQuery)
	}
	if FlagAppId != "" && FlagMsgId != "" && FlagChannelId == "" && FlagQuery == "" {
		apptoken := GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		ReadLatestMessage(fmt.Sprintf("%s/%s/%s/%s/%s",conf.PdexUrl,"apps",FlagAppId,"messages",FlagMsgId),apptoken)
	}
	if FlagAppId != "" && FlagMsgId == "" && FlagChannelId != "" && FlagQuery == "" {
		apptoken := GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		ReadLatestMessage(fmt.Sprintf("%s/%s/%s/%s",conf.PdexUrl,"channels",FlagChannelId,"messages"),apptoken)
	}
	if FlagAppId != "" && FlagMsgId == "" && FlagChannelId != "" && FlagQuery != "" {
		apptoken := GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		ReadFilteredMessage(fmt.Sprintf("%s/%s/%s/%s",conf.PdexUrl,"channels",FlagChannelId,"messages"),apptoken, FlagQuery)
	}
	if FlagAppId != "" && FlagMsgId != "" && FlagChannelId != "" && FlagQuery == "" {
		apptoken := GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		ReadLatestMessage(fmt.Sprintf("%s/%s/%s/%s/%s",conf.PdexUrl,"channels",FlagChannelId,"messages",FlagMsgId),apptoken)
	}

	return nil
}

func ReadLatestMessages(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagAppId == "" && FlagChannelId == "" {
		fmt.Println("pdex read msg-latest --app-id APP_ID")
		fmt.Println("pdex read msg-latest --channel-id CHANNEL_ID --app-id APP_ID")
	}
	if FlagAppId != "" && FlagChannelId == "" {
		apptoken := GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		ReadLatestMessage(fmt.Sprintf("%s/%s/%s/%s",conf.PdexUrl,"apps",FlagAppId,"messages/latest"),apptoken)
	}
	if FlagAppId != "" && FlagChannelId != "" {
		apptoken := GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		ReadLatestMessage(fmt.Sprintf("%s/%s/%s/%s",conf.PdexUrl,"channels",FlagChannelId,"messages/latest"),apptoken)
	}
	return nil
}
