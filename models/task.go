package models

import (
	"encoding/json"
	"strconv"
)

type (
	taskInterface interface {
		GetByElementId(id int) []byte
	}

	task struct {
		taskInterface
		generalInterfaces
		getInterface
		updateInterface
		request request
	}

	taskData struct {
		Id                int    `json:"id"`
		ElementId         int    `json:"element_id"`
		ElementType       int    `json:"element_type"`
		CompleteTillAt    int64  `json:"complete_till_at"`
		TaskType          int    `json:"task_type"`
		Text              string `json:"text"`
		CreatedAt         int64  `json:"created_at"`
		UpdatedAt         int64  `json:"updated_at"`
		ResponsibleUserId int    `json:"responsible_user_id"`
		IsCompleted       bool   `json:"is_completed"`
		CreatedBy         int    `json:"created_by"`
		AccountId         int    `json:"account_id"`
		GroupId           int    `json:"group_id"`
		Result            struct {
			Id   int    `json:"id"`
			Text string `json:"text"`
		} `json:"result"`
	}
	allTasks struct {
		Embedded struct {
			Items []taskData `json:"items"`
		} `json:"_embedded"`
	}
)

func (t task) GetAll() []byte {
	var tasks allTasks
	resultJson := t.request.Get(taskUrl)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &tasks)
	result, _ := json.Marshal(tasks.Embedded.Items)
	return result
}

func (t task) GetById(id int) []byte {
	var tasks allTasks
	url := constructUrlWithId(taskUrl, strconv.Itoa(id))
	resultJson := t.request.Get(url)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &tasks)
	result, _ := json.Marshal(tasks.Embedded.Items[0])
	return result
}

func (t task) GetByResponsibleUser(id int) []byte {
	var tasks allTasks
	url := constructUrlWithResponsible(taskUrl, strconv.Itoa(id))
	resultJson := t.request.Get(url)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &tasks)
	result, _ := json.Marshal(tasks.Embedded.Items)
	return result
}

func (t task) GetByElementId(id int) []byte {
	var tasks allTasks
	url := constructUrlWithElementId(taskUrl, strconv.Itoa(id))
	resultJson := t.request.Get(url)
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &tasks)
	result, _ := json.Marshal(tasks.Embedded.Items)
	return result
}

//todo: add new task
func (t task) Update(data []byte) []byte {
	resultJson := t.request.Post(taskUrl, data)
	return resultJson
}
