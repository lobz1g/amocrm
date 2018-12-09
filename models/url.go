package models

const accountUrl = "/api/v2/account"
const companyUrl = "/api/v2/companies"
const contactUrl = "/api/v2/contacts"
const leadUrl = "/api/v2/leads"
const taskUrl = "/api/v2/tasks"

func constructUrlWithId(url, id string) string {
	return url + "?id=" + id
}

func constructUrlWithResponsible(url, id string) string {
	return url + "?responsible_user_id=" + id
}

func constructUrlWithElementId(url, id string) string {
	return url + "?element_id=" + id
}

func getUrl(cfg config, url string) string {
	return "https://" + cfg.Domain + ".amocrm.ru" + url
}
