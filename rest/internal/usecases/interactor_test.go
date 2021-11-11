package usecases

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/becosuke/golang-examples/rest/internal/entities"
	"github.com/becosuke/golang-examples/rest/internal/interfaces/repository"
	"github.com/becosuke/golang-examples/rest/internal/registry/config"
	mock_repository "github.com/becosuke/golang-examples/rest/mocks/interfaces/repository"
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
	node := entities.NewNode("kkk", "vvv")
	mockRepository.EXPECT().Create(gomock.Eq(ctx), gomock.Eq(node)).Return(node, nil)

	res, err := SUT.Create(ctx, node)

	assert.Equal(t, node.Key, res.Key)
	assert.Equal(t, node.Value, res.Value)
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
	node := entities.NewNode("kkk", "vvv")
	nodeKey := entities.NewNodeKey("kkk")
	mockRepository.EXPECT().Read(gomock.Eq(ctx), gomock.Eq(nodeKey)).Return(node, nil)

	res, err := SUT.Read(ctx, nodeKey)

	assert.Equal(t, node.Key, res.Key)
	assert.Equal(t, node.Value, res.Value)
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
	node := entities.NewNode("kkk", "vvv")
	mockRepository.EXPECT().Update(gomock.Eq(ctx), gomock.Eq(node)).Return(node, nil)

	res, err := SUT.Update(ctx, node)

	assert.Equal(t, node.Key, res.Key)
	assert.Equal(t, node.Value, res.Value)
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
	nodeKey := entities.NewNodeKey("kkk")
	mockRepository.EXPECT().Delete(gomock.Eq(ctx), gomock.Eq(nodeKey)).Return(nil)

	err := SUT.Delete(ctx, nodeKey)

	assert.NoError(t, err)
}
