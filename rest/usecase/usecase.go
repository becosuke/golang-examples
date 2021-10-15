package usecase

import (
	"context"

	"github.com/becosuke/golang-examples/rest/domain/entity"
	"github.com/becosuke/golang-examples/rest/domain/processor"
	"github.com/becosuke/golang-examples/rest/domain/repository"
	"github.com/becosuke/golang-examples/rest/registry/config"
)

type Usecase interface {
	Create(context.Context, string, string) ([]byte, error)
	Read(context.Context, string) ([]byte, error)
	Update(context.Context, string, string) ([]byte, error)
	Delete(context.Context, string) error
}

func NewUsecase(cfg *config.Config, repository repository.Repository, processor processor.Processor) Usecase {
	return &usecaseImpl{
		config:     cfg,
		repository: repository,
		processor:   processor,
	}
}

type usecaseImpl struct {
	config     *config.Config
	repository repository.Repository
	processor  processor.Processor
}

func (impl *usecaseImpl) Create(ctx context.Context, key, value string) ([]byte, error) {
	if v, b := impl.processor.Const(key); b {
		return entity.NewEntity(key, v).ConvertToJson()
	}
	m, err := impl.repository.Create(ctx, key, value)
	if err != nil {
		return []byte{}, err
	}
	return m.ConvertToJson()
}

func (impl *usecaseImpl) Read(ctx context.Context, key string) ([]byte, error) {
	if v, b := impl.processor.Const(key); b {
		return entity.NewEntity(key, v).ConvertToJson()
	}
	m, err := impl.repository.Read(ctx, key)
	if err != nil {
		return []byte{}, err
	}
	return m.ConvertToJson()
}

func (impl *usecaseImpl) Update(ctx context.Context, key, value string) ([]byte, error) {
	if v, b := impl.processor.Const(key); b {
		return entity.NewEntity(key, v).ConvertToJson()
	}
	m, err := impl.repository.Update(ctx, key, value)
	if err != nil {
		return []byte{}, err
	}
	return m.ConvertToJson()
}

func (impl *usecaseImpl) Delete(ctx context.Context, key string) error {
	return impl.repository.Delete(ctx, key)
}
