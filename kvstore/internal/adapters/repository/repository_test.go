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
	repository := NewRepository(conf, mockSyncmap)

	assert.Implements(t, (*Repository)(nil), repository)
}

func TestRepositoryImpl_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncmap(ctrl)
	repository := NewRepository(conf, mockSyncmap)

	pack := entity.NewPack("kkk", "vvv")
	message := syncmap.NewMessage("kkk", "vvv")
	mockSyncmap.EXPECT().LoadOrStore(gomock.Eq(message)).Return(message, false, nil)

	ctx := context.Background()
	res, err := repository.Create(ctx, pack)

	assert.Equal(t, pack.Key, res.Key)
	assert.Equal(t, pack.Value, res.Value)
	assert.NoError(t, err)
}

func TestRepositoryImpl_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncmap(ctrl)
	repository := NewRepository(conf, mockSyncmap)

	seal := entity.NewSeal("kkk")
	messageKey := syncmap.NewMessageKey("kkk")
	message := syncmap.NewMessage("kkk", "vvv")
	mockSyncmap.EXPECT().Load(gomock.Eq(messageKey)).Return(message, nil)

	ctx := context.Background()
	res, err := repository.Read(ctx, seal)

	assert.Equal(t, seal.Key, res.Key)
	assert.NoError(t, err)
}

func TestRepositoryImpl_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncmap(ctrl)
	repository := NewRepository(conf, mockSyncmap)

	pack := entity.NewPack("kkk", "vvv")
	message := syncmap.NewMessage("kkk", "vvv")
	mockSyncmap.EXPECT().Store(gomock.Eq(message)).Return(message, nil)

	ctx := context.Background()
	res, err := repository.Update(ctx, pack)

	assert.Equal(t, pack.Key, res.Key)
	assert.Equal(t, pack.Value, res.Value)
	assert.NoError(t, err)
}

func TestRepositoryImpl_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncmap(ctrl)
	repository := NewRepository(conf, mockSyncmap)

	seal := entity.NewSeal("kkk")
	messageKey := syncmap.NewMessageKey("kkk")
	mockSyncmap.EXPECT().Delete(gomock.Eq(messageKey)).Return(nil)

	ctx := context.Background()
	err := repository.Delete(ctx, seal)

	assert.NoError(t, err)
}
