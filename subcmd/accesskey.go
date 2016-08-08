package subcmd

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func ShowAccessKey(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Println(os.Stderr, "error in the CreateUser context \n")
		os.Exit(1)
	}
	if FlagUsername == "" || FlagPassword == "" {
		fmt.Println("show access-key --username USERNAME --password PASSWORD")
		return nil
	} else {
		UserAccessKeyApi(fmt.Sprintf("%s/%s",conf.PdexUrl,"auth/token"), FlagUsername, FlagPassword)
	}
	return nil
}
