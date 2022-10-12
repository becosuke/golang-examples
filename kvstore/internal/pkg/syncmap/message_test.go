package syncmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMessage(t *testing.T) {
	m := NewMessage("kkk", "vvv")
	assert.Equal(t, Message{key: "kkk", value: "vvv"}, *m)
}

func TestMessage_GetKey(t *testing.T) {
	m := NewMessage("kkk", "vvv")
	assert.Equal(t, "kkk", m.Key())
}

func TestMessage_GetValue(t *testing.T) {
	m := NewMessage("kkk", "vvv")
	assert.Equal(t, "vvv", m.Value())
}
