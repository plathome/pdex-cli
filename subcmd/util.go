package subcmd

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func PingCommand(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	result, _ :=  GetUtils(fmt.Sprintf("%s%s",conf.PdexUrl,"/utils/ping"))
	fmt.Println(result)
	return nil
}

func VersionCommand(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	result, _ := GetUtils(fmt.Sprintf("%s%s",conf.PdexUrl,"/utils/version"))
	fmt.Println(result)
	return nil
}

func ChangelogCommand(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	result, _ := GetUtils(fmt.Sprintf("%s%s",conf.PdexUrl,"/utils/changelog"))
	fmt.Println(result)
	return nil
}

func HmacCommand(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagDeviceId == "" {
		fmt.Println("pdex util hmac --deid DEVICE-ID")
	} else {
		HmacCommandTask(conf.PdexUrl, conf.AccessKey, FlagDeviceId)
	}
	return nil
}

func DeleteChannel(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagDeviceId == "" || FlagChannelId == "" || FlagConfirmation == "" {
		fmt.Println("pdex delete channels --deid DEVICE-ID --app-id APP-ID --confirm true/false")
	}
	if FlagChannelId != "" && FlagDeviceId != "" && FlagConfirmation != "" {
		DeleteChannelTask(conf.PdexUrl, conf.AccessKey, FlagDeviceId, FlagChannelId, FlagConfirmation)
	}
	return nil
}

func DeleteAccount(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagConfirmation == "" {
		fmt.Println("pdex delete account --confirm true/false")
	} else {
		DeleteAccountTask(conf.PdexUrl, conf.AccessKey, FlagConfirmation)
	}
	return nil
}
