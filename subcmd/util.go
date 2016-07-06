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
