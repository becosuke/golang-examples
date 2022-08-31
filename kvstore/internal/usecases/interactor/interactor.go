package interactor

import (
	"context"
	domain "github.com/becosuke/golang-examples/kvstore/internal/domain/pack"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
)

func NewUsecase(config *config.Config, repository domain.Repository) domain.Usecase {
	return &usecaseImpl{
		config:     config,
		repository: repository,
	}
}

type usecaseImpl struct {
	config     *config.Config
	repository domain.Repository
}

func (impl *usecaseImpl) Read(ctx context.Context, key *domain.Key) (*domain.Pack, error) {
	result, err := impl.repository.Read(ctx, key)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (impl *usecaseImpl) Create(ctx context.Context, value *domain.Value) (*domain.Pack, error) {
	key, err := impl.repository.Create(ctx, value)
	if err != nil {
		return nil, err
	}
	return impl.Read(ctx, key)
}

func (impl *usecaseImpl) Update(ctx context.Context, pack *domain.Pack) (*domain.Pack, error) {
	err := impl.repository.Update(ctx, pack)
	if err != nil {
		return nil, err
	}
	return impl.Read(ctx, pack.Key())
}

func (impl *usecaseImpl) Delete(ctx context.Context, key *domain.Key) error {
	return impl.repository.Delete(ctx, key)
}
