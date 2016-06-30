package subcmd

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"github.com/urfave/cli"
)

type ConfigFile struct {
	PdexUrl 	string `json:"pdex_url"`
	AccessKey 	string `json:"access_key"`
}

var (
	h            = os.Getenv("HOME")
	confDirName  = ".pdex-cli"
	confFileName = "conf.json"
	confDir      = fmt.Sprintf("%s/%s", h, confDirName)
	confPath     = fmt.Sprintf("%s/%s", confDir, confFileName)
	FlagUrl 		string
	FlagAccessKey 	string
)

func ConfigureCommands(context *cli.Context) error {
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	urlstring := conf.PdexUrl
	keystring := conf.AccessKey

	if len(FlagUrl) != 0 {
		urlstring = FlagUrl
	}

	if len(FlagAccessKey) != 0 {
		keystring = FlagAccessKey
	}

	CreateConfig()
	confs := &ConfigFile{
		PdexUrl: urlstring,
		AccessKey: keystring,
	}
	WriteConfigs(confs)
	fmt.Println("successfully configured")
	return nil
}

func ListConfigureCommands(context *cli.Context) error {
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	key  := conf.AccessKey
	replacement := substr(key,0,len(key)-4)
	fmt.Println("End point Url: ",conf.PdexUrl)
	fmt.Println("Access key   : ",strings.Replace(key,replacement , "********", 1))
	fmt.Println("Config File  : ",confPath)
	return nil
}

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

func ReadConfigs() (c *ConfigFile, err error) {
	b, err := ioutil.ReadFile(confPath)
	if err != nil {
		return
	}
	c = new(ConfigFile)
	err = json.Unmarshal(b, c)
	return
}

func WriteConfigs(c *ConfigFile) (err error) {
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

func substr(s string,pos,length int) string {
    runes:=[]rune(s)
    ln := pos+length
    if ln > len(runes) {
        ln = len(runes)
    }
    return string(runes[pos:ln])
}
