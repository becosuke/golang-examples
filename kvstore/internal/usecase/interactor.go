package usecase

import (
	"context"
	"github.com/becosuke/golang-examples/kvstore/internal/entity/kv"

	"github.com/becosuke/golang-examples/kvstore/internal/adapter/repository"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
)

type Interactor interface {
	Create(context.Context, *kv.Pack) (*kv.Pack, error)
	Read(context.Context, *kv.Seal) (*kv.Pack, error)
	Update(context.Context, *kv.Pack) (*kv.Pack, error)
	Delete(context.Context, *kv.Seal) error
}

func NewInteractor(config *config.Config, repository repository.Repository) Interactor {
	return &interactorImpl{
		config:     config,
		repository: repository,
	}
}

type interactorImpl struct {
	config     *config.Config
	repository repository.Repository
}

func (impl *interactorImpl) Create(ctx context.Context, pack *kv.Pack) (*kv.Pack, error) {
	m, err := impl.repository.Create(ctx, pack)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (impl *interactorImpl) Read(ctx context.Context, key *kv.Seal) (*kv.Pack, error) {
	m, err := impl.repository.Read(ctx, key)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (impl *interactorImpl) Update(ctx context.Context, pack *kv.Pack) (*kv.Pack, error) {
	m, err := impl.repository.Update(ctx, pack)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (impl *interactorImpl) Delete(ctx context.Context, key *kv.Seal) error {
	return impl.repository.Delete(ctx, key)
}
