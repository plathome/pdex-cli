package subcmd

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"net/http"
	"net/url"
	"github.com/urfave/cli"
)

var (
	FlagChannelId string
	FlagCmdId string
)

func SendCommands(context *cli.Context) error {
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	ChannelSendCommand(conf.PdexUrl, conf.AccessKey, FlagChannelId, FlagAppId, context.Args().First())
	return nil
}

func ReadCommands(context *cli.Context) error {
	fmt.Println("ReadCommands method")
	return nil
}

func ChannelSendCommand(urlstr string, accesstoken string, channelid string, appid string, command string) {
	apptoken := GetAppToken(urlstr, accesstoken, appid)
	v := url.Values{}
	v.Add("content_type","text/plain")
	v.Add("payload",command)
	s := v.Encode()
	req, err := http.NewRequest("POST", urlstr + "/channels/" + channelid + "/command", strings.NewReader(s))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	req.Header.Add("Authorization", "Bearer " + apptoken)

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
