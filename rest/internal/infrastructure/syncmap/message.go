package syncmap

import (
	"github.com/becosuke/golang-examples/rest/internal/entities"
)

type Message struct {
	Key   string
	Value string
}

func NewMessage(key, value string) *Message {
	return &Message{Key: key, Value: value}
}

func (m *Message) ConvertToEntity() *entities.Node {
	return &entities.Node{
		Key:   m.Key,
		Value: m.Value,
	}
}
