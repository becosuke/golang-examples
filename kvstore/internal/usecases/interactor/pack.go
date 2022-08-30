package interactor

import (
	"context"
	"github.com/becosuke/golang-examples/kvstore/internal/domain/pack"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
)

func NewUsecase(config *config.Config, repository pack.Repository) pack.Usecase {
	return &usecaseImpl{
		config:     config,
		repository: repository,
	}
}

type usecaseImpl struct {
	config     *config.Config
	repository pack.Repository
}

func (impl *usecaseImpl) Read(ctx context.Context, key *pack.Key) (*pack.Pack, error) {
	result, err := impl.repository.Read(ctx, key)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (impl *usecaseImpl) Create(ctx context.Context, domain *pack.Pack) error {
	return impl.repository.Create(ctx, domain)
}

func (impl *usecaseImpl) Update(ctx context.Context, domain *pack.Pack) error {
	return impl.repository.Update(ctx, domain)
}

func (impl *usecaseImpl) Delete(ctx context.Context, key *pack.Key) error {
	return impl.repository.Delete(ctx, key)
}
