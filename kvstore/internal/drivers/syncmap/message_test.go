package syncmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	m := NewMessage("kkk", "vvv")
	assert.Equal(t, "kkk", m.Key)
	assert.Equal(t, "vvv", m.Value)
}

func TestMessage_GetKey(t *testing.T) {
	m := NewMessage("kkk", "vvv")
	assert.Equal(t, "kkk", m.GetKey())
}

func TestMessage_GetValue(t *testing.T) {
	m := NewMessage("kkk", "vvv")
	assert.Equal(t, "vvv", m.GetValue())
}

func TestNewMessageKey(t *testing.T) {
	m := NewMessageKey("kkk")
	assert.Equal(t, "kkk", m.Key)
}

func TestMessageKey_GetKey(t *testing.T) {
	m := NewMessageKey("kkk")
	assert.Equal(t, "kkk", m.GetKey())
}
