package boundary

import (
	"github.com/becosuke/golang-examples/kvstore/internal/domain/pack"
	"github.com/becosuke/golang-examples/kvstore/pb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBoundary(t *testing.T) {
	newBoundary := NewBoundary()
	assert.Implements(t, (*Boundary)(nil), newBoundary)
}

func TestBoundaryImpl_KeyDomainToResource(t *testing.T) {
	newBoundary := NewBoundary()
	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	domainKey := pack.NewKey(keyString)
	resourceKey := newBoundary.KeyDomainToResource(domainKey)
	assert.Equal(t, &pb.Key{Body: keyString}, resourceKey)
}

func TestBoundaryImpl_KeyResourceToDomain(t *testing.T) {
	newBoundary := NewBoundary()
	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	resourceKey := &pb.Key{Body: keyString}
	domainKey := newBoundary.KeyResourceToDomain(resourceKey)
	assert.Equal(t, pack.NewKey(keyString), domainKey)
}

func TestBoundaryImpl_ValueDomainToResource(t *testing.T) {
	newBoundary := NewBoundary()
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	domainValue := pack.NewValue(valueString)
	resourceValue := newBoundary.ValueDomainToResource(domainValue)
	assert.Equal(t, &pb.Value{Body: valueString}, resourceValue)
}

func TestBoundaryImpl_ValueResourceToDomain(t *testing.T) {
	newBoundary := NewBoundary()
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	resourceValue := &pb.Value{Body: valueString}
	domainValue := newBoundary.ValueResourceToDomain(resourceValue)
	assert.Equal(t, pack.NewValue(valueString), domainValue)
}

func TestBoundaryImpl_PackDomainToResource(t *testing.T) {
	newBoundary := NewBoundary()
	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	domainPack := pack.NewPack(pack.NewKey(keyString), pack.NewValue(valueString))
	resourcePack := newBoundary.PackDomainToResource(domainPack)
	assert.Equal(t, &pb.Pack{Key: &pb.Key{Body: keyString}, Value: &pb.Value{Body: valueString}}, resourcePack)
}

func TestBoundaryImpl_PackResourceToDomain(t *testing.T) {
	newBoundary := NewBoundary()
	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	resourceKey := &pb.Key{Body: keyString}
	resourceValue := &pb.Value{Body: valueString}
	toDomain := newBoundary.PackResourceToDomain(resourceKey, resourceValue)
	assert.Equal(t, pack.NewPack(pack.NewKey(keyString), pack.NewValue(valueString)), toDomain)
}
