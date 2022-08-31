package controller

import (
	"context"
	"github.com/becosuke/golang-examples/kvstore/internal/adapters/boundary"
	"github.com/becosuke/golang-examples/kvstore/internal/domain/pack"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
	mock "github.com/becosuke/golang-examples/kvstore/mocks/domain/pack"
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
	newMockUsecase := mock.NewMockUsecase(ctrl)
	newKVStoreServiceServer := NewKVStoreServiceServer(config.NewConfig(), newMockUsecase, boundary.NewBoundary())

	assert.Implements(t, (*pb.KVStoreServiceServer)(nil), newKVStoreServiceServer)
}

func TestKVStoreServiceServerImpl_GetPack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	newMockUsecase := mock.NewMockUsecase(ctrl)
	newKVStoreServiceServer := NewKVStoreServiceServer(config.NewConfig(), newMockUsecase, boundary.NewBoundary())

	ctx := context.Background()
	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	newKey := pack.NewKey(keyString)
	newValue := pack.NewValue(valueString)
	newPack := pack.NewPack(newKey, newValue)
	newMockUsecase.EXPECT().Read(ctx, newKey).Return(newPack, nil)
	req := &pb.GetPackRequest{Key: &pb.Key{Body: keyString}}
	res, err := newKVStoreServiceServer.GetPack(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, keyString, res.GetPack().GetKey().GetBody())
	assert.Equal(t, valueString, res.GetPack().GetValue().GetBody())
}

func TestKVStoreServiceServerImpl_CreatePack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	newMockUsecase := mock.NewMockUsecase(ctrl)
	newKVStoreServiceServer := NewKVStoreServiceServer(config.NewConfig(), newMockUsecase, boundary.NewBoundary())

	ctx := context.Background()

	var err error
	_, err = newKVStoreServiceServer.CreatePack(ctx, &pb.CreatePackRequest{})
	assert.Error(t, err)
	_, err = newKVStoreServiceServer.CreatePack(ctx, &pb.CreatePackRequest{Value: &pb.Value{}})
	assert.Error(t, err)

	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	newKey := pack.NewKey(keyString)
	newValue := pack.NewValue(valueString)
	newPack := pack.NewPack(newKey, newValue)
	newMockUsecase.EXPECT().Create(ctx, newValue).Return(newPack, nil)
	res, err := newKVStoreServiceServer.CreatePack(ctx, &pb.CreatePackRequest{Value: &pb.Value{Body: valueString}})
	require.NoError(t, err)
	assert.Equal(t, valueString, res.GetPack().GetValue().GetBody())
}

func TestKvstoreServiceServerImpl_UpdatePack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	newMockUsecase := mock.NewMockUsecase(ctrl)
	newKVStoreServiceServer := NewKVStoreServiceServer(config.NewConfig(), newMockUsecase, boundary.NewBoundary())

	ctx := context.Background()
	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	newKey := pack.NewKey(keyString)
	newValue := pack.NewValue(valueString)
	newPack := pack.NewPack(newKey, newValue)
	newMockUsecase.EXPECT().Update(ctx, newPack).Return(newPack, nil)
	req := &pb.UpdatePackRequest{Key: &pb.Key{Body: keyString}, Value: &pb.Value{Body: valueString}}
	res, err := newKVStoreServiceServer.UpdatePack(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, keyString, res.GetPack().GetKey().GetBody())
	assert.Equal(t, valueString, res.GetPack().GetValue().GetBody())
}

func TestKvstoreServiceServerImpl_DeletePack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	newMockUsecase := mock.NewMockUsecase(ctrl)
	newKVStoreServiceServer := NewKVStoreServiceServer(config.NewConfig(), newMockUsecase, boundary.NewBoundary())

	ctx := context.Background()
	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	newKey := pack.NewKey(keyString)
	newMockUsecase.EXPECT().Delete(ctx, newKey).Return(nil)
	req := &pb.DeletePackRequest{Key: &pb.Key{Body: keyString}}
	res, err := newKVStoreServiceServer.DeletePack(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, &emptypb.Empty{}, res)
}
