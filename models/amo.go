package models

import (
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const limit = 500
const delay = 25 * time.Millisecond

type amoSettings struct {
	Cfg    *config
	Client http.Client
}

var client *amoSettings

func OpenConnection() error {
	var err error
	client, err = newAmoSettings()
	if err != nil {
		return err
	}
	err = client.open()
	if err != nil {
		return err
	}
	return nil
}

func newAmoSettings() (*amoSettings, error) {
	cfg, err := getConfig()
	if err != nil {
		return nil, err
	}
	c := &amoSettings{Cfg: cfg, Client: http.Client{}}
	return c, nil
}

func (c *amoSettings) open() error {
	jar := newJar()
	c.Client = http.Client{Jar: jar, Timeout: 15 * time.Minute, CheckRedirect: nil, Transport: nil}

	values := url.Values{
		"USER_LOGIN": {c.Cfg.Login},
		"USER_HASH":  {c.Cfg.Key},
	}

	time.Sleep(delay)
	resp, err := c.Client.PostForm(getUrl(c.Cfg.Domain, "/private/api/auth.php?type=json"), values)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func getStrFromArr(arr []int) string {
	tmp := ""
	for _, value := range arr {
		if tmp != "" {
			tmp += ", "
		}
		tmp += strconv.Itoa(value)
	}

	return tmp
}
