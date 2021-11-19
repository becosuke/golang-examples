package syncmap

import (
	"github.com/becosuke/golang-examples/kvstore/internal/entity/kv"
)

type Message struct {
	Key   string
	Value string
}

func NewMessage(key, value string) *Message {
	return &Message{Key: key, Value: value}
}

func (m *Message) ConvertToEntity() *kv.Pack {
	return kv.NewPack(m.Key, m.Value)
}
