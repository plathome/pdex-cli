package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
	"gopkg.in/resty.v0"
//	"github.com/plathome/pdex-cli/conf"
	"./conf"
)

func main() {
	app := cli.NewApp()
	app.Name = "pdex-cli"
	app.Usage = "The cli tool for pd-exchange"
	app.Version = "0.0.1"
	app.Commands = []cli.Command {
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "url configuration",
			Subcommands: []cli.Command{
				{
					Name:    "url",
					Aliases: []string{"u"},
					Usage:   "set the pdex endpoint",
					Action: func(c *cli.Context) error {
						if c.Args().First() == "" {
							fmt.Fprint(os.Stderr, "Error: Please entry the url. \n")
							os.Exit(1)
						}
						configuration.CreateConfig()
						confs := &configuration.ConfigFile{
							PdexUrl: c.Args().First(),
							AccessKey: "ACCESS-KEY",
						}
						configuration.WriteConfigs(confs)
						fmt.Fprint(os.Stdout, "Success: register url config \n")
						return nil
					},
				},
			},
		},
		{
			Name:    "util",
			Aliases: []string{"n"},
			Usage:   "pdcli util",
			Subcommands: []cli.Command{
				{
					Name:    "ping",
					Aliases: []string{"p"},
					Usage:   "util ping",
					Action: func(c *cli.Context) error {
						conf, err := configuration.ReadConfigs()
						if err != nil {
							fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
							os.Exit(1)
						}
						CheckPing(conf.PdexUrl)
						return nil
					},
				},
				{
					Name:    "version",
					Aliases: []string{"v"},
					Usage:   "util version",
					Action: func(c *cli.Context) error {
						conf, err := configuration.ReadConfigs()
						if err != nil {
							fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
							os.Exit(1)
						}
						CheckVersion(conf.PdexUrl)
						return nil
					},
				},
				{
					Name:    "changelog",
					Aliases: []string{"c"},
					Usage:   "util changelog",
					Action: func(c *cli.Context) error {
						conf, err := configuration.ReadConfigs()
						if err != nil {
							fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
							os.Exit(1)
						}
						CheckChangeLog(conf.PdexUrl)
						return nil
					},
				},
			},
		},
		{
			Name:    "device",
			Aliases: []string{"d"},
			Usage:   "pdcli device",
			Subcommands: []cli.Command{
				{
					Name:    "sendmsg",
					Aliases: []string{"s"},
					Usage:   "device send",
					Action: func(c *cli.Context) error {
						if c.Args().First() == "" {
							fmt.Fprint(os.Stderr, "Error: Please entry the url. \n")
							os.Exit(1)
						}
						conf, err := configuration.ReadConfigs()
						if err != nil {
							fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
							os.Exit(1)
						}
						fmt.Println(conf.PdexUrl)
						fmt.Println(conf.AccessKey)
						//DeviceSendMessage(c.Args().First())
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

func CheckPing(url string) {
	resp, err := resty.R().Get(url + "/api/v1/utils/ping")
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
	fmt.Printf("%v\n", resp)
}

func CheckVersion(url string) {
	resp, err := resty.R().Get(url + "/api/v1/utils/version")
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
	fmt.Printf("%v\n", resp)
}

func CheckChangeLog(url string) {
	resp, err := resty.R().Get(url + "/api/v1/utils/changelog")
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
	fmt.Printf("%v\n", resp)
}

func DeviceSendMessage(conf string) {
	// resp, err := resty.R().Get("http://localhost:9292/api/v1/utils/changelog")
	// if err != nil {
	// 	fmt.Printf("\nError: %v", err)
	// }
	// fmt.Printf("%v\n", resp)
	fmt.Printf("\n%v\n",conf)
}

func ReadMessageByApp() {
	// resp, err := resty.R().Get("http://localhost:9292/api/v1/utils/changelog")
	// if err != nil {
	// 	fmt.Printf("\nError: %v", err)
	// }
	// fmt.Printf("%v\n", resp)
	fmt.Printf("%v\n","MessageByAppInfo")
}

func ReadMessageByChannel() {
	// resp, err := resty.R().Get("http://localhost:9292/api/v1/utils/changelog")
	// if err != nil {
	// 	fmt.Printf("\nError: %v", err)
	// }
	// fmt.Printf("%v\n", resp)
	fmt.Printf("%v\n","ReadMessageByChannel")
}
