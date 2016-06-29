package subcmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var (
	FlagDeviceGroup string
)

func ListDeviceGroups(context *cli.Context) error {
	fmt.Println("ListDeviceGroups method " + FlagDeviceGroup)
	return nil
}
