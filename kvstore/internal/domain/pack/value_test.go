package pack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewValue(t *testing.T) {
	newString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	newValue := NewValue(newString)
	assert.Equal(t, Value(newString), *newValue)
}

func TestValue_String(t *testing.T) {
	newString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	newValue := NewValue(newString)
	assert.Equal(t, newString, newValue.String())
}
