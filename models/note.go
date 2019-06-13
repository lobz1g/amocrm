package models

import (
	"encoding/json"
)

const taskType = 4
const taskResultType = 13

type (
	Nt struct {
		request request
	}

	note struct {
		Id                int
		ElementId         int `json:"element_id"`
		ElementType       int `json:"element_type"`
		Text              string
		NoteType          int `json:"note_type"`
		ResponsibleUserId int `json:"responsible_user_id"`
	}

	allNotes struct {
		Embedded struct {
			Items []*note
		} `json:"_embedded"`
	}
)

func (n Nt) Create() *note {
	return &note{}
}

func (n Nt) Add(nt *note) (int, error) {
	data := map[string]interface{}{}
	data["element_id"] = nt.ElementId
	data["element_type"] = nt.ElementType
	data["text"] = nt.Text
	data["note_type"] = nt.NoteType
	if nt.ResponsibleUserId != 0 {
		data["responsible_user_id"] = nt.ResponsibleUserId
	}

	fullData := map[string][]interface{}{"add": {data}}
	jsonData, _ := json.Marshal(fullData)

	resp, err := n.request.Post(noteUrl, jsonData)
	if err != nil {
		return 0, err
	}
	var newNote allNotes
	json.Unmarshal(resp, &newNote)
	return newNote.Embedded.Items[0].Id, nil
}
