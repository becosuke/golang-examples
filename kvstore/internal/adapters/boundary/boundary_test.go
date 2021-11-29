package boundary

import (
	"github.com/becosuke/golang-examples/kvstore/internal/domain/entity"
	"github.com/becosuke/golang-examples/kvstore/internal/pb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBoundary(t *testing.T) {
	SUT := NewBoundary()
	assert.Implements(t, (*Boundary)(nil), SUT)
}

func TestBoundaryImpl_PackDomainToResource(t *testing.T) {
	SUT := NewBoundary()
	pack := entity.NewPack("test-key", "test-value")
	resource := SUT.PackDomainToResource(pack)
	assert.Equal(t, "test-key", resource.Key)
	assert.Equal(t, "test-value", resource.Value)
}

func TestBoundaryImpl_PackResourceToDomain(t *testing.T) {
	SUT := NewBoundary()
	pack := &pb.Pack{Key: "test-key", Value: "test-value"}
	domain := SUT.PackResourceToDomain(pack)
	assert.Equal(t, "test-key", domain.Key)
	assert.Equal(t, "test-value", domain.Value)
}

func TestBoundaryImpl_SealDomainToResource(t *testing.T) {
	SUT := NewBoundary()
	seal := entity.NewSeal("test-key")
	resource := SUT.SealDomainToResource(seal)
	assert.Equal(t, "test-key", resource.Key)
}

func TestBoundaryImpl_SealResourceToDomain(t *testing.T) {
	SUT := NewBoundary()
	seal := &pb.Seal{Key: "test-key"}
	domain := SUT.SealResourceToDomain(seal)
	assert.Equal(t, "test-key", domain.Key)
}
