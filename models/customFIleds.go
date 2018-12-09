package models

type CustomField struct {
	Id     int                 `json:"id"`
	Name   string              `json:"name"`
	Values []ValuesCustomField `json:"values"`
}

type ValuesCustomField struct {
	Value string `json:"value"`
	Enum  int    `json:"enum"`
}
