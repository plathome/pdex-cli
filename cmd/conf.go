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
	ConfigFile 	string `json:"config_file"`
}

var (
	h            = os.Getenv("HOME")
	confDirName  = ".pdex-cli"
	confFileName = "conf.json"
	confDir      = fmt.Sprintf("%s/%s", h, confDirName)
	confPath     = fmt.Sprintf("%s/%s", confDir, confFileName)
	FlagUrl 		string
	FlagAccessKey 	string
	FlagProfileName string
)

func ConfigureCommands(context *cli.Context) error {
	urlstring  := ""
	keystring  := ""

	if FlagUrl == "" && FlagAccessKey == "" {
		fmt.Println("configure set --url API_END_POINT --accesskey ACCESS_KEY")
		fmt.Println("configure set --url API_END_POINT")
		fmt.Println("configure set --access-key ACCESS_KEY")
		return nil
	}

	if FlagUrl != "" && FlagAccessKey == "" {
		conf, err := ReadConfigs()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
			os.Exit(1)
		}
		urlstring  = conf.PdexUrl
		keystring  = conf.AccessKey

		urlstring = FlagUrl
	}

	if FlagUrl == "" && FlagAccessKey != "" {
		conf, err := ReadConfigs()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
			os.Exit(1)
		}
		urlstring  = conf.PdexUrl
		keystring  = conf.AccessKey

		keystring = FlagAccessKey
	}

	if FlagUrl != "" && FlagAccessKey != "" {
		urlstring 	= FlagUrl
		keystring 	= FlagAccessKey
	}

	CreateConfig()
	confs := &ConfigFile{
		PdexUrl: urlstring,
		AccessKey: keystring,
		ConfigFile: "conf.json",
	}
	WriteConfigs(confs)
	fmt.Println("successfully configured")
	return nil
}

func ConfigureCommandsProfile(context *cli.Context) error {
	if FlagProfileName != "" {

		conf, err := ReadConfigs()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
			os.Exit(1)
		}
		urlstring 	:= conf.PdexUrl
		keystring 	:= conf.AccessKey
		configfile 	:= FlagProfileName + ".json"
		CreateConfig()
		confs := &ConfigFile{
			PdexUrl: urlstring,
			AccessKey: keystring,
			ConfigFile: configfile,
		}
		WriteConfigs(confs)

		confPath     = fmt.Sprintf("%s/%s", confDir, FlagProfileName + ".json")
		CreateConfig()
		configs := &ConfigFile {
			PdexUrl: "",
			AccessKey: "",
			ConfigFile: FlagProfileName + ".json",
		}
		WriteConfigs(configs)
	}

	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	urlstring 	:= conf.PdexUrl
	keystring 	:= conf.AccessKey
	configfile 	:= conf.ConfigFile

	if FlagUrl == "" && FlagAccessKey == "" {
		fmt.Println("configure profile --name PROFILE_NAME --url API_END_POINT --accesskey ACCESS_KEY")
		fmt.Println("configure profile --name PROFILE_NAME --url API_END_POINT")
		fmt.Println("configure profile --name PROFILE_NAME --access-key ACCESS_KEY")
		return nil
	}

	if FlagUrl != "" && FlagAccessKey == "" {
		urlstring = FlagUrl
	}

	if FlagUrl == "" && FlagAccessKey != "" {
		keystring = FlagAccessKey
	}

	if FlagUrl != "" && FlagAccessKey != "" {
		urlstring = FlagUrl
		keystring = FlagAccessKey
	}

	CreateConfig()
	confs := &ConfigFile{
		PdexUrl: urlstring,
		AccessKey: keystring,
		ConfigFile: configfile,
	}
	WriteConfigs(confs)
	fmt.Println("successfully configured")
	return nil
}

func ListConfigureCommands(context *cli.Context) error {
	SetActingProfile()

	confr, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	key  := confr.AccessKey
	replacement := SubStr(key,0,len(key)-4)
	fmt.Println("End point Url: ",confr.PdexUrl)
	fmt.Println("Access key   : ",strings.Replace(key,replacement , "********", 1))
	fmt.Println("Config File  : ",confr.ConfigFile)
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

func SubStr(s string,pos,length int) string {
    runes:=[]rune(s)
    ln := pos+length
    if ln > len(runes) {
        ln = len(runes)
    }
    return string(runes[pos:ln])
}

func SetConfigFile() string {
	file := "conf.json"
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	if conf.ConfigFile != "" {
		file = conf.ConfigFile
	}
	return file
}
