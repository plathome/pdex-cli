package subcmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var (
	FlagChannelId string
)

func SendCommands(context *cli.Context) error {
	fmt.Println("SendCommands method --channelid " + FlagChannelId)
	return nil
}

func ReadCommands(context *cli.Context) error {
	fmt.Println("ReadCommands method")
	return nil
}
