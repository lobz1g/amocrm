package models

import (
	"../logHandler"
	"net/http"
	"net/url"
	"time"
)

type amo struct {
	Account account
	Company company
	Contact contact
	Lead    lead
	Note    note
	Task    task
	Service service
}

type amoSettings struct {
	Cfg    config
	Client http.Client
}

var client amoSettings

//todo: add error handler
func New() amo {
	client.Cfg = getConfig()
	jar := newJar()
	client.Client = http.Client{Jar: jar, Timeout: 15 * time.Minute, CheckRedirect: nil, Transport: nil}

	values := url.Values{
		"USER_LOGIN": {client.Cfg.Login},
		"USER_HASH":  {client.Cfg.Key},
	}

	resp, err := client.Client.PostForm(getUrl(client.Cfg, "/private/api/auth.php?type=json"), values)
	if err != nil {
		defer logHandler.WriteLogFile(err, "amo", "New()")
	}
	//todo: add checking error
	_ = resp.Body.Close()

	return amo{}
}
