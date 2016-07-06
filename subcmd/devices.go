package subcmd

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func ListDevices(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagDeviceGroup == "" {
		fmt.Println("list devices --deid-prefix DEVICE-ID-PREFIX")
		return nil
	} else {
		ListApi(fmt.Sprintf("%s/%s/%s/%s",conf.PdexUrl,"devicegroups",FlagDeviceGroup,"devices"), conf.AccessKey)
	}
	return nil
}

func CreateDevice(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagDeviceGroup == "" {
		fmt.Println("create devices --deid-prefix DEVICE-ID-PREFIX")
		return nil
	} else {
		CreateApi(fmt.Sprintf("%s/%s/%s/%s", conf.PdexUrl, "devicegroups", FlagDeviceGroup, "new"), conf.AccessKey,  "", "")
	}
	return nil
}
