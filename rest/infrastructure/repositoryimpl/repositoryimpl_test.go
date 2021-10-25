package repositoryimpl

import (
	"context"
	"github.com/becosuke/golang-examples/rest/infrastructure/memorymap"
	"testing"

	"github.com/becosuke/golang-examples/rest/domain/entity"
	"github.com/becosuke/golang-examples/rest/domain/repository"
	mock_memorymap "github.com/becosuke/golang-examples/rest/mock/memorymap"
	"github.com/becosuke/golang-examples/rest/registry/config"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	key   = "key"
	value = "value"
)

func TestNewRepository(t *testing.T) {
	conf := config.NewConfig()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_memorymap.NewMockMemoryMap(ctrl)
	SUT := NewRepository(conf, mock)

	assert.Implements(t, (*repository.Repository)(nil), SUT)
}

func TestRepositoryImpl_Create(t *testing.T) {
	conf := config.NewConfig()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_memorymap.NewMockMemoryMap(ctrl)
	SUT := NewRepository(conf, mock)

	ctx := context.Background()
	mock.EXPECT().LoadOrStore(gomock.Eq(key), gomock.Eq(value)).Return(&memorymap.Message{Key: key, Value: value}, false, nil)
	result, err := SUT.Create(ctx, key, value)
	assert.Equal(t, &entity.Entity{Key: key, Value: value}, result)
	assert.Nil(t, err)
}

func TestRepositoryImpl_Read(t *testing.T) {
	conf := config.NewConfig()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_memorymap.NewMockMemoryMap(ctrl)
	SUT := NewRepository(conf, mock)

	ctx := context.Background()
	mock.EXPECT().Load(gomock.Eq(key)).Return(&memorymap.Message{Key: key, Value: value}, nil)
	result, err := SUT.Read(ctx, key)
	assert.Equal(t, &entity.Entity{Key: key, Value: value}, result)
	assert.Nil(t, err)
}

func TestRepositoryImpl_Update(t *testing.T) {
	conf := config.NewConfig()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_memorymap.NewMockMemoryMap(ctrl)
	SUT := NewRepository(conf, mock)

	ctx := context.Background()
	mock.EXPECT().Store(gomock.Eq(key), gomock.Eq(value)).Return(&memorymap.Message{Key: key, Value: value}, nil)
	result, err := SUT.Update(ctx, key, value)
	assert.Equal(t, &entity.Entity{Key: key, Value: value}, result)
	assert.Nil(t, err)
}

func TestRepositoryImpl_Delete(t *testing.T) {
	conf := config.NewConfig()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_memorymap.NewMockMemoryMap(ctrl)
	SUT := NewRepository(conf, mock)

	ctx := context.Background()
	mock.EXPECT().Delete(gomock.Eq(key)).Return(nil)
	err := SUT.Delete(ctx, key)
	assert.Nil(t, err)
}
