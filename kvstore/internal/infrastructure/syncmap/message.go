package syncmap

import (
	"github.com/becosuke/golang-examples/kvstore/internal/entities/node"
)

type Message struct {
	Key   string
	Value string
}

func NewMessage(key, value string) *Message {
	return &Message{Key: key, Value: value}
}

func (m *Message) ConvertToEntity() *node.KeyValueNode {
	return node.NewKeyValueNode(m.Key, m.Value)
}
