package amocrm

import (
	"github.com/lobz1g/amocrm/models"
)

type amo struct {
	Account models.Acc
	Company models.Cmpn
	Lead    models.Ld
	Task    models.Tsk
	Note    models.Nt
}

func NewAmo(login, key, domain string) *amo {
	models.OpenConnection(login, key, domain)
	return &amo{}
}
