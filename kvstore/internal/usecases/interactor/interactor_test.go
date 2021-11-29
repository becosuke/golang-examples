package interactor

import (
	"context"
	"github.com/becosuke/golang-examples/kvstore/internal/adapters/repository"
	"github.com/becosuke/golang-examples/kvstore/internal/domain/entity"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
	mock_repository "github.com/becosuke/golang-examples/kvstore/mock/adapters/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInteractor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockRepository(ctrl)
	SUT := NewInteractor(
		config.NewConfig(),
		mockRepository,
	)

	assert.Implements(t, (*repository.Repository)(nil), SUT)
}

func TestInteractorImpl_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockRepository(ctrl)
	SUT := NewInteractor(
		config.NewConfig(),
		mockRepository,
	)

	ctx := context.Background()
	pack := entity.NewPack("kkk", "vvv")
	mockRepository.EXPECT().Create(gomock.Eq(ctx), gomock.Eq(pack)).Return(pack, nil)

	res, err := SUT.Create(ctx, pack)

	assert.Equal(t, pack.Key, res.Key)
	assert.Equal(t, pack.Value, res.Value)
	assert.NoError(t, err)
}

func TestInteractorImpl_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockRepository(ctrl)
	SUT := NewInteractor(
		config.NewConfig(),
		mockRepository,
	)

	ctx := context.Background()
	pack := entity.NewPack("kkk", "vvv")
	packKey := entity.NewSeal("kkk")
	mockRepository.EXPECT().Read(gomock.Eq(ctx), gomock.Eq(packKey)).Return(pack, nil)

	res, err := SUT.Read(ctx, packKey)

	assert.Equal(t, pack.Key, res.Key)
	assert.Equal(t, pack.Value, res.Value)
	assert.NoError(t, err)
}

func TestInteractorImpl_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockRepository(ctrl)
	SUT := NewInteractor(
		config.NewConfig(),
		mockRepository,
	)

	ctx := context.Background()
	pack := entity.NewPack("kkk", "vvv")
	mockRepository.EXPECT().Update(gomock.Eq(ctx), gomock.Eq(pack)).Return(pack, nil)

	res, err := SUT.Update(ctx, pack)

	assert.Equal(t, pack.Key, res.Key)
	assert.Equal(t, pack.Value, res.Value)
	assert.NoError(t, err)

}

func TestInteractorImpl_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockRepository(ctrl)
	SUT := NewInteractor(
		config.NewConfig(),
		mockRepository,
	)

	ctx := context.Background()
	packKey := entity.NewSeal("kkk")
	mockRepository.EXPECT().Delete(gomock.Eq(ctx), gomock.Eq(packKey)).Return(nil)

	err := SUT.Delete(ctx, packKey)

	assert.NoError(t, err)
}
