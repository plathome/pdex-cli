package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/plathome/pdex-cli/cmd"
//	"./cmd"
)

func main() {
	app := cli.NewApp()
	app.Usage = "The cli tool for pd-exchange"
	app.Version = "0.1.8"
	app.Commands = []cli.Command {
		cmd.ListCmd(),
		cmd.SendCmd(),
		cmd.ReadCmd(),
		cmd.ConfigureCmd(),
		cmd.UtilsCmd(),
		cmd.ShowCmd(),
		cmd.CreateCmd(),
		cmd.UpdateCmd(),
		cmd.DeleteCmd(),
	}
	app.Run(os.Args)
}

