package models

type (
	note struct {
		addInterface
		request request
	}
)

//todo: add notes for other models (e.g. company)
func (n note) Add(data []byte) []byte {
	resultJson := n.request.Post(leadUrl, data)
	return resultJson
}
