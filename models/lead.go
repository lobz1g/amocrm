package models

import (
	"encoding/json"
)

type (
	Ld struct {
		request request
	}

	lead struct {
		Id                int
		Name              string
		ResponsibleUserId int   `json:"responsible_user_id"`
		CreatedBy         int   `json:"created_by"`
		CreatedAt         int64 `json:"created_at"`
		UpdatedAt         int64 `json:"updated_at"`
		AccountId         int   `json:"account_id"`
		PipelineId        int   `json:"pipeline_id"`
		StatusId          int   `json:"status_id"`
		IsDeleted         bool  `json:"is_deleted"`
		MainContact       struct {
			Id int
		}
		GroupId int `json:"group_id"`
		Company struct {
			Id   int
			Name string
		}
		ClosedAt      int           `json:"closed_at"`
		ClosestTaskAt int           `json:"closest_task_at"`
		CustomFields  []customField `json:"custom_fields"`
		Contacts      struct {
			Id []int
		}
		Sale         int
		LossReasonId int `json:"loss_reason_id"`
		Pipeline     struct {
			Id int
		}
	}

	allLeads struct {
		Embedded struct {
			Items []*lead
		} `json:"_embedded"`
	}
)

// Method creates empty struct
func (l Ld) Create() *lead {
	return &lead{}
}

// Method gets all leads from API AmoCRM
//
// Example
//    api := amocrm.NewAmo("login", "key", "domain")
//    allLeads, _ := api.Lead.All()
func (l Ld) All() ([]*lead, error) {
	return l.multiplyRequest(leadUrl)
}

// Method gets all leads by responsible from API AmoCRM
//
// Example
//    api := amocrm.NewAmo("login", "key", "domain")
//    leads, _ := api.Lead.Responsible(12345)
func (l Ld) Responsible(id int) ([]*lead, error) {
	url := constructUrlWithResponsible(leadUrl, id)
	return l.multiplyRequest(url)
}

// Method gets all leads by status from API AmoCRM
//
// Example
//    api := amocrm.NewAmo("login", "key", "domain")
//    leads, _ := api.Lead.Status(12345)
func (l Ld) Status(id int) ([]*lead, error) {
	url := constructUrlWithStatus(leadUrl, id)
	return l.multiplyRequest(url)
}

func (l Ld) multiplyRequest(url string) ([]*lead, error) {
	leads := allLeads{}
	// API returns only 500 rows per request
	// this loop count necessary offset and request data again
	for i := 0; ; i++ {
		var tmpLeads allLeads
		resultJson, err := l.request.Get(constructUrlWithOffset(url, i))
		if err != nil {
			return nil, err
		}
		json.Unmarshal(resultJson, &tmpLeads)
		// sets current data after request to general slice
		leads.Embedded.Items = append(leads.Embedded.Items, tmpLeads.Embedded.Items...)
		if len(tmpLeads.Embedded.Items) < 500 {
			break
		}
	}

	return leads.Embedded.Items, nil
}

// Method gets only one row by ID
func (l Ld) Id(id int) (*lead, error) {
	var leads allLeads
	url := constructUrlWithId(leadUrl, id)
	resultJson, err := l.request.Get(url)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resultJson, &leads)
	return leads.Embedded.Items[0], nil
}

// Note:
//    Field `Name` is required
//    Return id of new lead
//
// Example:
//    api := amocrm.NewAmo("login", "key", "domain")
//    lead := api.Lead.Create()
//    lead.Name = "test"
//    id, err := api.Lead.Add(lead)
func (l Ld) Add(ld *lead) (int, error) {
	data := map[string]interface{}{}
	data["name"] = ld.Name
	if ld.ResponsibleUserId != 0 {
		data["responsible_user_id"] = ld.ResponsibleUserId
	}
	if ld.StatusId != 0 {
		data["status_id"] = ld.StatusId
	}
	if ld.PipelineId != 0 {
		data["pipeline_id"] = ld.PipelineId
	}
	if ld.Sale != 0 {
		data["sale"] = ld.Sale
	}
	if ld.Company.Id != 0 {
		data["company_id"] = ld.Company.Id
	}
	if len(ld.Contacts.Id) != 0 {
		data["contacts_id"] = getStrFromArr(ld.Contacts.Id)
	}
	if len(ld.CustomFields) != 0 {
		data["custom_fields"] = ld.CustomFields
	}

	fullData := map[string][]interface{}{"add": {data}}
	jsonData, _ := json.Marshal(fullData)

	resp, err := l.request.Post(leadUrl, jsonData)
	if err != nil {
		return 0, err
	}
	var newLead allLeads
	json.Unmarshal(resp, &newLead)
	return newLead.Embedded.Items[0].Id, nil
}

// Note:
//    Id is required
//
// Example:
//	   api := amocrm.NewAmo("login", "key", "domain")
//	   lead, _ := api.Lead.Id(123456)
//	   lead.Name = "test"
//	   _ = api.Lead.Update(lead)
func (l Ld) Update(ld *lead) error {
	data := map[string]interface{}{}
	data["id"] = ld.Id
	data["name"] = ld.Name
	data["status_id"] = ld.StatusId
	data["pipeline_id"] = ld.PipelineId
	data["sale"] = ld.Sale
	data["updated_at"] = ld.UpdatedAt + 1
	if ld.Company.Id != 0 {
		data["company_id"] = ld.Company.Id
	}
	data["responsible_user_id"] = ld.ResponsibleUserId
	data["custom_fields"] = ld.CustomFields
	data["created_by"] = ld.CreatedBy
	if len(ld.Contacts.Id) != 0 {
		data["contacts_id"] = getStrFromArr(ld.Contacts.Id)
	}

	fullData := map[string][]interface{}{"update": {data}}
	jsonData, _ := json.Marshal(fullData)

	_, err := l.request.Post(leadUrl, jsonData)
	if err != nil {
		return err
	}
	return nil
}
