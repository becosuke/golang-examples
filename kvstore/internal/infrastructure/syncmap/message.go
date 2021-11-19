package syncmap

import "github.com/becosuke/golang-examples/kvstore/internal/entities/pack"

type Message struct {
	Key   string
	Value string
}

func NewMessage(key, value string) *Message {
	return &Message{Key: key, Value: value}
}

func (m *Message) ConvertToEntity() *pack.Pack {
	return pack.NewPack(m.Key, m.Value)
}
