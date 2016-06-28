package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
	"github.com/plathome/pdex-cli/conf"
	"strings"
)

type DigestString struct {
	digest string
}

func main() {
	app := cli.NewApp()
	app.Name = "pdex-cli"
	app.Usage = "The cli tool for pd-exchange"
	app.Version = "0.0.8"
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
						GetUtils(conf.PdexUrl,"/utils/ping")
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
						GetUtils(conf.PdexUrl,"/utils/version")
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
						GetUtils(conf.PdexUrl,"/utils/changelog")
						return nil
					},
				},
				{
					Name:    "hmac",
					Aliases: []string{"c"},
					Usage:   "util hmac",
					Action: func(c *cli.Context) error {
						if c.Args().First() == "" {
							fmt.Fprint(os.Stderr, "Error: Please entry the key and deid. \n")
							os.Exit(1)
						}
						conf, err := configuration.ReadConfigs()
						if err != nil {
							fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
							os.Exit(1)
						}
						HmacDigest(conf.PdexUrl, c.Args().First(), c.Args().Get(1))
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
							fmt.Fprint(os.Stderr, "Error: Please entry the key, deid and message \n")
							os.Exit(1)
						}
						conf, err := configuration.ReadConfigs()
						if err != nil {
							fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
							os.Exit(1)
						}
						DeviceSendMessage(conf.PdexUrl, c.Args().First(), c.Args().Get(1), c.Args().Get(2))
						return nil
					},
				},
			},
		},
	}
	app.Run(os.Args)
}

func HmacDigest(urlstr string, secretkey string, deid string) {
	Hmac(urlstr, []string{"key","message"} , []string{secretkey,deid} )
}

func Hmac(link string, parameters []string, values []string) (body string, err error) {
   data := url.Values{}
   for i := range parameters {
      data.Set(parameters[i], values[i])
   }
   resp, err := http.PostForm(link + "/utils/hmac", data)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   htmlData, err := ioutil.ReadAll(resp.Body)
   if err != nil {
		fmt.Println(err)
		os.Exit(1)
    }
	fmt.Println(string(htmlData))
    return string(htmlData), nil
}

func DeviceSendMessage(link string, secretkey string, deid string, message string) (body string, err error) {
   parameters := []string{"key","message"}
   values := []string{secretkey,deid}
   data := url.Values{}
   for i := range parameters {
      data.Set(parameters[i], values[i])
   }
   resp, err := http.PostForm(link + "/utils/hmac", data)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   b, _ := ioutil.ReadAll(resp.Body)

   var data2 map[string]interface{}
   json.Unmarshal(b, &data2)

   digest, _ := data2["digest"].(string)
   SendMessage(link, deid, digest, message)
   return string(b), nil
}

func SendMessage(urlStr string, deid string, digestkey string, message string) {
	v := url.Values{}
	v.Add("msg", message)
	s := v.Encode()
	req, err := http.NewRequest("POST", urlStr + "/devices/" + deid + "/message", strings.NewReader(s))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	req.Header.Add("Authorization", "Bearer " + digestkey)

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("http.Do() error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("%v\n", string(data))
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

