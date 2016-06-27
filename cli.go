package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
//	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
//	"github.com/plathome/pdex-cli/conf"
	"./conf"
)

// type HmacApiResponse struct {
// 	HmacValue string `json:"digest"`
// }

type Result struct {
    Digest string
}

// type HMACS struct {
//     Hmac  string
// }

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
							AccessKey: c.Args().Get(1),
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
						GetUtils(conf.PdexUrl,"/api/v1/utils/ping")
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
						GetUtils(conf.PdexUrl,"/api/v1/utils/version")
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
						GetUtils(conf.PdexUrl,"/api/v1/utils/changelog")
						return nil
					},
				},
				{
					Name:    "hmac",
					Aliases: []string{"c"},
					Usage:   "util hmac",
					Action: func(c *cli.Context) error {
						conf, err := configuration.ReadConfigs()
						if err != nil {
							fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
							os.Exit(1)
						}
						DeviceSendMessage(conf.PdexUrl + "/api/v1/utils/hmac", "e63ab20b0771", "01.8529e6.26387394")
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
						conf, err := configuration.ReadConfigs()
						if err != nil {
							fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
							os.Exit(1)
						}
						DeviceSendMessage(conf.PdexUrl + "/api/v1/utils/hmac", "e63ab20b0771", "01.8529e6.26387394")
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

func DeviceSendMessage(urlstr string, secretkey string, deid string) {
	PostRequest(urlstr, []string{"key","message"} , []string{secretkey,deid} )
}

func ReadMessageByApp() {
	fmt.Printf("%v\n","MessageByAppInfo")
}

func ReadMessageByChannel() {
	fmt.Printf("%v\n","ReadMessageByChannel")
}

func PostRequest(link string, parameters []string, values []string) (body string, err error) {
   data := url.Values{}
   for i := range parameters {
      data.Set(parameters[i], values[i])
   }
   resp, err := http.PostForm(link, data)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   ret, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      return "", err
   }
   fmt.Println(string(ret))
   return string(ret), nil
}

func GetUtils(urlstr string, utils string) {
	resp, err := http.Get(urlstr + utils)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(htmlData))
}

