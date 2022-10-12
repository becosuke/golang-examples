package interactor

import (
	"context"
	domain "github.com/becosuke/golang-examples/kvstore/internal/domain/pack"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
	mockDomain "github.com/becosuke/golang-examples/kvstore/mocks/domain/pack"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInteractor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newMockRepository := mockDomain.NewMockRepository(ctrl)
	SUT := NewUsecase(
		config.NewConfig(),
		newMockRepository,
	)

	assert.Implements(t, (*domain.Usecase)(nil), SUT)
}

func TestInteractorImpl_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newMockRepository := mockDomain.NewMockRepository(ctrl)
	SUT := NewUsecase(
		config.NewConfig(),
		newMockRepository,
	)

	ctx := context.Background()
	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	newKey := domain.NewKey(keyString)
	newPack := domain.NewPack(newKey, domain.NewValue(valueString))
	newMockRepository.EXPECT().Read(gomock.Eq(ctx), gomock.Eq(newKey)).Return(newPack, nil)

	res, err := SUT.Read(ctx, newKey)
	assert.NoError(t, err)
	assert.Equal(t, newPack, res)
}

func TestInteractorImpl_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newMockRepository := mockDomain.NewMockRepository(ctrl)
	SUT := NewUsecase(
		config.NewConfig(),
		newMockRepository,
	)

	ctx := context.Background()
	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	newKey := domain.NewKey(keyString)
	newValue := domain.NewValue(valueString)
	newPack := domain.NewPack(newKey, newValue)
	newMockRepository.EXPECT().Create(gomock.Eq(ctx), gomock.Eq(newValue)).Return(newKey, nil)
	newMockRepository.EXPECT().Read(gomock.Eq(ctx), gomock.Eq(newKey)).Return(newPack, nil)

	res, err := SUT.Create(ctx, newValue)
	assert.NoError(t, err)
	assert.Equal(t, newPack, res)
}

func TestInteractorImpl_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newMockRepository := mockDomain.NewMockRepository(ctrl)
	SUT := NewUsecase(
		config.NewConfig(),
		newMockRepository,
	)

	ctx := context.Background()
	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	newKey := domain.NewKey(keyString)
	newPack := domain.NewPack(newKey, domain.NewValue(valueString))
	newMockRepository.EXPECT().Update(gomock.Eq(ctx), gomock.Eq(newPack)).Return(nil)
	newMockRepository.EXPECT().Read(gomock.Eq(ctx), gomock.Eq(newKey)).Return(newPack, nil)

	res, err := SUT.Update(ctx, newPack)
	assert.NoError(t, err)
	assert.Equal(t, newPack, res)
}

func TestInteractorImpl_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newMockRepository := mockDomain.NewMockRepository(ctrl)
	SUT := NewUsecase(
		config.NewConfig(),
		newMockRepository,
	)

	ctx := context.Background()
	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	newKey := domain.NewKey(keyString)
	newMockRepository.EXPECT().Delete(gomock.Eq(ctx), gomock.Eq(newKey)).Return(nil)

	err := SUT.Delete(ctx, newKey)
	assert.NoError(t, err)
}
