package boundary

import (
	"github.com/becosuke/golang-examples/kvstore/internal/domain/entity"
	"github.com/becosuke/golang-examples/kvstore/pb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBoundary(t *testing.T) {
	boundary := NewBoundary()
	assert.Implements(t, (*Boundary)(nil), boundary)
}

func TestBoundaryImpl_PackDomainToResource(t *testing.T) {
	boundary := NewBoundary()
	pack := entity.NewPack("test-key", "test-value")
	resource := boundary.PackDomainToResource(pack)
	assert.Equal(t, "test-key", resource.Key)
	assert.Equal(t, "test-value", resource.Value)
}

func TestBoundaryImpl_PackResourceToDomain(t *testing.T) {
	boundary := NewBoundary()
	pack := &pb.Pack{Key: "test-key", Value: "test-value"}
	domain := boundary.PackResourceToDomain(pack)
	assert.Equal(t, "test-key", domain.Key)
	assert.Equal(t, "test-value", domain.Value)
}

func TestBoundaryImpl_SealDomainToResource(t *testing.T) {
	boundary := NewBoundary()
	seal := entity.NewSeal("test-key")
	resource := boundary.SealDomainToResource(seal)
	assert.Equal(t, "test-key", resource.Key)
}

func TestBoundaryImpl_SealResourceToDomain(t *testing.T) {
	boundary := NewBoundary()
	seal := &pb.Seal{Key: "test-key"}
	domain := boundary.SealResourceToDomain(seal)
	assert.Equal(t, "test-key", domain.Key)
}
