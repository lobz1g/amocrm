package models

import (
	"encoding/json"
)

type (
	service struct {
		serviceInterface
		request request
	}
)

type (
	userData struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Login    string `json:"login"`
		GroupId  int    `json:"group_id"`
		IsActive bool   `json:"is_active"`
	}

	allUsers struct {
		List struct {
			Users map[string]userData `json:"users"`
		} `json:"_embedded"`
	}
)

type (
	taskType struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	allTaskTypes struct {
		List struct {
			TaskTypes map[string]taskType `json:"task_types"`
		} `json:"_embedded"`
	}
)

type (
	leadStatus struct {
		Statuses map[string]struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"statuses"`
	}

	allLeadStatuses struct {
		List struct {
			Pipelines map[string]leadStatus `json:"pipelines"`
		} `json:"_embedded"`
	}
)

type (
	groupData struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	allGroups struct {
		List struct {
			Groups []groupData `json:"groups"`
		} `json:"_embedded"`
	}
)

type (
	customFields struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	allLeadCustomFields struct {
		Embedded struct {
			List struct {
				Items map[string]customFields `json:"leads"`
			} `json:"custom_fields"`
		} `json:"_embedded"`
	}

	allCompanyCustomFields struct {
		Embedded struct {
			List struct {
				Items map[string]customFields `json:"companies"`
			} `json:"custom_fields"`
		} `json:"_embedded"`
	}
)

func (s service) GetUsers() []byte {
	var users allUsers
	resultJson := s.request.Get(accountUrl + "?with=users")
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &users)
	result, _ := json.Marshal(users.List.Users)
	return result
}

func (s service) GetGroups() []byte {
	var groups allGroups
	resultJson := s.request.Get(accountUrl + "?with=groups")
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &groups)
	result, _ := json.Marshal(groups.List.Groups)
	return result
}

func (s service) GetLeadStatuses() []byte {
	var statuses allLeadStatuses
	resultJson := s.request.Get(accountUrl + "?with=pipelines")
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &statuses)
	result, _ := json.Marshal(statuses.List.Pipelines)
	return result
}

func (s service) GetLeadCustomFields() []byte {
	var cf allLeadCustomFields
	resultJson := s.request.Get(accountUrl + "?with=custom_fields")
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &cf)
	result, _ := json.Marshal(cf.Embedded.List.Items)
	return result
}

func (s service) GetCompanyCustomFields() []byte {
	var cf allCompanyCustomFields
	resultJson := s.request.Get(accountUrl + "?with=custom_fields")
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &cf)
	result, _ := json.Marshal(cf.Embedded.List.Items)
	return result
}

func (s service) GetTaskTypes() []byte {
	var taskTypes allTaskTypes
	resultJson := s.request.Get(accountUrl + "?with=task_types")
	//todo: add checking error
	_ = json.Unmarshal(resultJson, &taskTypes)
	result, _ := json.Marshal(taskTypes.List.TaskTypes)
	return result
}
