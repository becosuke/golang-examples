package pack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKey(t *testing.T) {
	newString := "b3b8500d-3502-4420-a600-49081c68d24b"
	newKey := NewKey(newString)
	assert.Equal(t, Key(newString), *newKey)
}

func TestKey_String(t *testing.T) {
	newString := "b3b8500d-3502-4420-a600-49081c68d24b"
	newKey := NewKey(newString)
	assert.Equal(t, newString, newKey.String())
}
