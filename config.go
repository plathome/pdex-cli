package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	URL string `json:"url"`
}

var (
	h            = os.Getenv("HOME")
	confDirName  = ".pdex_cli"
	confFileName = "conf.json"
	confDir      = fmt.Sprintf("%s/%s", h, confDirName)
	confPath     = fmt.Sprintf("%s/%s", confDir, confFileName)
)

func ExistConfig() bool {
	_, err := os.Stat(confPath)
	return err == nil
}

func CreateConfig() (err error) {
	err = os.MkdirAll(confDir, 0755)
	if err != nil {
		return
	}

	_, err = os.Create(confPath)
	return
}

func ReadConfig() (c *Config, err error) {
	b, err := ioutil.ReadFile(confPath)
	if err != nil {
		return
	}

	c = new(Config)
	err = json.Unmarshal(b, c)
	return
}

func WriteConfig(c *Config) (err error) {
	b, err := json.Marshal(c)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(confPath, b, 0755)
	return
}

func RemoveConfig() (err error) {
	err = os.RemoveAll(confPath)
	return
}
