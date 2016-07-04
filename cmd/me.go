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

func ShowMyself(context *cli.Context) error {
	SetActingProfile()
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	ShowMe(conf.PdexUrl,conf.AccessKey)
	return nil
}

func ShowMe(urlstr string, accesstoken string) {
	v := url.Values{}
	s := v.Encode()
	req, err := http.NewRequest("GET", urlstr + "/me", strings.NewReader(s))
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
