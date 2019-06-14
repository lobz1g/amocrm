package models

import (
	"strconv"
	"strings"
)

const authUrl = "/private/api/auth.php?type=json"
const accountUrl = "/api/v2/account?with=users,pipelines,groups,note_types,task_types,custom_fields"
const companyUrl = "/api/v2/companies"
const leadUrl = "/api/v2/leads"
const taskUrl = "/api/v2/tasks"
const noteUrl = "/api/v2/notes"

func constructUrlWithId(url string, id int) string {
	return url + "?id=" + strconv.Itoa(id)
}

func constructUrlWithResponsible(url string, id int) string {
	return url + "?responsible_user_id=" + strconv.Itoa(id)
}

func constructUrlWithStatus(url string, id int) string {
	return url + "?responsible_user_id=" + strconv.Itoa(id)
}

func constructUrlWithOffset(url string, rows int) string {
	var newUrl string
	if strings.Contains(url, "?") {
		newUrl = url + "&"
	} else {
		newUrl = url + "?"
	}
	newUrl += "limit_rows=" + strconv.Itoa(limit) + "&limit_offset=" + strconv.Itoa(rows*limit)
	return newUrl
}

func getUrl(url string) string {
	return "https://" + client.Cfg.Domain + ".amocrm.ru" + url
}
