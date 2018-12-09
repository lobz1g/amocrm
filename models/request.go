package models

import (
	"../logHandler"
	"bytes"
	"io/ioutil"
	"net/http"
)

type (
	requestInterface interface {
		Get(address string) []byte
		Post(address string, data interface{}) []byte
	}

	request struct {
		requestInterface
	}
)

func (request) Get(address string) []byte {
	resp, err := client.Client.Get(getUrl(client.Cfg, address))
	if err != nil {
		defer logHandler.WriteLogFile(err, "request", "Get()")
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		defer logHandler.WriteLogFile(err, "request", "ReadAll()")
	}
	defer resp.Body.Close()

	return b
}

func (request) Post(address string, data []byte) []byte {
	req, err := http.NewRequest("POST", getUrl(client.Cfg, address), bytes.NewBuffer(data))
	if err != nil {
		defer logHandler.WriteLogFile(err, "request", "NewRequest()")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Client.Do(req)
	if err != nil {
		defer logHandler.WriteLogFile(err, "request", "Do()")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		defer logHandler.WriteLogFile(err, "request", "ReadAll()")
	}

	return body
}
