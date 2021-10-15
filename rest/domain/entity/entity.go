package entity

import "encoding/json"

type Entity struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewEntity(key, value string) *Entity {
	return &Entity{
		Key:   key,
		Value: value,
	}
}

func (s *Entity) GetKey() string {
	return s.Key
}

func (s *Entity) GetValue() string {
	return s.Value
}

func (s *Entity) ConvertToJson() ([]byte, error) {
	return json.Marshal(s)
}
