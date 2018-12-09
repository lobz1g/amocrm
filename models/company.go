package models

import (
	"encoding/json"
	"strconv"
)

type (
	company struct {
		generalInterfaces
		getInterface
		updateInterface
		request request
	}

	companyData struct {
		Id                int    `json:"id"`
		Name              string `json:"name"`
		ResponsibleUserId int    `json:"responsible_user_id"`
		CreatedBy         int    `json:"created_by"`
		CreatedAt         int64  `json:"created_at"`
		UpdatedAt         int64  `json:"updated_at"`
		AccountId         int    `json:"account_id"`
		UpdatedBy         int    `json:"updated_by"`
		GroupId           int    `json:"group_id"`
		Contacts          struct {
			Id []int `json:"id"`
		} `json:"contacts"`
		CustomFields []CustomField `json:"custom_fields"`
	}

	allCompanies struct {
		Embedded struct {
			Items []companyData `json:"items"`
		} `json:"_embedded"`
	}
)

func (c company) GetAll() []byte {
	var companies allCompanies
	resultJson := c.request.Get(companyUrl)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &companies)
	result, _ := json.Marshal(companies.Embedded.Items)
	return result
}

func (c company) GetById(id int) []byte {
	var companies allCompanies
	url := constructUrlWithId(companyUrl, strconv.Itoa(id))
	resultJson := c.request.Get(url)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &companies)
	result, _ := json.Marshal(companies.Embedded.Items[0])
	return result
}

func (c company) GetByResponsibleUser(id int) []byte {
	var companies allCompanies
	url := constructUrlWithResponsible(companyUrl, strconv.Itoa(id))
	resultJson := c.request.Get(url)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &companies)
	result, _ := json.Marshal(companies.Embedded.Items)
	return result
}

//todo: add new company

func (c company) Update(data []byte) []byte {
	resultJson := c.request.Post(companyUrl, data)
	return resultJson
}
