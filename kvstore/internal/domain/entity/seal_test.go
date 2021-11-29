package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSeal(t *testing.T) {
	seal := NewSeal("test-key")
	assert.Equal(t, "test-key", seal.Key)
}

func TestSeal_GetKey(t *testing.T) {
	seal := NewSeal("test-key")
	assert.Equal(t, "test-key", seal.GetKey())
}
