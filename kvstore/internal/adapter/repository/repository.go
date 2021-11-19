package repository

import (
	"context"
	"errors"
	"github.com/becosuke/golang-examples/kvstore/internal/entity/kv"
	"github.com/becosuke/golang-examples/kvstore/internal/infrastructure/syncmap"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
)

type Repository interface {
	Create(context.Context, *kv.Pack) (*kv.Pack, error)
	Read(context.Context, *kv.Seal) (*kv.Pack, error)
	Update(context.Context, *kv.Pack) (*kv.Pack, error)
	Delete(context.Context, *kv.Seal) error
}

func NewRepository(config *config.Config, syncMap syncmap.SyncMap) Repository {
	return &repositoryImpl{
		config:  config,
		syncMap: syncMap,
	}
}

type repositoryImpl struct {
	config  *config.Config
	syncMap syncmap.SyncMap
}

func (s *repositoryImpl) Create(ctx context.Context, pack *kv.Pack) (*kv.Pack, error) {
	message, loaded, err := s.syncMap.LoadOrStore(pack.Key, pack.Value)
	if loaded {
		return nil, errors.New("already exists")
	}
	return message.ConvertToEntity(), err
}

func (s *repositoryImpl) Read(ctx context.Context, packKey *kv.Seal) (*kv.Pack, error) {
	message, err := s.syncMap.Load(packKey.Key)
	return message.ConvertToEntity(), err
}

func (s *repositoryImpl) Update(ctx context.Context, pack *kv.Pack) (*kv.Pack, error) {
	message, err := s.syncMap.Store(pack.Key, pack.Value)
	return message.ConvertToEntity(), err
}

func (s *repositoryImpl) Delete(ctx context.Context, packKey *kv.Seal) error {
	err := s.syncMap.Delete(packKey.Key)
	return err
}
