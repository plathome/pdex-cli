package configuration

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Pdex struct {
	C *Config
}

type Message struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

var cli = &http.Client{Timeout: time.Duration(5) * time.Second}

func (s *Pdex) Post(msg Message) (err error) {
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(b)

	url := s.C.URL
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}

	res, err := cli.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		return errors.New("post.Failed")
	}

	return
}
