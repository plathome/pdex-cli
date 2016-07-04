package subcmd

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func ListApps(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	ListApi(fmt.Sprintf("%s/%s", conf.PdexUrl, "apps"), conf.AccessKey)
	return nil
}

func ShowApp(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}

	if FlagAppId == "" {
		fmt.Println("show apps --app-id APP_ID")
		return nil
	} else {
		ListApi(fmt.Sprintf("%s/%s/%s", conf.PdexUrl, "apps", FlagAppId), conf.AccessKey)
	}
	return nil
}

func CreateApp(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagAppNameSuffix == "" {
		fmt.Println("create apps --app-name-suffix APP-NAME-SUFFIX")
		return nil
	} else {
		CreateApi(fmt.Sprintf("%s/%s", conf.PdexUrl, "apps"), conf.AccessKey,  "app_name_suffix", FlagAppNameSuffix)
	}
	return nil
}
