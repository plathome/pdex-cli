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
	FlagDeviceGroup string
)

func ShowDeviceGroup(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}

	if FlagDeviceGroup == "" {
		fmt.Println("show devicegroups --deid-prefix API_END_POINT")
		return nil
	} else {
		ListDeviceGroupsApi(conf.PdexUrl,conf.AccessKey,fmt.Sprintf("/%s", FlagDeviceGroup))
	}
	return nil
}

func ListDeviceGroups(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	ListDeviceGroupsApi(conf.PdexUrl,conf.AccessKey,"")
	return nil
}

func ListDeviceGroupsApi(urlstr string, accesstoken string, searchstr string) {
	v := url.Values{}
	s := v.Encode()
	req, err := http.NewRequest("GET", urlstr + "/devicegroups" + searchstr, strings.NewReader(s))
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
	fmt.Printf("%v\n", string(data))
}

func CreateDG(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	CreateDeviceGroupApi(conf.PdexUrl,conf.AccessKey)
	return nil
}

func CreateDeviceGroupApi(urlstr string, accesskey string) {
	v := url.Values{}
	s := v.Encode()
	req, err := http.NewRequest("POST", urlstr + "/devicegroups", strings.NewReader(s))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	req.Header.Add("Authorization", "Bearer " + accesskey)

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

