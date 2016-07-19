package subcmd

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func ShowDeviceGroup(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}

	if FlagDeviceGroup == "" {
		fmt.Println("show devicegroups --deid-prefix DEVICE-ID-PREFIX")
		return nil
	} else {
		ListApi(fmt.Sprintf("%s/%s/%s", conf.PdexUrl, "devicegroups", FlagDeviceGroup), conf.AccessKey)
	}
	return nil
}

func ListDeviceGroups(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	ListApi(fmt.Sprintf("%s/%s", conf.PdexUrl, "devicegroups"), conf.AccessKey)
	return nil
}

func CreateDG(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	CreateApi(fmt.Sprintf("%s/%s", conf.PdexUrl, "devicegroups") , conf.AccessKey,  "", "")
	return nil
}

func UpdateSession(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed in reading the config file. \n")
		os.Exit(1)
	}
	UpdateSessionApi(fmt.Sprintf("%s/%s",conf.PdexUrl,"auth/token"), conf.AccessKey)
	return nil
}

func UpdatePassword(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the CreateUser context \n")
		os.Exit(1)
	}
	if FlagCurrentPassword == "" || FlagNewPassword == "" {
		fmt.Println("update password --current-password CUR-PASSWD --new-password PASSWORD")
		return nil
	} else {
		UpdatePasswordApi(fmt.Sprintf("%s/%s",conf.PdexUrl,"auth/secret"), conf.AccessKey, FlagCurrentPassword, FlagNewPassword)
	}
	return nil
}

func UpdateApp(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the CreateUser context \n")
		os.Exit(1)
	}
	if FlagAppNameSuffix == "" || FlagAppId == "" {
		fmt.Println("update apps --app-name-suffix APP-NAME-SIFFIX --app-id APPID")
		return nil
	} else {
		UpdateAppApi(fmt.Sprintf("%s/%s/%s",conf.PdexUrl,"apps",FlagAppId), conf.AccessKey, FlagAppNameSuffix)
	}
	return nil
}

func UpdateDgTag(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the CreateUser context \n")
		os.Exit(1)
	}
	if FlagDeviceGroup != "" && FlagKey != "" && FlagValue != "" {
		parameters 	:=	[]string{"value"}
		values 		:=	[]string{FlagValue}
		UpdateTagApi(fmt.Sprintf("%s/%s/%s/%s/%s",conf.PdexUrl,"devicegroups",FlagDeviceGroup,"tags",FlagKey), conf.AccessKey, parameters, values)
	} else {
		fmt.Println("update dg-tags --deid-prefix DEVICE-ID-PREFIX --key KEY --value VALUE")
		return nil
	}
	return nil
}

func UpdateDeviceTag(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the Update Device Tags context \n")
		os.Exit(1)
	}
	if FlagDeviceId != "" && FlagKey != "" && FlagValue != "" {
		UpdateDeviceTapApi(conf.PdexUrl, conf.AccessKey, FlagDeviceId, FlagKey, FlagValue)
	} else {
		fmt.Println("update dg-tags --deid DEVICE-ID --key KEY --value VALUE")
		return nil
	}
	return nil
}

func UpdateAppTag(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the Update App Tags context \n")
		os.Exit(1)
	}
	if FlagAppId != "" && FlagKey != "" && FlagValue != "" {
		UpdateAppTagApi(conf.PdexUrl, conf.AccessKey, FlagAppId, FlagKey, FlagValue)
	} else {
		fmt.Println("update app-tags --app-id APP-ID --key KEY --value VALUE")
		return nil
	}
	return nil
}

func CreateUser(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the CreateUser context \n")
		os.Exit(1)
	}
	if FlagUsername == "" || FlagPassword == "" {
		fmt.Println("create session --username USERNAME --password PASSWORD")
		return nil
	} else {
		CreateUserApi(fmt.Sprintf("%s/%s",conf.PdexUrl,"users"), conf.AccessKey, FlagUsername, FlagPassword)
	}
	return nil
}

func CreateDgTags(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the CreateDgTags Context \n")
		os.Exit(1)
	}
	if FlagKey != "" && FlagValue != "" && FlagDeviceGroup != "" {
		CreateApi(fmt.Sprintf("%s/%s/%s/%s/%s", conf.PdexUrl, "devicegroups", FlagDeviceGroup, "tags", FlagKey) , conf.AccessKey,  "value", FlagValue)
	} else {
		fmt.Println("create dg-tags --deid-prefix DEVICE-ID-PREFIX --key KEY --value VALUE")
	}
	return nil
}

func CreateDeviceTags(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the CreateDeviceTags Context \n")
		os.Exit(1)
	}
	if FlagKey != "" && FlagValue != "" && FlagDeviceId != "" {
		CreateDeviceTagsApi(conf.PdexUrl, FlagDeviceId, conf.AccessKey, FlagKey, FlagValue)
	} else {
		fmt.Println("create device-tags --deid DEVICE-ID --key KEY --value VALUE")
	}
	return nil
}

func CreateAppTags(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the CreateAppTags Context \n")
		os.Exit(1)
	}
	if FlagKey != "" && FlagValue != "" && FlagAppId != "" {
		CreateAppTagsApi(conf.PdexUrl, FlagAppId, conf.AccessKey, FlagKey, FlagValue)
	} else {
		fmt.Println("create ap-tags --app-id APP-ID --key KEY --value VALUE")
	}
	return nil
}

func CreateSession(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the CreateUser context \n")
		os.Exit(1)
	}
	if FlagUsername == "" || FlagPassword == "" {
		fmt.Println("create session --username USERNAME --password PASSWORD")
		return nil
	} else {
		CreateUserApi(fmt.Sprintf("%s/%s",conf.PdexUrl,"auth/token"), conf.AccessKey, FlagUsername, FlagPassword)
	}
	return nil
}
