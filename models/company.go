package models

import (
	"encoding/json"
	"strconv"
)

type (
	Cmpn struct {
		request request
	}

	company struct {
		Id                int    `json:"id"`
		Name              string `json:"name"`
		ResponsibleUserId int    `json:"responsible_user_id"`
		CreatedBy         int    `json:"created_by"`
		CreatedAt         int64  `json:"created_at"`
		UpdatedAt         int64  `json:"updated_at"`
		AccountId         int    `json:"account_id"`
		UpdatedBy         int    `json:"updated_by"`
		GroupId           int    `json:"group_id"`
		Leads             struct {
			Id []int
		}
		ClosestTaskAt int           `json:"closest_task_at"`
		CustomFields  []customField `json:"custom_fields"`
		Contacts      struct {
			Id []int
		}
	}

	allCompanies struct {
		Embedded struct {
			Items []*company
		} `json:"_embedded"`
	}
)

func (c Cmpn) Create() *company {
	return &company{}
}

func (c Cmpn) All() ([]*company, error) {
	companies := allCompanies{}
	for i := 0; ; i++ {
		var tmpCompanies allCompanies
		resultJson, err := c.request.Get(companyUrl + "?limit_rows=" + strconv.Itoa(limit) + "&limit_offset=" + strconv.Itoa(i*limit))
		if err != nil {
			return nil, err
		}
		json.Unmarshal(resultJson, &tmpCompanies)
		companies.Embedded.Items = append(companies.Embedded.Items, tmpCompanies.Embedded.Items...)
		if len(tmpCompanies.Embedded.Items) < 500 {
			break
		}
	}

	return companies.Embedded.Items, nil
}

func (c Cmpn) Id(id int) (*company, error) {
	var companies allCompanies
	url := constructUrlWithId(companyUrl, id)
	resultJson, err := c.request.Get(url)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resultJson, &companies)
	return companies.Embedded.Items[0], nil
}

func (c Cmpn) Add(cmpn *company) (int, error) {
	data := map[string]interface{}{}
	data["name"] = cmpn.Name
	if cmpn.ResponsibleUserId != 0 {
		data["responsible_user_id"] = cmpn.ResponsibleUserId
	}
	if cmpn.CreatedBy != 0 {
		data["created_by"] = cmpn.CreatedBy
	}
	if len(cmpn.Leads.Id) != 0 {
		data["leads_id"] = getStrFromArr(cmpn.Leads.Id)
	}
	if len(cmpn.Contacts.Id) != 0 {
		data["contacts_id"] = getStrFromArr(cmpn.Contacts.Id)
	}
	if len(cmpn.CustomFields) != 0 {
		data["custom_fields"] = cmpn.CustomFields
	}

	fullData := map[string][]interface{}{"add": {data}}
	jsonData, _ := json.Marshal(fullData)

	resp, err := c.request.Post(companyUrl, jsonData)
	if err != nil {
		return 0, err
	}
	var newCompany allCompanies
	json.Unmarshal(resp, &newCompany)
	return newCompany.Embedded.Items[0].Id, nil
}

func (c Cmpn) Update(cmpn *company) error {
	data := map[string]interface{}{}
	data["id"] = cmpn.Id
	data["name"] = cmpn.Name
	data["updated_at"] = cmpn.UpdatedAt + 1
	data["responsible_user_id"] = cmpn.ResponsibleUserId
	data["custom_fields"] = cmpn.CustomFields
	data["created_by"] = cmpn.CreatedBy
	if len(cmpn.Leads.Id) != 0 {
		data["leads_id"] = getStrFromArr(cmpn.Leads.Id)
	}
	if len(cmpn.Contacts.Id) != 0 {
		data["contacts_id"] = getStrFromArr(cmpn.Contacts.Id)
	}

	fullData := map[string][]interface{}{"update": {data}}
	jsonData, _ := json.Marshal(fullData)

	_, err := c.request.Post(companyUrl, jsonData)
	if err != nil {
		return err
	}
	return nil
}


