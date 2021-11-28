package repository

import (
	"context"
	"errors"
	"github.com/becosuke/golang-examples/kvstore/internal/domain/entity"
	"github.com/becosuke/golang-examples/kvstore/internal/drivers/syncmap"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
)

type Repository interface {
	Create(context.Context, *entity.Pack) (*entity.Pack, error)
	Read(context.Context, *entity.Seal) (*entity.Pack, error)
	Update(context.Context, *entity.Pack) (*entity.Pack, error)
	Delete(context.Context, *entity.Seal) error
}

func NewRepository(config *config.Config, syncmap syncmap.Syncmap) Repository {
	return &repositoryImpl{
		config: config,
		store:  syncmap,
	}
}

type repositoryImpl struct {
	config *config.Config
	store  syncmap.Syncmap
}

func (s *repositoryImpl) Create(ctx context.Context, pack *entity.Pack) (*entity.Pack, error) {
	message, loaded, err := s.store.LoadOrStore(pack.Key, pack.Value)
	if loaded {
		return nil, errors.New("already exists")
	}
	return message.ConvertToEntity(), err
}

func (s *repositoryImpl) Read(ctx context.Context, packKey *entity.Seal) (*entity.Pack, error) {
	message, err := s.store.Load(packKey.Key)
	return message.ConvertToEntity(), err
}

func (s *repositoryImpl) Update(ctx context.Context, pack *entity.Pack) (*entity.Pack, error) {
	message, err := s.store.Store(pack.Key, pack.Value)
	return message.ConvertToEntity(), err
}

func (s *repositoryImpl) Delete(ctx context.Context, packKey *entity.Seal) error {
	err := s.store.Delete(packKey.Key)
	return err
}
