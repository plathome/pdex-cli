package subcmd

import (
	"bytes"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"
	"net/http"
	"net/url"
)

type DigestData struct {
	Digest string `json:"digest"`
}

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
	FlagUrl 			string
	FlagAccessKey 		string
	FlagProfileName 	string
	FlagAppNameSuffix 	string
	FlagChannelId 		string
	FlagCmdId 			string
	FlagDeviceGroup 	string
	FlagDeviceId 		string
	FlagAppId 			string
	FlagMsgId 			string
)

func GetUtils(urlstr string) (string, error) {
	resp, err := http.Get(urlstr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(htmlData), nil
}

func HmacDigest(urlstr string, secretkey string, deid string) {
	Hmac(urlstr, []string{"key","message"} , []string{secretkey,deid} )
}

func Hmac(link string, parameters []string, values []string) (body string, err error) {
   data := url.Values{}
   for i := range parameters {
      data.Set(parameters[i], values[i])
   }
   resp, err := http.PostForm(link + "/utils/hmac", data)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   htmlData, err := ioutil.ReadAll(resp.Body)
   if err != nil {
		fmt.Println(err)
		os.Exit(1)
    }
    return string(htmlData), nil
}

func ListChannelApi(urlstr string, accesskey string, deid string) (body string, err error) {
   dgparts 		:= strings.Split(deid,".")
   devicegroup  := dgparts[0] + "." + dgparts[1]
   secretkey 	:= GetSecretKey(urlstr + "/devicegroups/" + devicegroup , accesskey)
   parameters 	:= []string{"key","message"}
   values 		:= []string{secretkey,deid}
   data 		:= url.Values{}
   for i := range parameters {
      data.Set(parameters[i], values[i])
   }
   resp, err := http.PostForm(urlstr + "/utils/hmac", data)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   b, _ := ioutil.ReadAll(resp.Body)
   var data2 map[string]interface{}
   json.Unmarshal(b, &data2)
   digest, _ := data2["digest"].(string)
   ListApi(fmt.Sprintf("%s/%s/%s/%s", urlstr, "devices",deid,"channels"), digest)
   return string(digest), nil
}

func ChannelCreateApi(urlstr string, accesskey string, deid string, appid string) (body string, err error) {
   dgparts 		:= strings.Split(deid,".")
   devicegroup  := dgparts[0] + "." + dgparts[1]
   secretkey 	:= GetSecretKey(urlstr + "/devicegroups/" + devicegroup , accesskey)
   parameters 	:= []string{"key","message"}
   values 		:= []string{secretkey,deid}
   data 		:= url.Values{}
   for i := range parameters {
      data.Set(parameters[i], values[i])
   }
   resp, err := http.PostForm(urlstr + "/utils/hmac", data)
   if err != nil {
      return "", err
   }

   defer resp.Body.Close()
   b, _ := ioutil.ReadAll(resp.Body)
   var data2 map[string]interface{}
   json.Unmarshal(b, &data2)
   digest, _ := data2["digest"].(string)
   v := url.Values{}
   v.Add("deid", deid)
   v.Add("app_id", appid)
   s := v.Encode()
   req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", urlstr, "channels"), strings.NewReader(s))
   if err != nil {
   	fmt.Printf("http.NewRequest() error: %v\n", err)
   	return
   }
   req.Header.Add("Authorization", "Bearer " + digest)
   c := &http.Client{}
   ch_resp, err := c.Do(req)
   if err != nil {
   	fmt.Printf("http.Do() error: %v\n", err)
   	return
   }
   defer ch_resp.Body.Close()
   ch_data, err := ioutil.ReadAll(ch_resp.Body)
   if err != nil {
   	fmt.Printf("error: %v\n", err)
   	return
   }
   return string(ch_data), nil
}

func ShowChannelApi(urlstr string, accesskey string, deid string, channelid string) (body string, err error) {
   dgparts 		:= strings.Split(deid,".")
   devicegroup  := dgparts[0] + "." + dgparts[1]
   secretkey 	:= GetSecretKey(urlstr + "/devicegroups/" + devicegroup , accesskey)
   parameters 	:= []string{"key","message"}
   values 		:= []string{secretkey,deid}
   data 		:= url.Values{}
   for i := range parameters {
      data.Set(parameters[i], values[i])
   }
   resp, err := http.PostForm(urlstr + "/utils/hmac", data)
   if err != nil {
      return "", err
   }

   defer resp.Body.Close()
   b, _ := ioutil.ReadAll(resp.Body)
   var data2 map[string]interface{}
   json.Unmarshal(b, &data2)
   digest, _ := data2["digest"].(string)

   initial_str 	:= ListApiReturn(urlstr + "/channels/" + channelid  , digest)
   final_str 	:= ""

   if strings.Contains(initial_str, "error") == true {
   	 final_str = initial_str
   } else {
   		replace_str := fmt.Sprintf(",%s", "\"digest\":\"" + digest +"\"}")
   		final_str   = strings.Replace(initial_str,"}",replace_str,-1)
   	}

   return final_str, nil
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
	req, err := http.NewRequest("POST", urlStr + "/devices/" + deid + "/message", bytes.NewBuffer([]byte(message)))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	req.Header.Set("content_type", "text/plain")
	req.Header.Set("Authorization", "Bearer " + digestkey)

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

func GetAppToken(urlstr string, accesstoken string, appid string) (body string) {
	v := url.Values{}
	s := v.Encode()
	req, err := http.NewRequest("GET", urlstr + "/apps/" + appid , strings.NewReader(s))
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
	key, _ := secretkey["app_token"].(string)
	return key
}

func ReadMessage(sourcetype string, urlstr string, typeid string, apptoken string) {
	v := url.Values{}
	s := v.Encode()
	req, err := http.NewRequest("GET", urlstr + "/" + sourcetype + "/" + typeid + "/messages", strings.NewReader(s))
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

func ReadSingleMessage(sourcetype string, urlstr string, typeid string, apptoken string, msgid string) {
	v := url.Values{}
	s := v.Encode()
	req, err := http.NewRequest("GET", urlstr + "/" + sourcetype + "/" + typeid + "/messages/" + msgid, strings.NewReader(s))
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

func ShowMeApi(urlstr string, accesstoken string) {
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

func ListApi(urlstr string, accesstoken string) {
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
	fmt.Printf("%v\n", string(data))
}

func ListApiReturn(urlstr string, accesstoken string) (body string) {
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
	return string(data)
}

func CreateApi(urlstr string, accesskey string, key string, value string) {
	v := url.Values{}
	s := v.Encode()
	v.Add(key, value)
	req, err := http.NewRequest("POST", urlstr, strings.NewReader(s))
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

func ChannelSendCommand(urlstr string, accesstoken string, channelid string, appid string, command string) {
	apptoken 	:= GetAppToken(urlstr, accesstoken, appid)
	req, err := http.NewRequest("POST", urlstr + "/channels/" + channelid + "/command", bytes.NewBuffer([]byte(command)))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	req.Header.Set("content_type", "text/plain")
	req.Header.Set("Authorization", "Bearer " + apptoken)

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

func SendCommandToAllChannels(urlstr string, accesskey string, deid string) (body string, err error) {
   dgparts 		:= strings.Split(deid,".")
   devicegroup  := dgparts[0] + "." + dgparts[1]
   secretkey 	:= GetSecretKey(urlstr + "/devicegroups/" + devicegroup , accesskey)
   parameters 	:= []string{"key","message"}
   values 		:= []string{secretkey,deid}
   data 		:= url.Values{}
   for i := range parameters {
      data.Set(parameters[i], values[i])
   }
   resp, err := http.PostForm(urlstr + "/utils/hmac", data)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   b, _ := ioutil.ReadAll(resp.Body)
   var data2 map[string]interface{}
   json.Unmarshal(b, &data2)
   digest, _ := data2["digest"].(string)
   ListApi(fmt.Sprintf("%s/%s/%s/%s", urlstr, "devices",deid,"channels"), digest)
   return string(digest), nil
}

func SetActingProfile() {
	conf, err := ReadConfigs()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Failed reading config file. \n")
		os.Exit(1)
	}
	confPath = fmt.Sprintf("%s/%s", confDir, conf.ConfigFile)
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

func FileExists(filename string) bool {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

func AtLeastTwo(a bool, b bool, c bool) bool {
	return a && (b || c) || (b && c)
}
