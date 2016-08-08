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

	if 	(FlagUrl == "" && FlagAccessKey == "" && FlagUsername == "" && FlagPassword == "") {
		fmt.Println("configure set --url API_END_POINT --access-key ACCESS_KEY")
		fmt.Println("configure set --url API_END_POINT")
		fmt.Println("configure set --access-key ACCESS_KEY")
		fmt.Println("configure set --url API_END_POINT --username USER_NAME --password PASS_WORD")
		return nil
	}

	if FlagUrl != "" && FlagUsername != "" && FlagPassword != "" {
		keystring = UserAccessKeyApi(fmt.Sprintf("%s/%s",FlagUrl,"auth/token"), FlagUsername, FlagPassword)
		urlstring = FlagUrl
	}

	if FlagUrl != "" && FlagAccessKey == "" && FlagUsername == "" && FlagPassword == "" {
		conf, err := ReadConfigs()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
			os.Exit(1)
		}
		urlstring  = conf.PdexUrl
		keystring  = conf.AccessKey

		urlstring = FlagUrl
	}

	if FlagUrl == "" && FlagAccessKey != "" && FlagUsername == "" && FlagPassword == "" {
		conf, err := ReadConfigs()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
			os.Exit(1)
		}
		urlstring  = conf.PdexUrl
		keystring  = conf.AccessKey

		keystring = FlagAccessKey
	}

	if FlagUrl != "" && FlagAccessKey != "" && FlagUsername == "" && FlagPassword == "" {
		urlstring 	= FlagUrl
		keystring 	= FlagAccessKey
	}

	if keystring == "error" {
		fmt.Println("Error in the access-key, try again!")
	} else {
		CreateConfig()
		confs := &ConfigFile{
			PdexUrl: urlstring,
			AccessKey: keystring,
			ConfigFile: "conf.json",
		}
		WriteConfigs(confs)
		fmt.Println("successfully configured")
	}
	return nil
}

func ConfigureCommandsProfile(context *cli.Context) error {
	urlstring  := ""
	keystring  := ""

	if FlagProfileName == "" {
		fmt.Println("configure profile --name PROFILE_NAME")
		return nil
	} else {

		if FileExists(confPath) == false {
			CreateConfig()
			confs := &ConfigFile{
				PdexUrl: "",
				AccessKey: "",
				ConfigFile: FlagProfileName + ".json",
			}
			WriteConfigs(confs)
		} else {
			conf, err := ReadConfigs()
			if err != nil {
				fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
				os.Exit(1)
			}
			urlstring 	:= conf.PdexUrl
			keystring 	:= conf.AccessKey
			configfile 	:= conf.ConfigFile

			if (configfile != FlagProfileName + ".json" ) {
				CreateConfig()
				confs := &ConfigFile{
					PdexUrl: urlstring,
					AccessKey: keystring,
					ConfigFile: FlagProfileName + ".json",
				}
				WriteConfigs(confs)
			}
		}

		confPath = fmt.Sprintf("%s/%s", confDir, FlagProfileName + ".json")

		if FileExists(confPath) == false {
			CreateConfig()
			confs := &ConfigFile{
				PdexUrl: "",
				AccessKey: "",
				ConfigFile: FlagProfileName + ".json",
			}
			WriteConfigs(confs)
		}

		if FlagUrl != "" && FlagAccessKey == "" && FlagUsername == "" && FlagPassword == ""  {
			conf, err := ReadConfigs()
			if err != nil {
				fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
				os.Exit(1)
			}
			urlstring  = FlagUrl
			keystring  = conf.AccessKey
		}

		if FlagUrl == "" && FlagAccessKey != "" && FlagUsername == "" && FlagPassword == "" {
			conf, err := ReadConfigs()
			if err != nil {
				fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
				os.Exit(1)
			}
			keystring = FlagAccessKey
			urlstring = conf.PdexUrl
		}

		if FlagUrl != "" && FlagAccessKey != "" && FlagUsername == "" && FlagPassword == "" {
			urlstring = FlagUrl
			keystring = FlagAccessKey
		}

		if FlagUrl != "" && FlagUsername != "" && FlagPassword != "" {
			keystring = UserAccessKeyApi(fmt.Sprintf("%s/%s",FlagUrl,"auth/token"), FlagUsername, FlagPassword)
			urlstring = FlagUrl
		}

		if 	(FlagAccessKey == "" && FlagUrl == "" && FlagUsername == "" && FlagPassword == "") {
			fmt.Println("configure profile --name PROFILE_NAME --url API_END_POINT --access-key ACCESS_KEY")
			fmt.Println("configure profile --name PROFILE_NAME --url API_END_POINT")
			fmt.Println("configure profile --name PROFILE_NAME --access-key ACCESS_KEY")
			fmt.Println("configure profile --name PROFILE_NAME --url API_END_POINT --username USER_NAME --password PASS_WORD")
			return nil
		}
	}

	if keystring == "error" {
		fmt.Println("Error in the access-key, try again!")
	} else {
		CreateConfig()
		confs := &ConfigFile{
			PdexUrl: urlstring,
			AccessKey: keystring,
			ConfigFile: FlagProfileName + ".json",
		}
		WriteConfigs(confs)
		fmt.Println("successfully configured")
	}
	return nil
}

func ListConfigureCommands(context *cli.Context) error {
	SetActingProfile()
	confr, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	keys := confr.AccessKey
	if len(keys) < 4 {
		keys = "bad-key-"+keys
	}
	replacement := SubStr(keys,0,len(keys)-4)
	fmt.Println("End point Url: ",confr.PdexUrl)
	fmt.Println("Access key   : ",strings.Replace(keys,replacement,"********",1))
	fmt.Println("Config File  : ",confr.ConfigFile)
	return nil
}
