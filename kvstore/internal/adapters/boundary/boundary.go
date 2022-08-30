package boundary

import (
	"github.com/becosuke/golang-examples/kvstore/internal/domain/pack"
	"github.com/becosuke/golang-examples/kvstore/pb"
	"github.com/google/uuid"
)

type Boundary interface {
	KeyDomainToResource(domainKey *pack.Key) *pb.Key
	KeyResourceToDomain(resourceKey *pb.Key) *pack.Key
	ValueDomainToResource(domainValue *pack.Value) *pb.Value
	ValueResourceToDomain(resourceValue *pb.Value) *pack.Value
	PackDomainToResource(domainPack *pack.Pack) *pb.Pack
	PackResourceToDomain(resourcePack *pb.Pack) *pack.Pack
}

func NewBoundary() Boundary {
	return &boundaryImpl{}
}

type boundaryImpl struct {
}

func (impl *boundaryImpl) KeyDomainToResource(domainKey *pack.Key) *pb.Key {
	return &pb.Key{Body: domainKey.String()}
}

func (impl *boundaryImpl) KeyResourceToDomain(resourceKey *pb.Key) *pack.Key {
	if err := resourceKey.Validate(); err != nil || resourceKey == nil {
		return pack.NewKey(uuid.Nil)
	}
	return pack.NewKey(uuid.MustParse(resourceKey.GetBody()))
}

func (impl *boundaryImpl) ValueDomainToResource(domainValue *pack.Value) *pb.Value {
	return &pb.Value{Body: domainValue.String()}
}

func (impl *boundaryImpl) ValueResourceToDomain(resourceValue *pb.Value) *pack.Value {
	if err := resourceValue.Validate(); err != nil || resourceValue == nil {
		return pack.NewValue("")
	}
	return pack.NewValue(resourceValue.GetBody())
}

func (impl *boundaryImpl) PackDomainToResource(domainPack *pack.Pack) *pb.Pack {
	return &pb.Pack{
		Key:   impl.KeyDomainToResource(domainPack.Key()),
		Value: impl.ValueDomainToResource(domainPack.Value()),
	}
}

func (impl *boundaryImpl) PackResourceToDomain(resourcePack *pb.Pack) *pack.Pack {
	return pack.NewPack(
		impl.KeyResourceToDomain(resourcePack.GetKey()),
		impl.ValueResourceToDomain(resourcePack.GetValue()),
	)
}
