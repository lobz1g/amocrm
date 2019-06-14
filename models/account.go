package models

import (
	"encoding/json"
)

type (
	Acc struct {
		request request
	}

	customFieldsUser struct {
		Id          int
		Name        string
		FieldType   int `json:"field_type"`
		Sort        int
		Code        string
		IsMultiply  bool `json:"is_multiply"`
		IsSystem    bool `json:"is_system"  `
		IsEditable  bool `json:"is_editable"`
		IsRequired  bool `json:"is_required"`
		IsDeletable bool `json:"is_deletable"`
		IsVisible   bool `json:"is_visible"`
		Enums       map[string]string
	}

	account struct {
		Id             int
		Name           string
		Subdomain      string
		Currency       string
		Timezone       string
		TimezoneOffset string `json:"timezone_offset"`
		Language       string
		DatePattern    struct {
			Date     string
			Time     string
			DateTime string `json:"date_time"`
			TimeFull string `json:"time_full"`
		} `json:"date_pattern"`
		CurrentUser int `json:"current_user"`
		Embedded    struct {
			Users map[string]struct {
				Id       int
				Name     string
				Login    string
				Language string
				GroupId  int  `json:"group_id"`
				IsActive bool `json:"is_active"`
				IsFree   bool `json:"is_free"`
				IsAdmin  bool `json:"is_admin"`
				Rights   struct {
					Mail          string
					IncomingLeads string `json:"incoming_leads"`
					Catalogs      string
					LeadAdd       string `json:"lead_add"`
					LeadView      string `json:"lead_view"`
					LeadEdit      string `json:"lead_edit"`
					LeadDelete    string `json:"lead_delete"`
					LeadExport    string `json:"lead_export"`
					ContactView   string `json:"contact_view"`
					ContactEdit   string `json:"contact_edit"`
					ContactDelete string `json:"contact_delete"`
					ContactExport string `json:"contact_export"`
					CompanyAdd    string `json:"company_add"`
					CompanyView   string `json:"company_view"`
					CompanyEdit   string `json:"company_edit"`
					CompanyDelete string `json:"company_delete"`
					CompanyExport string `json:"company_export"`
				}
			}
			CustomFields struct {
				Contacts  map[string]customFieldsUser
				Leads     map[string]customFieldsUser
				Companies map[string]customFieldsUser
				Catalogs  map[string]map[string]customFieldsUser
			} `json:"custom_fields"`
			NoteTypes map[string]struct {
				Id       int
				Name     string
				Code     string
				Editable bool
			} `json:"note_types"`
			Groups []struct {
				Id   int
				Name string
			}
			TaskTypes map[string]struct {
				Id   int
				Name string
			} `json:"task_types"`
			Pipelines map[string]struct {
				Id       int
				Name     string
				Sort     int
				IsMain   bool `json:"is_main"`
				Statuses map[string]struct {
					Id       int
					Name     string
					Sort     int
					Color    string
					Editable bool
				}
			}
		} `json:"_embedded"`
	}
)

func (a Acc) Get() (*account, error) {
	resultJson, err := a.request.get(accountUrl)
	if err != nil {
		return nil, err
	}

	var account *account
	err = json.Unmarshal(resultJson, &account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
