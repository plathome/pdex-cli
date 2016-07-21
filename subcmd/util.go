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

func DeleteDgTagKey(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the Delete DeviceGroup tag context \n")
		os.Exit(1)
	}
	if FlagDeviceGroup != "" && FlagKey != "" {
		parameters 		:=	[]string{""}
		values 			:=	[]string{""}
		DeleteApi(fmt.Sprintf("%s/%s/%s/%s/%s", conf.PdexUrl, "devicegroups", FlagDeviceGroup, "tags", FlagKey), conf.AccessKey, parameters, values, "DELETE")
	} else {
		fmt.Println("pdex delete dg-tags --deid-prefix DEVICE-ID-PREFIX --key KEY")
		return nil
	}
	return nil
}

func DeleteDeviceTagKey(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the Delete Device Tag context \n")
		os.Exit(1)
	}
	if FlagDeviceId != "" && FlagKey != "" {
		DeleteDeviceTagApi(conf.PdexUrl, FlagDeviceId, conf.AccessKey, FlagKey)
	} else {
		fmt.Println("pdex delete de-tags --deid DEVICE-ID --key KEY")
		return nil
	}
	return nil
}

func DeleteApppTagKey(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the Delete Application Tag context \n")
		os.Exit(1)
	}
	if FlagAppId != "" && FlagKey != "" {
		DeleteAppTagApi(conf.PdexUrl, conf.AccessKey, FlagAppId,  FlagKey)
	} else {
		fmt.Println("pdex delete app-tags --app-id APP-ID --key KEY")
		return nil
	}
	return nil
}

