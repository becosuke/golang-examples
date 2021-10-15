package repositoryimpl

import (
	"context"

	"github.com/becosuke/golang-examples/rest/domain/entity"
	"github.com/becosuke/golang-examples/rest/domain/repository"
	"github.com/becosuke/golang-examples/rest/infrastructure/syncmap"
	"github.com/becosuke/golang-examples/rest/registry/config"
)

func NewRepository(conf *config.Config, syncMap syncmap.SyncMap) repository.Repository {
	return &repositoryImpl{
		config:  conf,
		syncMap: syncMap,
	}
}

type repositoryImpl struct {
	config  *config.Config
	syncMap syncmap.SyncMap
}

func (impl *repositoryImpl) Create(ctx context.Context, key, value string) (*entity.Entity, error) {
	message, err := impl.syncMap.LoadOrStore(key, value)
	return message.ConvertToEntity(), err
}

func (impl *repositoryImpl) Read(ctx context.Context, key string) (*entity.Entity, error) {
	message, err := impl.syncMap.Load(key)
	return message.ConvertToEntity(), err
}

func (impl *repositoryImpl) Update(ctx context.Context, key, value string) (*entity.Entity, error) {
	message, err := impl.syncMap.Store(key, value)
	return message.ConvertToEntity(), err
}

func (impl *repositoryImpl) Delete(ctx context.Context, key string) error {
	err := impl.syncMap.Delete(key)
	return err
}
