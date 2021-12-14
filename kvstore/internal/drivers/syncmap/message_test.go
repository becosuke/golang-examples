package syncmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMessage(t *testing.T) {
	m := NewMessage("kkk", "vvv")
	assert.Equal(t, "kkk", m.Key.String())
	assert.Equal(t, "vvv", m.Value.String())
}

func TestMessage_GetKey(t *testing.T) {
	m := NewMessage("kkk", "vvv")
	assert.Equal(t, "kkk", m.GetKey().String())
}

func TestMessage_GetValue(t *testing.T) {
	m := NewMessage("kkk", "vvv")
	assert.Equal(t, "vvv", m.GetValue().String())
}
