package model

import "encoding/json"

type Model struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewModel(key, value string) *Model {
	return &Model{
		Key:   key,
		Value: value,
	}
}

func (s *Model) GetKey() string {
	return s.Key
}

func (s *Model) GetValue() string {
	return s.Value
}

func (s *Model) ConvertToJson() ([]byte, error) {
	return json.Marshal(s)
}
