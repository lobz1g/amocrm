package models

import (
	"encoding/json"
	"strconv"
)

type (
	lead struct {
		generalInterfaces
		getInterface
		updateInterface
		request request
	}

	leadData struct {
		Id                int    `json:"id"`
		Name              string `json:"name"`
		ResponsibleUserId int    `json:"responsible_user_id"`
		CreatedBy         int    `json:"created_by"`
		CreatedAt         int64  `json:"created_at"`
		UpdatedAt         int64  `json:"updated_at"`
		ClosedAt          int64  `json:"closed_at"`
		ClosestTaskAt     int64  `json:"closest_task_at"`
		AccountId         int    `json:"account_id"`
		UpdatedBy         int    `json:"updated_by"`
		GroupId           int    `json:"group_id"`
		IsDeleted         bool   `json:"is_deleted"`
		StatusId          int    `json:"status_id"`
		Sale              int    `json:"sale"`
		MainContact       struct {
			Id int `json:"id"`
		} `json:"main_contact"`
		Contacts struct {
			Id []int `json:"id"`
		} `json:"contacts"`
		CustomFields []CustomField `json:"custom_fields"`
		Company      struct {
			Id int `json:"id"`
		} `json:"company"`
		Tags struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"tags"`
	}

	allLeads struct {
		Embedded struct {
			Items []leadData `json:"items"`
		} `json:"_embedded"`
	}
)

func (l lead) GetAll() []byte {
	var leads allLeads
	resultJson := l.request.Get(leadUrl)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &leads)
	result, _ := json.Marshal(leads.Embedded.Items)
	return result
}

func (l lead) GetById(id int) []byte {
	var leads allLeads
	url := constructUrlWithId(leadUrl, strconv.Itoa(id))
	resultJson := l.request.Get(url)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &leads)
	result, _ := json.Marshal(leads.Embedded.Items[0])
	return result
}

func (l lead) GetByResponsibleUser(id int) []byte {
	var leads allLeads
	url := constructUrlWithResponsible(leadUrl, strconv.Itoa(id))
	resultJson := l.request.Get(url)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &leads)
	result, _ := json.Marshal(leads.Embedded.Items)
	return result
}

//todo: add new lead

func (l lead) Update(data []byte) []byte {
	resultJson := l.request.Post(leadUrl, data)
	return resultJson
}
