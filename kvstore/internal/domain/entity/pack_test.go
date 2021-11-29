package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPack(t *testing.T) {
	pack := NewPack("test-key", "test-value")
	assert.Equal(t, "test-key", pack.Key)
	assert.Equal(t, "test-value", pack.Value)
}

func TestPack_GetKey(t *testing.T) {
	pack := NewPack("test-key", "test-value")
	assert.Equal(t, "test-key", pack.GetKey())
}

func TestPack_GetValue(t *testing.T) {
	pack := NewPack("test-key", "test-value")
	assert.Equal(t, "test-value", pack.GetValue())
}
