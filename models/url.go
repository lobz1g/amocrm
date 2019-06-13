package models

import (
	"strconv"
)

const accountUrl = "/api/v2/account?with=users,pipelines,groups,note_types,task_types,custom_fields"
const companyUrl = "/api/v2/companies"
const leadUrl = "/api/v2/leads"
const taskUrl = "/api/v2/tasks"
const noteUrl = "/api/v2/notes"

func constructUrlWithId(url string, id int) string {
	return url + "?id=" + strconv.Itoa(id)
}

func getUrl(domain string, url string) string {
	return "https://" + domain + ".amocrm.ru" + url
}
