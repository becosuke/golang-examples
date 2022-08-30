package pack

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPack(t *testing.T) {
	newKey := NewKey(uuid.MustParse("b3b8500d-3502-4420-a600-49081c68d24b"))
	newValue := NewValue("25338aef-9462-4c0e-bc8d-e701d3f66cc3")
	pack := NewPack(newKey, newValue)
	assert.Equal(t, Pack{key: newKey, value: newValue}, *pack)
}

func TestPack_Key(t *testing.T) {
	newKey := NewKey(uuid.MustParse("b3b8500d-3502-4420-a600-49081c68d24b"))
	newValue := NewValue("25338aef-9462-4c0e-bc8d-e701d3f66cc3")
	pack := NewPack(newKey, newValue)
	assert.Equal(t, newKey, pack.Key())
}

func TestPack_Value(t *testing.T) {
	newKey := NewKey(uuid.MustParse("b3b8500d-3502-4420-a600-49081c68d24b"))
	newValue := NewValue("25338aef-9462-4c0e-bc8d-e701d3f66cc3")
	pack := NewPack(newKey, newValue)
	assert.Equal(t, newValue, pack.Value())
}
