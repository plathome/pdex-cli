package subcmd

import (
	"fmt"
	"os"
	"strings"
	"github.com/urfave/cli"
)

func ListChannels(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagDeviceId == "" {
		fmt.Println("list channels --deid DEVICE-ID")
		return nil
	} else {
		ListChannelApi(conf.PdexUrl, conf.AccessKey, FlagDeviceId)
	}
	return nil
}

func ShowChannel(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if AtLeastTwo(FlagDeviceId == "", FlagChannelId == "", FlagAppId == "") {
		fmt.Println("pdex show channels --deid DEVICE-ID --channel-id CHANNEL-ID")
		fmt.Println("pdex show channels --app-id APP-ID --channel-id CHANNEL-ID")
	}
	if FlagDeviceId != "" &&  FlagChannelId != "" && FlagAppId == "" {
		digest_key, _ := ShowChannelApi(conf.PdexUrl, conf.AccessKey, FlagDeviceId, FlagChannelId)
		fmt.Println(digest_key)
	}
	if FlagAppId != "" && FlagChannelId != "" && FlagDeviceId == "" {
		apptoken 	:= GetAppToken(conf.PdexUrl, conf.AccessKey, FlagAppId)
		initial_str := ListApiReturn(fmt.Sprintf("%s/%s/%s", conf.PdexUrl, "channels", FlagChannelId), apptoken)

		final_str	:= ""

		if strings.Contains(initial_str, "error") == true {
			final_str = initial_str
		} else {
			replace_str := fmt.Sprintf(",%s", "\"app_token\":\"" + apptoken +"\"}")
			final_str   = strings.Replace(initial_str,"}",replace_str,-1)
		}

		fmt.Println(final_str)
	}
	return nil
}

func CreateChannel(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagAppId == "" || FlagDeviceId == "" {
		fmt.Println("pdex create channels --deid DEVICE-ID --app-id APP-ID")
	}
	if FlagAppId != "" && FlagDeviceId != "" {
		channel_created, _ := ChannelCreateApi(conf.PdexUrl, conf.AccessKey, FlagDeviceId, FlagAppId)
		fmt.Println(channel_created)
	}
	return nil
}
