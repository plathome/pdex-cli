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

var (
	FlagDeviceId string
)

func SendMessages(context *cli.Context) error {
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	DeviceSendMessage(conf.PdexUrl, conf.AccessKey, FlagDeviceId, context.Args().First())
	return nil
}

func ReadMessages(context *cli.Context) error {
	fmt.Println("ReadMessages method")
	return nil
}

func DeviceSendMessage(link string, accesstoken string, deid string, message string) (body string, err error) {
   dgparts := strings.Split(deid,".")
   devicegroup  := dgparts[0] + "." + dgparts[1]
   secretkey := GetSecretKey(link + "/devicegroups/" + devicegroup , accesstoken)
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

func GetSecretKey(urlstr string, accesstoken string) (body string) {
	v := url.Values{}
	s := v.Encode()
	req, err := http.NewRequest("GET", urlstr, strings.NewReader(s))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	req.Header.Add("Authorization", "Bearer " + accesstoken)
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
	var secretkey map[string]interface{}
	json.Unmarshal(data, &secretkey)
	key, _ := secretkey["secret_key"].(string)
	return key
}
