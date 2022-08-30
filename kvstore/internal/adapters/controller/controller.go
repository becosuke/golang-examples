package controller

import (
	"context"
	"github.com/becosuke/golang-examples/kvstore/internal/adapters/boundary"
	"github.com/becosuke/golang-examples/kvstore/internal/domain/pack"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
	"github.com/becosuke/golang-examples/kvstore/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type kvstoreServiceServerImpl struct {
	pb.UnimplementedKVStoreServiceServer
	config   *config.Config
	usecase  pack.Usecase
	boundary boundary.Boundary
}

func NewKVStoreServiceServer(config *config.Config, usecase pack.Usecase, boundary boundary.Boundary) pb.KVStoreServiceServer {
	return &kvstoreServiceServerImpl{
		config:   config,
		usecase:  usecase,
		boundary: boundary,
	}
}

func (impl *kvstoreServiceServerImpl) GetPack(ctx context.Context, req *pb.GetPackRequest) (*pb.GetPackResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	res, err := impl.usecase.Read(ctx, impl.boundary.KeyResourceToDomain(req.GetKey()))
	if err != nil {
		return nil, err
	}
	return &pb.GetPackResponse{Pack: impl.boundary.PackDomainToResource(res)}, nil
}

func (impl *kvstoreServiceServerImpl) CreatePack(ctx context.Context, req *pb.CreatePackRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	err := impl.usecase.Create(ctx, impl.boundary.PackResourceToDomain(nil, req.GetValue()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (impl *kvstoreServiceServerImpl) UpdatePack(ctx context.Context, req *pb.UpdatePackRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	err := impl.usecase.Update(ctx, impl.boundary.PackResourceToDomain(req.GetKey(), req.GetValue()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (impl *kvstoreServiceServerImpl) DeletePack(ctx context.Context, req *pb.DeletePackRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	err := impl.usecase.Delete(ctx, impl.boundary.KeyResourceToDomain(req.GetKey()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
