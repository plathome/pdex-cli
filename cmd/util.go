package subcmd

import (
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/urfave/cli"
)

func PingCommand(context *cli.Context) error {
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	GetUtils(conf.PdexUrl,"/utils/ping")
	return nil
}

func VersionCommand(context *cli.Context) error {
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	GetUtils(conf.PdexUrl,"/utils/version")
	return nil
}

func ChangelogCommand(context *cli.Context) error {
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	GetUtils(conf.PdexUrl,"/utils/changelog")
	return nil
}

func HmacCommand(context *cli.Context) error {
	if context.Args().First() == "" {
		fmt.Fprint(os.Stderr, "Error: Please entry the key and deid. \n")
		os.Exit(1)
	}
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	HmacDigest(conf.PdexUrl, context.Args().First(), context.Args().Get(1))
	return nil
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
