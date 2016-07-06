package subcmd

import (
	"os"
	"fmt"
	"strings"
	"github.com/urfave/cli"
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
		if FileExists(confPath) == false {
			CreateConfig()
			confs := &ConfigFile{
				PdexUrl: "",
				AccessKey: "",
				ConfigFile: FlagProfileName + ".json",
			}
			WriteConfigs(confs)
		}

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
