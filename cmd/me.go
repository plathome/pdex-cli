package subcmd

import (
	"fmt"
	"github.com/urfave/cli"
)

func ListMyself(context *cli.Context) error {
	fmt.Println("ListMyself method ")
	return nil
}
