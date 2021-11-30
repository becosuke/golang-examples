package controller

import (
	"context"
	"github.com/becosuke/golang-examples/kvstore/internal/adapters/boundary"
	"github.com/becosuke/golang-examples/kvstore/internal/pb"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
	"github.com/becosuke/golang-examples/kvstore/internal/usecases/interactor"
	"google.golang.org/protobuf/types/known/emptypb"
)

type kvstoreServiceServerImpl struct {
	pb.UnimplementedKVStoreServiceServer
	config     *config.Config
	interactor interactor.Interactor
	boundary   boundary.Boundary
}

func NewKVStoreServiceServer(config *config.Config, interactor interactor.Interactor, boundary boundary.Boundary) pb.KVStoreServiceServer {
	return &kvstoreServiceServerImpl{
		config:     config,
		interactor: interactor,
		boundary:   boundary,
	}
}

func (i *kvstoreServiceServerImpl) CreatePack(ctx context.Context, req *pb.CreatePackRequest) (*pb.CreatePackResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	err = req.GetPack().Validate()
	if err != nil {
		return nil, err
	}
	pack, err := i.interactor.Create(ctx, i.boundary.PackResourceToDomain(req.GetPack()))
	if err != nil {
		return nil, err
	}
	return &pb.CreatePackResponse{Pack: i.boundary.PackDomainToResource(pack)}, nil
}

func (i *kvstoreServiceServerImpl) GetPack(ctx context.Context, req *pb.GetPackRequest) (*pb.GetPackResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	err = req.GetSeal().Validate()
	if err != nil {
		return nil, err
	}
	pack, err := i.interactor.Read(ctx, i.boundary.SealResourceToDomain(req.GetSeal()))
	if err != nil {
		return nil, err
	}
	return &pb.GetPackResponse{Pack: i.boundary.PackDomainToResource(pack)}, nil
}

func (i *kvstoreServiceServerImpl) UpdatePack(ctx context.Context, req *pb.UpdatePackRequest) (*pb.UpdatePackResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	err = req.GetPack().Validate()
	if err != nil {
		return nil, err
	}
	pack, err := i.interactor.Update(ctx, i.boundary.PackResourceToDomain(req.GetPack()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdatePackResponse{Pack: i.boundary.PackDomainToResource(pack)}, nil
}

func (i *kvstoreServiceServerImpl) DeletePack(ctx context.Context, req *pb.DeletePackRequest) (*emptypb.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	err = req.GetSeal().Validate()
	if err != nil {
		return nil, err
	}
	err = i.interactor.Delete(ctx, i.boundary.SealResourceToDomain(req.GetSeal()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
