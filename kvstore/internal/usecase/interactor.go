package usecase

import (
	"context"
	"github.com/becosuke/golang-examples/kvstore/internal/adapters/repository"
	"github.com/becosuke/golang-examples/kvstore/internal/domain/entity"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
)

type Interactor interface {
	Create(context.Context, *entity.Pack) (*entity.Pack, error)
	Read(context.Context, *entity.Seal) (*entity.Pack, error)
	Update(context.Context, *entity.Pack) (*entity.Pack, error)
	Delete(context.Context, *entity.Seal) error
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

func (impl *interactorImpl) Create(ctx context.Context, pack *entity.Pack) (*entity.Pack, error) {
	result, err := impl.repository.Create(ctx, pack)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (impl *interactorImpl) Read(ctx context.Context, key *entity.Seal) (*entity.Pack, error) {
	result, err := impl.repository.Read(ctx, key)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (impl *interactorImpl) Update(ctx context.Context, pack *entity.Pack) (*entity.Pack, error) {
	result, err := impl.repository.Update(ctx, pack)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (impl *interactorImpl) Delete(ctx context.Context, key *entity.Seal) error {
	return impl.repository.Delete(ctx, key)
}
