package pack

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPack(t *testing.T) {
	newKey := NewKey(uuid.New())
	newValue := NewValue(uuid.NewString())
	pack := NewPack(newKey, newValue)
	assert.Equal(t, Pack{key: newKey, value: newValue}, *pack)
}

func TestPack_Key(t *testing.T) {
	newKey := NewKey(uuid.New())
	newValue := NewValue(uuid.NewString())
	pack := NewPack(newKey, newValue)
	assert.Equal(t, newKey, pack.Key())
}

func TestPack_Value(t *testing.T) {
	newKey := NewKey(uuid.New())
	newValue := NewValue(uuid.NewString())
	pack := NewPack(newKey, newValue)
	assert.Equal(t, newValue, pack.Value())
}
