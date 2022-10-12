package gateway

import (
	"context"
	"github.com/becosuke/golang-examples/kvstore/internal/domain/pack"
	"github.com/becosuke/golang-examples/kvstore/internal/pkg/syncmap"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
	mockDomain "github.com/becosuke/golang-examples/kvstore/mocks/domain/pack"
	mockSyncmap "github.com/becosuke/golang-examples/kvstore/mocks/drivers/syncmap"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newConfig := config.NewConfig()
	newMockSyncmap := mockSyncmap.NewMockSyncmap(ctrl)
	newMockGenerator := mockDomain.NewMockGenerator(ctrl)
	newRepository := NewRepository(newConfig, newMockSyncmap, newMockGenerator)

	assert.Implements(t, (*pack.Repository)(nil), newRepository)
}

func TestRepositoryImpl_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newConfig := config.NewConfig()
	newMockSyncmap := mockSyncmap.NewMockSyncmap(ctrl)
	newMockGenerator := mockDomain.NewMockGenerator(ctrl)
	newRepository := NewRepository(newConfig, newMockSyncmap, newMockGenerator)

	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	newKey := pack.NewKey(keyString)
	newMessage := syncmap.NewMessage(keyString, valueString)
	newMockSyncmap.EXPECT().Load(gomock.Eq(keyString)).Return(newMessage, nil)

	ctx := context.Background()
	res, err := newRepository.Read(ctx, newKey)
	assert.NoError(t, err)
	assert.Equal(t, keyString, res.Key().String())
	assert.Equal(t, valueString, res.Value().String())
}

func TestRepositoryImpl_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newConfig := config.NewConfig()
	newMockSyncmap := mockSyncmap.NewMockSyncmap(ctrl)
	newMockGenerator := mockDomain.NewMockGenerator(ctrl)
	newRepository := NewRepository(newConfig, newMockSyncmap, newMockGenerator)

	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	newKey := pack.NewKey(keyString)
	newValue := pack.NewValue(valueString)
	newMockGenerator.EXPECT().GenerateKey().Return(newKey)
	newMessage := syncmap.NewMessage(keyString, valueString)
	newMockSyncmap.EXPECT().LoadOrStore(gomock.Eq(newMessage)).Return(newMessage, false, nil)

	ctx := context.Background()
	res, err := newRepository.Create(ctx, newValue)
	assert.NoError(t, err)
	assert.Equal(t, newKey, res)
}

func TestRepositoryImpl_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newConfig := config.NewConfig()
	newMockSyncmap := mockSyncmap.NewMockSyncmap(ctrl)
	newMockGenerator := mockDomain.NewMockGenerator(ctrl)
	newRepository := NewRepository(newConfig, newMockSyncmap, newMockGenerator)

	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	newPack := pack.NewPack(pack.NewKey(keyString), pack.NewValue(valueString))
	newMessage := syncmap.NewMessage(keyString, valueString)
	newMockSyncmap.EXPECT().Store(gomock.Eq(newMessage)).Return(newMessage, nil)

	ctx := context.Background()
	err := newRepository.Update(ctx, newPack)
	assert.NoError(t, err)
}

func TestRepositoryImpl_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newConfig := config.NewConfig()
	newMockSyncmap := mockSyncmap.NewMockSyncmap(ctrl)
	newMockGenerator := mockDomain.NewMockGenerator(ctrl)
	newRepository := NewRepository(newConfig, newMockSyncmap, newMockGenerator)

	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	newKey := pack.NewKey(keyString)
	newMockSyncmap.EXPECT().Delete(gomock.Eq(keyString)).Return(nil)

	ctx := context.Background()
	err := newRepository.Delete(ctx, newKey)
	assert.NoError(t, err)
}
