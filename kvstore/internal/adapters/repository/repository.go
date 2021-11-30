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

func (ri *repositoryImpl) Create(ctx context.Context, pack *entity.Pack) (*entity.Pack, error) {
	message, loaded, err := ri.store.LoadOrStore(ri.ToMessage(pack))
	if loaded {
		return nil, errors.New("already exists")
	}
	return ri.ToEntity(message), err
}

func (ri *repositoryImpl) Read(ctx context.Context, seal *entity.Seal) (*entity.Pack, error) {
	message, err := ri.store.Load(ri.ToMessageKey(seal))
	return ri.ToEntity(message), err
}

func (ri *repositoryImpl) Update(ctx context.Context, pack *entity.Pack) (*entity.Pack, error) {
	message, err := ri.store.Store(ri.ToMessage(pack))
	return ri.ToEntity(message), err
}

func (ri *repositoryImpl) Delete(ctx context.Context, seal *entity.Seal) error {
	err := ri.store.Delete(ri.ToMessageKey(seal))
	return err
}

func (ri *repositoryImpl) ToEntity(m *syncmap.Message) *entity.Pack {
	if m == nil {
		return &entity.Pack{}
	}
	return &entity.Pack{
		Key:   m.GetKey(),
		Value: m.GetValue(),
	}
}

func (ri *repositoryImpl) ToMessage(pack *entity.Pack) *syncmap.Message {
	if pack == nil {
		return &syncmap.Message{}
	}
	return &syncmap.Message{
		Key:   pack.GetKey(),
		Value: pack.GetValue(),
	}
}

func (ri *repositoryImpl) ToMessageKey(seal *entity.Seal) *syncmap.MessageKey {
	if seal == nil {
		return &syncmap.MessageKey{}
	}
	return &syncmap.MessageKey{
		Key: seal.GetKey(),
	}
}
