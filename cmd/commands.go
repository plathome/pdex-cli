package subcmd

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"
	"net/http"
	"net/url"
	"github.com/urfave/cli"
)

type DigestData struct {
	Digest string `json:"digest"`
}

var (
	FlagChannelId string
	FlagCmdId string
)

func SendCommands(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}

	if FlagAppId == "" || FlagChannelId == "" {
		fmt.Println("pdex send commands --channel-id CHANNEL_ID --app-id APP_ID")
	}
	if FlagAppId != "" && FlagChannelId != "" {
		ChannelSendCommand(conf.PdexUrl, conf.AccessKey, FlagChannelId, FlagAppId, context.Args().First())
	}
	return nil
}

func ReadCommands(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if FlagDeviceId == "" && FlagCmdId == "" {
		fmt.Println("pdex read commands --deid DEVICE_ID")
		fmt.Println("pdex read commands --deid DEVICE_ID --cmdid COMMAND_ID")
	}
	if FlagDeviceId != "" && FlagCmdId == "" {
		ReadCommandsApi(conf.PdexUrl, FlagDeviceId, conf.AccessKey, "commands")
	}
	if FlagDeviceId != "" && FlagCmdId != "" {
		ReadCommandsApi(conf.PdexUrl, FlagDeviceId, conf.AccessKey, "commands/" + FlagCmdId)
	}
	return nil
}

func ChannelSendCommand(urlstr string, accesstoken string, channelid string, appid string, command string) {
	apptoken := GetAppToken(urlstr, accesstoken, appid)
	v := url.Values{}
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

func ReadCommandsApi(urlstr string, deid string, accesstoken string, commandstr string) {
	dgparts := strings.Split(deid,".")
	devicegroup  := dgparts[0] + "." + dgparts[1]
	secretkey := GetSecretKey(urlstr + "/devicegroups/" + devicegroup , accesstoken)
	digestdata, err := Hmac(urlstr, []string{"key","message"} , []string{secretkey, deid} )
	jd := new(DigestData)
    err = json.Unmarshal([]byte(digestdata), &jd)
    if err != nil {
        fmt.Println(err)
    }
	v := url.Values{}
	s := v.Encode()
	req, err := http.NewRequest("GET", urlstr + "/devices/" + deid + "/" + commandstr, strings.NewReader(s))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	req.Header.Add("Authorization", "Bearer " + jd.Digest)
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

func SetActingProfile() {
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	confPath = fmt.Sprintf("%s/%s", confDir, conf.ConfigFile)
}
