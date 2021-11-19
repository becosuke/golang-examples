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

func TestMessage_ConvertToEntity(t *testing.T) {
	m := NewMessage("kkk", "vvv")
	e := m.ConvertToEntity()
	assert.Equal(t, "kkk", e.Key)
	assert.Equal(t, "vvv", e.Value)
}
