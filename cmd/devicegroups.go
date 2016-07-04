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
