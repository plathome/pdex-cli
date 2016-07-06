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
	GetUtils(conf.PdexUrl,"/utils/ping")
	return nil
}

func VersionCommand(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	GetUtils(conf.PdexUrl,"/utils/version")
	return nil
}

func ChangelogCommand(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	GetUtils(conf.PdexUrl,"/utils/changelog")
	return nil
}

func HmacCommand(context *cli.Context) error {
	if context.Args().First() == "" {
		fmt.Fprint(os.Stderr, "Error: Please entry the key and deid. \n")
		os.Exit(1)
	}
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	HmacDigest(conf.PdexUrl, context.Args().First(), context.Args().Get(1))
	return nil
}
