package subcmd

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func ShowMyself(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	ShowMeApi(conf.PdexUrl,conf.AccessKey)
	return nil
}
