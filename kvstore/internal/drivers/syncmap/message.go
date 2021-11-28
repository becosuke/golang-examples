package syncmap

import (
	"github.com/becosuke/golang-examples/kvstore/internal/domain/entity"
)

type Message struct {
	Key   string
	Value string
}

func NewMessage(key, value string) *Message {
	return &Message{Key: key, Value: value}
}

func (m *Message) ConvertToEntity() *entity.Pack {
	return entity.NewPack(m.Key, m.Value)
}
