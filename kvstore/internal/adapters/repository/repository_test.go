package repository

import (
	"context"
	"github.com/becosuke/golang-examples/kvstore/internal/domain/entity"
	"github.com/becosuke/golang-examples/kvstore/internal/drivers/syncmap"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
	mock_syncmap "github.com/becosuke/golang-examples/kvstore/mock/drivers/syncmap"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncmap(ctrl)
	SUT := NewRepository(conf, mockSyncmap)

	assert.Implements(t, (*Repository)(nil), SUT)
}

func TestRepositoryImpl_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncmap(ctrl)
	SUT := NewRepository(conf, mockSyncmap)

	pack := entity.NewPack("kkk", "vvv")
	message := &syncmap.Message{Key: "kkk", Value: "vvv"}
	mockSyncmap.EXPECT().LoadOrStore(gomock.Eq(pack.Key), gomock.Eq(pack.Value)).Return(message, false, nil)

	ctx := context.Background()
	res, err := SUT.Create(ctx, pack)

	assert.Equal(t, pack.Key, res.Key)
	assert.Equal(t, pack.Value, res.Value)
	assert.Nil(t, err)
}

func TestRepositoryImpl_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncmap(ctrl)
	SUT := NewRepository(conf, mockSyncmap)

	seal := entity.NewSeal("kkk")
	message := &syncmap.Message{Key: "kkk", Value: "vvv"}
	mockSyncmap.EXPECT().Load(gomock.Eq(seal.Key)).Return(message, nil)

	ctx := context.Background()
	res, err := SUT.Read(ctx, seal)

	assert.Equal(t, seal.Key, res.Key)
	assert.Nil(t, err)
}

func TestRepositoryImpl_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncmap(ctrl)
	SUT := NewRepository(conf, mockSyncmap)

	pack := entity.NewPack("kkk", "vvv")
	message := &syncmap.Message{Key: "kkk", Value: "vvv"}
	mockSyncmap.EXPECT().Store(gomock.Eq(pack.Key), gomock.Eq(pack.Value)).Return(message, nil)

	ctx := context.Background()
	res, err := SUT.Update(ctx, pack)

	assert.Equal(t, pack.Key, res.Key)
	assert.Equal(t, pack.Value, res.Value)
	assert.Nil(t, err)
}

func TestRepositoryImpl_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncmap(ctrl)
	SUT := NewRepository(conf, mockSyncmap)

	packKey := entity.NewSeal("kkk")
	mockSyncmap.EXPECT().Delete(gomock.Eq(packKey.Key)).Return(nil)

	ctx := context.Background()
	err := SUT.Delete(ctx, packKey)

	assert.Nil(t, err)
}
