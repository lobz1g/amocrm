package models

import (
	"encoding/json"
	"strconv"
)

type (
	contact struct {
		generalInterfaces
		getInterface
		request request
	}

	contactData struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	allContacts struct {
		Embedded struct {
			Items []contactData `json:"items"`
		} `json:"_embedded"`
	}
)

func (c contact) GetAll() []byte {
	var contacts allContacts
	resultJson := c.request.Get(contactUrl)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &contacts)
	result, _ := json.Marshal(contacts.Embedded.Items)
	return result
}

func (c contact) GetById(id int) []byte {
	var contacts allContacts
	url := constructUrlWithId(contactUrl, strconv.Itoa(id))
	resultJson := c.request.Get(url)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &contacts)
	result, _ := json.Marshal(contacts.Embedded.Items[0])
	return result
}

func (c contact) GetByResponsibleUser(id int) []byte {
	var contacts allContacts
	url := constructUrlWithResponsible(contactUrl, strconv.Itoa(id))
	resultJson := c.request.Get(url)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &contacts)
	result, _ := json.Marshal(contacts.Embedded.Items)
	return result
}

//todo: add new contact
//todo: add update contact
