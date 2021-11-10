package repository

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/becosuke/golang-examples/rest/internal/entities"
	"github.com/becosuke/golang-examples/rest/internal/infrastructure/syncmap"
	"github.com/becosuke/golang-examples/rest/internal/registry/config"
	mock_syncmap "github.com/becosuke/golang-examples/rest/tests/mocks/infrastructure/syncmap"
)

func TestNewRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncMap(ctrl)
	SUT := NewRepository(conf, mockSyncmap)

	assert.Implements(t, (*Repository)(nil), SUT)
}

func TestRepositoryImpl_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncMap(ctrl)
	SUT := NewRepository(conf, mockSyncmap)

	node := entities.NewNode("kkk", "vvv")
	message := &syncmap.Message{Key: "kkk", Value: "vvv"}
	mockSyncmap.EXPECT().LoadOrStore(gomock.Eq(node.Key), gomock.Eq(node.Value)).Return(message, false, nil)

	ctx := context.Background()
	res, err := SUT.Create(ctx, node)

	assert.Equal(t, node.Key, res.Key)
	assert.Equal(t, node.Value, res.Value)
	assert.Nil(t, err)
}

func TestRepositoryImpl_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncMap(ctrl)
	SUT := NewRepository(conf, mockSyncmap)

	nodeKey := entities.NewNodeKey("kkk")
	message := &syncmap.Message{Key: "kkk", Value: "vvv"}
	mockSyncmap.EXPECT().Load(gomock.Eq(nodeKey.Key)).Return(message, nil)

	ctx := context.Background()
	res, err := SUT.Read(ctx, nodeKey)

	assert.Equal(t, nodeKey.Key, res.Key)
	assert.Nil(t, err)
}

func TestRepositoryImpl_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncMap(ctrl)
	SUT := NewRepository(conf, mockSyncmap)

	node := entities.NewNode("kkk", "vvv")
	message := &syncmap.Message{Key: "kkk", Value: "vvv"}
	mockSyncmap.EXPECT().Store(gomock.Eq(node.Key), gomock.Eq(node.Value)).Return(message, nil)

	ctx := context.Background()
	res, err := SUT.Update(ctx, node)

	assert.Equal(t, node.Key, res.Key)
	assert.Equal(t, node.Value, res.Value)
	assert.Nil(t, err)
}

func TestRepositoryImpl_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := config.NewConfig()
	mockSyncmap := mock_syncmap.NewMockSyncMap(ctrl)
	SUT := NewRepository(conf, mockSyncmap)

	nodeKey := entities.NewNodeKey("kkk")
	mockSyncmap.EXPECT().Delete(gomock.Eq(nodeKey.Key)).Return(nil)

	ctx := context.Background()
	err := SUT.Delete(ctx, nodeKey)

	assert.Nil(t, err)
}
