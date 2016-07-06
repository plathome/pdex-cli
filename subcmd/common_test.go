package subcmd

import (
	"fmt"
	"testing"
	"strings"
)

func TestAtleastTwo(t *testing.T) {
	result := AtLeastTwo(true,true,false)
	if result != true {
		t.Error("The logic is giving error")
	}
}

func TestFileExists(t *testing.T) {
	result := FileExists("../Makefile")
	if result != true {
		t.Error("file is not existing")
	}
}

func TestSubString(t *testing.T) {
	mainstr := "xyzpassword"
	result := SubStr(mainstr,0,len(mainstr)-4)
	if result != "xyzpass" {
		t.Error("incorrect")
	}
}

func TestGetUtils(t *testing.T) {
	urlstr 		:= "http://localhost:9292/api/v1"
	reststr 	:= "/utils/ping"
	result, _ 	:= GetUtils(fmt.Sprintf("%s%s",urlstr,reststr))
	if strings.Contains(result, "no such host") {
		t.Error("the PDExchange is not running")
	} else if strings.Contains(result, "404 Not Found") {
		t.Error("the url is incorrect")
	} else if !strings.Contains(result, "pong") {
		t.Error("the response if incorrect")
	}
}
