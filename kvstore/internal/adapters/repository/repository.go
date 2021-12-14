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

func (ri *repositoryImpl) Create(_ context.Context, pack *entity.Pack) (*entity.Pack, error) {
	message, loaded, err := ri.store.LoadOrStore(ri.ToSyncmapMessage(pack))
	if loaded {
		return nil, errors.New("already exists")
	}
	return ri.ToEntity(message), err
}

func (ri *repositoryImpl) Read(_ context.Context, seal *entity.Seal) (*entity.Pack, error) {
	message, err := ri.store.Load(ri.ToSyncmapKey(seal))
	return ri.ToEntity(message), err
}

func (ri *repositoryImpl) Update(_ context.Context, pack *entity.Pack) (*entity.Pack, error) {
	message, err := ri.store.Store(ri.ToSyncmapMessage(pack))
	return ri.ToEntity(message), err
}

func (ri *repositoryImpl) Delete(_ context.Context, seal *entity.Seal) error {
	err := ri.store.Delete(ri.ToSyncmapKey(seal))
	return err
}

func (ri *repositoryImpl) ToEntity(m *syncmap.Message) *entity.Pack {
	if m == nil {
		return &entity.Pack{}
	}
	return &entity.Pack{
		Key:   m.GetKey().String(),
		Value: m.GetValue().String(),
	}
}

func (ri *repositoryImpl) ToSyncmapMessage(pack *entity.Pack) *syncmap.Message {
	if pack == nil {
		return &syncmap.Message{}
	}
	return &syncmap.Message{
		Key:   syncmap.Key(pack.GetKey()),
		Value: syncmap.Value(pack.GetValue()),
	}
}

func (ri *repositoryImpl) ToSyncmapKey(seal *entity.Seal) syncmap.Key {
	if seal == nil {
		return ""
	}
	return syncmap.Key(seal.GetKey())
}
