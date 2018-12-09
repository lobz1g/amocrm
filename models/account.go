package models

type (
	account struct {
		generalInterfaces
		request request
	}
)

func (a account) GetAll() []byte {
	resultJson := a.request.Get(accountUrl)
	return resultJson
}
