package boundary

import (
	"github.com/becosuke/golang-examples/kvstore/internal/domain/entity"
	"github.com/becosuke/golang-examples/kvstore/pb"
)

type Boundary interface {
	PackDomainToResource(pack *entity.Pack) *pb.Pack
	PackResourceToDomain(pack *pb.Pack) *entity.Pack
	SealDomainToResource(seal *entity.Seal) *pb.Seal
	SealResourceToDomain(seal *pb.Seal) *entity.Seal
}

func NewBoundary() Boundary {
	return &boundaryImpl{}
}

type boundaryImpl struct {
}

func (*boundaryImpl) PackDomainToResource(pack *entity.Pack) *pb.Pack {
	return &pb.Pack{
		Key:   pack.GetKey(),
		Value: pack.GetValue(),
	}
}

func (*boundaryImpl) PackResourceToDomain(pack *pb.Pack) *entity.Pack {
	return &entity.Pack{
		Key:   pack.GetKey(),
		Value: pack.GetValue(),
	}
}

func (*boundaryImpl) SealDomainToResource(seal *entity.Seal) *pb.Seal {
	return &pb.Seal{Key: seal.GetKey()}
}

func (*boundaryImpl) SealResourceToDomain(seal *pb.Seal) *entity.Seal {
	return &entity.Seal{Key: seal.GetKey()}
}
