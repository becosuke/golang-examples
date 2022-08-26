package pack

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewValue(t *testing.T) {
	newString := uuid.NewString()
	newValue := NewValue(newString)
	assert.Equal(t, Value(newString), *newValue)
}

func TestValue_String(t *testing.T) {
	newValue := NewValue("b3b8500d-3502-4420-a600-49081c68d24b")
	assert.Equal(t, "b3b8500d-3502-4420-a600-49081c68d24b", newValue.String())
}
