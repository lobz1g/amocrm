package models

type (
	generalInterfaces interface {
		GetAll() []byte
	}

	getInterface interface {
		GetById(id int) []byte
		GetByResponsibleUser(id int) []byte
	}

	updateInterface interface {
		Update(data []byte) []byte
	}

	addInterface interface {
		Add(data []byte) []byte
	}

	serviceInterface interface {
		GetUsers() []byte
		GetGroups() []byte
		GetLeadStatuses() []byte
		GetLeadCustomFields() []byte
		GetCompanyCustomFields() []byte
		GetTaskTypes() []byte
	}
)
