package controller

import (
	"context"
	"github.com/becosuke/golang-examples/kvstore/internal/adapters/boundary"
	"github.com/becosuke/golang-examples/kvstore/internal/domain/entity"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
	"github.com/becosuke/golang-examples/kvstore/mocks/usecases/interactor"
	"github.com/becosuke/golang-examples/kvstore/pb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestNewKVStoreServiceServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockInteractor := interactor.NewMockInteractor(ctrl)
	serviceServer := NewKVStoreServiceServer(config.NewConfig(), mockInteractor, boundary.NewBoundary())

	assert.Implements(t, (*pb.KVStoreServiceServer)(nil), serviceServer)
}

func TestKVStoreServiceServerImpl_CreatePack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockInteractor := interactor.NewMockInteractor(ctrl)
	serviceServer := NewKVStoreServiceServer(config.NewConfig(), mockInteractor, boundary.NewBoundary())

	ctx := context.Background()

	var err error
	_, err = serviceServer.CreatePack(ctx, &pb.CreatePackRequest{})
	assert.Error(t, err)
	_, err = serviceServer.CreatePack(ctx, &pb.CreatePackRequest{Pack: &pb.Pack{}})
	assert.Error(t, err)

	mockInteractor.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&entity.Pack{Key: "test-key", Value: "test-value"}, nil)
	res, err := serviceServer.CreatePack(ctx, &pb.CreatePackRequest{Pack: &pb.Pack{Key: "test-key", Value: "test-value"}})
	require.NoError(t, err)
	assert.Equal(t, "test-key", res.Pack.Key)
	assert.Equal(t, "test-value", res.Pack.Value)
}

func TestKVStoreServiceServerImpl_GetPack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockInteractor := interactor.NewMockInteractor(ctrl)
	serviceServer := NewKVStoreServiceServer(config.NewConfig(), mockInteractor, boundary.NewBoundary())

	seal := &pb.Seal{Key: "test-key"}
	mockInteractor.EXPECT().Read(gomock.Any(), gomock.Any()).Return(&entity.Pack{Key: "test-key", Value: "test-value"}, nil)
	ctx := context.Background()
	req := &pb.GetPackRequest{Seal: seal}
	res, err := serviceServer.GetPack(ctx, req)

	assert.Equal(t, "test-key", res.Pack.Key)
	assert.Equal(t, "test-value", res.Pack.Value)
	assert.NoError(t, err)
}

func TestKvstoreServiceServerImpl_UpdatePack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockInteractor := interactor.NewMockInteractor(ctrl)
	serviceServer := NewKVStoreServiceServer(config.NewConfig(), mockInteractor, boundary.NewBoundary())

	pack := &pb.Pack{Key: "test-key", Value: "test-value"}
	mockInteractor.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&entity.Pack{Key: "test-key", Value: "test-value"}, nil)
	ctx := context.Background()
	req := &pb.UpdatePackRequest{Pack: pack}
	res, err := serviceServer.UpdatePack(ctx, req)

	assert.Equal(t, "test-key", res.Pack.Key)
	assert.Equal(t, "test-value", res.Pack.Value)
	assert.NoError(t, err)
}

func TestKvstoreServiceServerImpl_DeletePack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockInteractor := interactor.NewMockInteractor(ctrl)
	serviceServer := NewKVStoreServiceServer(config.NewConfig(), mockInteractor, boundary.NewBoundary())

	seal := &pb.Seal{Key: "test-key"}
	mockInteractor.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
	ctx := context.Background()
	req := &pb.DeletePackRequest{Seal: seal}
	res, err := serviceServer.DeletePack(ctx, req)

	assert.Equal(t, &emptypb.Empty{}, res)
	assert.NoError(t, err)
}
