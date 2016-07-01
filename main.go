package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/plathome/pdex-cli/subcmd"
//	"./subcmd"
)

func main() {
	app := cli.NewApp()
	app.Usage = "The cli tool for pd-exchange"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		cmd.ListCmd(),
		cmd.SendCmd(),
		cmd.ReadCmd(),
		cmd.ConfigureCmd(),
		cmd.UtilsCmd(),
	}
	app.Run(os.Args)
}

