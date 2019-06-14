package models

import (
	"encoding/json"
	"strconv"
)

type (
	Tsk struct {
		request request
	}

	task struct {
		Id                int
		ElementId         int   `json:"element_id"`
		ElementType       int   `json:"element_type"`
		CompleteTillAt    int64 `json:"complete_till_at"`
		TaskType          int   `json:"task_type"`
		Text              string
		CreatedAt         int64 `json:"created_at"`
		UpdatedAt         int64 `json:"updated_at"`
		ResponsibleUserId int   `json:"responsible_user_id"`
		IsCompleted       bool  `json:"is_completed"`
		CreatedBy         int   `json:"created_by"`
		AccountId         int   `json:"account_id"`
		GroupId           int   `json:"group_id"`
		Result            struct {
			Id   int    `json:"id"`
			Text string `json:"text"`
		} `json:"result"`
	}

	allTasks struct {
		Embedded struct {
			Items []*task
		} `json:"_embedded"`
	}
)

func (t Tsk) Create() *task {
	return &task{}
}

func (t Tsk) All() ([]*task, error) {
	tasks := allTasks{}
	for i := 0; ; i++ {
		var tmpTasks allTasks
		resultJson, err := t.request.get(taskUrl + "?limit_rows=" + strconv.Itoa(limit) + "&limit_offset=" + strconv.Itoa(i*limit))
		if err != nil {
			return nil, err
		}
		json.Unmarshal(resultJson, &tmpTasks)
		tasks.Embedded.Items = append(tasks.Embedded.Items, tmpTasks.Embedded.Items...)
		if len(tmpTasks.Embedded.Items) < 500 {
			break
		}
	}

	return tasks.Embedded.Items, nil
}

func (t Tsk) Id(id int) (*task, error) {
	var tasks allTasks
	url := constructUrlWithId(taskUrl, id)
	resultJson, err := t.request.get(url)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resultJson, &tasks)
	return tasks.Embedded.Items[0], nil
}

func (t Tsk) Add(tsk *task) (int, error) {
	data := map[string]interface{}{}
	data["element_id"] = tsk.ElementId
	data["element_type"] = tsk.ElementType
	data["complete_till_at"] = tsk.CompleteTillAt
	data["task_type"] = tsk.TaskType
	data["text"] = tsk.Text
	data["is_completed"] = tsk.IsCompleted
	if tsk.ResponsibleUserId != 0 {
		data["responsible_user_id"] = tsk.ResponsibleUserId
	}
	if tsk.CreatedBy != 0 {
		data["created_by"] = tsk.CreatedBy
	}

	fullData := map[string][]interface{}{"add": {data}}
	jsonData, _ := json.Marshal(fullData)

	resp, err := t.request.Post(taskUrl, jsonData)
	if err != nil {
		return 0, err
	}
	var newTask allTasks
	json.Unmarshal(resp, &newTask)
	return newTask.Embedded.Items[0].Id, nil
}

func (t Tsk) Update(tsk *task) error {
	data := map[string]interface{}{}
	data["id"] = tsk.Id
	data["updated_at"] = tsk.UpdatedAt + 1
	data["text"] = tsk.Text
	data["element_id"] = tsk.ElementId
	data["element_type"] = tsk.ElementType
	data["complete_till_at"] = tsk.CompleteTillAt
	data["task_type"] = tsk.TaskType
	data["is_completed"] = tsk.IsCompleted
	data["responsible_user_id"] = tsk.ResponsibleUserId
	data["created_by"] = tsk.CreatedBy

	fullData := map[string][]interface{}{"update": {data}}
	jsonData, _ := json.Marshal(fullData)

	_, err := t.request.Post(taskUrl, jsonData)
	if err != nil {
		return err
	}
	return nil
}

func (t Tsk) Close(tsk *task) error {
	n := &note{
		ElementId:         tsk.ElementId,
		ElementType:       taskType,
		NoteType:          taskResultType,
		Text:              tsk.Result.Text,
		ResponsibleUserId: tsk.ResponsibleUserId,
	}
	_, err := Nt{}.Add(n)
	if err != nil {
		return err
	}
	return nil
}
