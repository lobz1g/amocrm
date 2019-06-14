package models

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type (
	request struct{}
)

func (r request) Get(address string) ([]byte, error) {
	resp, err := client.Client.Get(getUrl(address))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("status code: " + strconv.Itoa(resp.StatusCode) + "\nerror: " + string(body))
	}

	return body, nil
}

func (r request) Post(address string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", getUrl(address), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("status code: " + strconv.Itoa(resp.StatusCode) + "\nerror: " + string(body))
	}

	return body, nil
}
