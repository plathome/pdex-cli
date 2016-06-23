package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
	"gopkg.in/resty.v0"
	"github.com/plathome/pdex-cli/conf"
)

func main() {
	app := cli.NewApp()
	app.Name = "pdcli"
	app.Usage = "the cli for pd-exchange"
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
						conf := &configuration.Config{
							URL: c.Args().First(),
						}
						configuration.WriteConfig(conf)
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
					conf, err := configuration.ReadConfig()
					if err != nil {
						fmt.Fprint(os.Stderr, "Error: Failed read config. \n Require [config] command first. \n")
						os.Exit(1)
					}
					checkPing()
					pd := &configuration.Pdex{C: c}
					return nil
					},
				},
				{
					Name:    "version",
					Aliases: []string{"v"},
					Usage:   "util version",
					Action: func(c *cli.Context) error {
						checkVersion()
						return nil
					},
				},
				{
					Name:    "changelog",
					Aliases: []string{"c"},
					Usage:   "util changelog",
					Action: func(c *cli.Context) error {
						checkChangeLog()
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

func checkPing() {
	resp, err := resty.R().Get("http://localhost:9292/api/v1/utils/ping")
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
	fmt.Printf("Response Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v\n", resp)
}

func checkVersion() {
	resp, err := resty.R().Get("http://localhost:9292/api/v1/utils/version")
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
	fmt.Printf("Response Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v\n", resp)
}

func checkChangeLog() {
	resp, err := resty.R().Get("http://localhost:9292/api/v1/utils/changelog")
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
	fmt.Printf("Response Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v\n", resp)
}
