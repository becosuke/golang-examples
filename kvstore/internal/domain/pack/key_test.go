package pack

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKey(t *testing.T) {
	newUUID := uuid.MustParse("b3b8500d-3502-4420-a600-49081c68d24b")
	newKey := NewKey(newUUID)
	assert.Equal(t, Key(newUUID), *newKey)
}

func TestKey_UUID(t *testing.T) {
	newUUID := uuid.MustParse("b3b8500d-3502-4420-a600-49081c68d24b")
	newKey := NewKey(newUUID)
	assert.Equal(t, newUUID, newKey.UUID())

	var emptyKey *Key
	assert.Equal(t, uuid.Nil, emptyKey.UUID())

	assert.Equal(t, uuid.Nil, (&Key{}).UUID())
}

func TestKey_String(t *testing.T) {
	newUUID := uuid.MustParse("b3b8500d-3502-4420-a600-49081c68d24b")
	newKey := NewKey(newUUID)
	assert.Equal(t, "b3b8500d-3502-4420-a600-49081c68d24b", newKey.String())
}
