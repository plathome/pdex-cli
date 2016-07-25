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
	FlagUsername 		string
	FlagPassword 		string
	FlagCurrentPassword	string
	FlagNewPassword		string
	FlagConfirmation 	string
	FlagKey 			string
	FlagValue 			string
	FlagQuery			string
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

func HmacCommandTask(urlstr string, accesskey string, deid string) {
	dgparts 	:= strings.Split(deid,".")
	devicegroup := dgparts[0] + "." + dgparts[1]
	secretkey 	:= GetSecretKey(urlstr + "/devicegroups/" + devicegroup, accesskey)
	result, _ 	:= HmacDigest(urlstr, secretkey, deid)
	fmt.Println(result)
}

func HmacDigest(urlstr string, secretkey string, deid string) (digestkey string, err error) {
	return Hmac(urlstr, []string{"key","message"} , []string{secretkey,deid} )
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

func ReadLatestMessage(urlstr string, apptoken string) {
	v := url.Values{}
	s := v.Encode()
	req, err := http.NewRequest("GET", urlstr, strings.NewReader(s))
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

func ReadFilteredMessage(urlstr string, apptoken string, filterquery string) {
	v := url.Values{}
	queryparts := strings.Split(filterquery, "&")
	for i := 0; i < len(queryparts); i++ {
		subparts := strings.Split(queryparts[i],"=")
		v.Set(subparts[0], subparts[1])
	}
	s := v.Encode()
	req, err := http.NewRequest("GET", urlstr, strings.NewReader(s))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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

func ListAppsApi(baseurl string, urlstr string, accesskey string, appid string) {
	apptoken := GetAppToken(baseurl, accesskey, appid)
	ListApi(urlstr, apptoken)
}

func ListDevicesApi(baseurl string, urlstr string, accesskey string, deid string) {
	dgparts 		:= strings.Split(deid,".")
	devicegroup  	:= dgparts[0] + "." + dgparts[1]
	secretkey 		:= GetSecretKey(baseurl + "/devicegroups/" + devicegroup , accesskey)
	digestdata, err := Hmac(baseurl, []string{"key","message"} , []string{secretkey, deid} )
	jd := new(DigestData)
    err = json.Unmarshal([]byte(digestdata), &jd)
    if err != nil {
        fmt.Println(err)
    }
	ListApi(urlstr, jd.Digest)
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

func CreateUserApi(urlstr string, accesstoken string, username string, password string) {
   parameters 	:=	[]string{"username", "password"}
   values 		:=	[]string{username, password}
   v := url.Values{}
   for i := range parameters {
   	v.Set(parameters[i], values[i])
   }
   s := v.Encode()
   req, err := http.NewRequest("POST", urlstr, strings.NewReader(s))
   if err != nil {
   	fmt.Printf("http.NewRequest() error: %v\n", err)
   	return
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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

func UpdatePasswordApi(urlstr string, accesstoken string, current_password string, new_password string) {
   parameters 	:=	[]string{"current_password", "new_password"}
   values 		:=	[]string{current_password, new_password}
   v := url.Values{}
   for i := range parameters {
   	v.Set(parameters[i], values[i])
   }
   s := v.Encode()
   req, err := http.NewRequest("PUT", urlstr, strings.NewReader(s))
   if err != nil {
   	fmt.Printf("http.NewRequest() error: %v\n", err)
   	return
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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

func UpdateDeviceTapApi(baseurl string, accesskey string, deid string, tagkey string, value string) {
	parameters 		:=	[]string{"value"}
	values 			:=	[]string{value}
	dgparts 		:= 	strings.Split(deid,".")
	devicegroup 	:= 	dgparts[0] + "." + dgparts[1]
	secretkey 		:= 	GetSecretKey(baseurl + "/devicegroups/" + devicegroup , accesskey)
	digestdata, err := Hmac(baseurl, []string{"key","message"} , []string{secretkey, deid} )
	jd 				:= new(DigestData)
    err 			= json.Unmarshal([]byte(digestdata), &jd)
    if err != nil {
        fmt.Println(err)
    }
    updatestr		:= fmt.Sprintf("%s/%s/%s/%s/%s",baseurl,"devices",deid,"tags",tagkey)
	UpdateTagApi(updatestr, jd.Digest, parameters, values)
}

func UpdateAppTagApi(baseurl string, accesskey string, appid string, tagkey string, value string) {
	parameters 	:=	[]string{"value"}
	values 		:=	[]string{value}
	apptoken 	:= 	GetAppToken(baseurl, accesskey, appid)
	updatestr 	:= 	fmt.Sprintf("%s/%s/%s/%s/%s",baseurl,"apps",appid,"tags",tagkey)
	UpdateTagApi(updatestr, apptoken, parameters, values)
}

func UpdateTagApi(urlstr string, accesstoken string, parameters []string, values []string) {
   v := url.Values{}
   for i := range parameters {
   	v.Set(parameters[i], values[i])
   }
   s := v.Encode()
   req, err := http.NewRequest("PUT", urlstr, strings.NewReader(s))
   if err != nil {
   	fmt.Printf("http.NewRequest() error: %v\n", err)
   	return
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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

func UpdateAppApi(urlstr string, accesstoken string, appnamesuffix string) {
   parameters 	:=	[]string{"app_name_suffix"}
   values 		:=	[]string{appnamesuffix}
   v := url.Values{}
   for i := range parameters {
   	v.Set(parameters[i], values[i])
   }
   s := v.Encode()
   req, err := http.NewRequest("PUT", urlstr, strings.NewReader(s))
   if err != nil {
   	fmt.Printf("http.NewRequest() error: %v\n", err)
   	return
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=app_name_suffix")
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

func UpdateSessionApi(urlstr string, accesstoken string) {
	v := url.Values{}
	s := v.Encode()
	req, err := http.NewRequest("PUT", urlstr, strings.NewReader(s))
	if err != nil {
		fmt.Printf("http put request erros: %v\n", err)
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
		fmt.Printf("error: %v\n",  err)
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

func CreateDeviceTagsApi(urlstr string, deid string, accesskey string, key string, value string) {
	dgparts := strings.Split(deid,".")
	devicegroup  := dgparts[0] + "." + dgparts[1]
	secretkey := GetSecretKey(urlstr + "/devicegroups/" + devicegroup , accesskey)
	digestdata, err := Hmac(urlstr, []string{"key","message"} , []string{secretkey, deid} )
	jd := new(DigestData)
    err = json.Unmarshal([]byte(digestdata), &jd)
    if err != nil {
        fmt.Println(err)
    }
	CreateApi(fmt.Sprintf("%s/%s/%s/%s/%s", urlstr, "devices", deid, "tags", key) , jd.Digest,  "value", value)
}

func CreateAppTagsApi(urlstr string, appid string, accesskey string, key string, value string) {
	apptoken := GetAppToken(urlstr, accesskey, appid)
	CreateApi(fmt.Sprintf("%s/%s/%s/%s/%s", urlstr, "apps", appid, "tags", key) , apptoken,  "value", value)
}

func CreateApi(urlstr string, accesskey string, key string, value string) {
	parameters 	:= []string{key}
	values 		:= []string{value}
	v 			:= url.Values{}
	for i := range parameters {
		v.Set(parameters[i], values[i])
	}
	s := v.Encode()
	req, err := http.NewRequest("POST", urlstr, strings.NewReader(s))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=" + key )
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

func ReadCommandsApiFilter(urlstr string, deid string, accesstoken string, commandstr string, filterquery string) {
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
	queryparts := strings.Split(filterquery, "&")
	for i := 0; i < len(queryparts); i++ {
		subparts := strings.Split(queryparts[i],"=")
		v.Set(subparts[0], subparts[1])
	}
	s := v.Encode()
	req, err := http.NewRequest("GET", urlstr + "/devices/" + deid + "/" + commandstr, strings.NewReader(s))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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

func DeleteAccountTask(urlstr string, accesskey string, confirm string) {
	parameters 		:=	[]string{"confirm"}
	values 			:=	[]string{confirm}
	if confirm == "true" {
		DeleteApi(fmt.Sprintf("%s/%s", urlstr, "me"), accesskey, parameters, values, "DELETE")
	} else {
		fmt.Println("Please provide true for delete confirmation")
	}
}

func DeleteChannelTask(urlstr string, accesskey string, deid string, channelid string, confirm string) {
	dgparts 		:= strings.Split(deid,".")
	devicegroup 	:= dgparts[0] + "." + dgparts[1]
	secretkey 		:= GetSecretKey(urlstr + "/devicegroups/" + devicegroup , accesskey)
	digestkey, err 	:= HmacDigest(urlstr, secretkey, deid)
	jd 				:= new(DigestData)
	err 			= json.Unmarshal([]byte(digestkey), &jd)
	if err != nil {
		fmt.Println(err)
	}
	parameters 		:=	[]string{"confirm"}
	values 			:=	[]string{confirm}
	if confirm == "true" {
		DeleteApi(fmt.Sprintf("%s/%s/%s", urlstr, "channels", channelid), jd.Digest, parameters, values, "DELETE")
	} else {
		fmt.Println("Please provide true for delete confirmation")
	}
}

func DeleteDeviceTagApi(urlstr string, deid string, accesskey string, key string) {
	dgparts := strings.Split(deid,".")
	devicegroup  := dgparts[0] + "." + dgparts[1]
	secretkey := GetSecretKey(urlstr + "/devicegroups/" + devicegroup , accesskey)
	digestdata, err := Hmac(urlstr, []string{"key","message"} , []string{secretkey, deid} )
	jd := new(DigestData)
    err = json.Unmarshal([]byte(digestdata), &jd)
    if err != nil {
        fmt.Println(err)
    }
    parameters 	:=	[]string{""}
    values 		:=	[]string{""}
    DeleteApi(fmt.Sprintf("%s/%s/%s/%s/%s", urlstr, "devices", deid, "tags", key), jd.Digest, parameters, values, "DELETE")
}

func DeleteAppTagApi(urlstr string, accesskey string, appid string, key string) {
	apptoken := GetAppToken(urlstr, accesskey, appid)
	parameters 		:=	[]string{""}
	values 			:=	[]string{""}
	DeleteApi(fmt.Sprintf("%s/%s/%s/%s/%s", urlstr, "apps", appid,"tags",key), apptoken, parameters, values, "DELETE")
}

func DeleteApi(urlstr string, key string, parameters []string, values []string, method string) {
	v := url.Values{}
	for i := range parameters {
		v.Set(parameters[i], values[i])
	}
	s := v.Encode()
	req, err := http.NewRequest(method, urlstr, strings.NewReader(s))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Add("Authorization", "Bearer " + key)
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
