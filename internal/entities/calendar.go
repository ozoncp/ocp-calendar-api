package entities

import (
	"encoding/json"
)

type Calendar struct {
	Id     uint64 `json:"id"`
	UserId uint64 `json:"user_id"`
	Type   uint64 `json:"name"`
	Link   string `json:"link"`
}

func (calendar Calendar) ToJson() (string, error) {
	result, err := json.MarshalIndent(calendar, "", "    ")
	return string(result), err
}
