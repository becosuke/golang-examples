package repository

import (
	"context"
	"errors"
	"github.com/becosuke/golang-examples/kvstore/internal/domain/pack"
	"github.com/becosuke/golang-examples/kvstore/internal/drivers/syncmap"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
	"github.com/google/uuid"
)

func NewRepository(config *config.Config, syncmap syncmap.Syncmap) pack.Repository {
	return &repositoryImpl{
		config: config,
		store:  syncmap,
	}
}

type repositoryImpl struct {
	config *config.Config
	store  syncmap.Syncmap
}

func (impl *repositoryImpl) Read(_ context.Context, key *pack.Key) (*pack.Pack, error) {
	message, err := impl.store.Load(key.String())
	return impl.ToEntity(message), err
}

func (impl *repositoryImpl) Create(_ context.Context, pack *pack.Pack) error {
	_, loaded, err := impl.store.LoadOrStore(impl.ToSyncmapMessage(pack))
	if loaded {
		return errors.New("already exists")
	}
	return err
}

func (impl *repositoryImpl) Update(_ context.Context, pack *pack.Pack) error {
	_, err := impl.store.Store(impl.ToSyncmapMessage(pack))
	return err
}

func (impl *repositoryImpl) Delete(_ context.Context, key *pack.Key) error {
	err := impl.store.Delete(key.String())
	return err
}

func (impl *repositoryImpl) ToEntity(m *syncmap.Message) *pack.Pack {
	if m == nil {
		return &pack.Pack{}
	}
	return pack.NewPack(
		pack.NewKey(uuid.MustParse(m.Key())),
		pack.NewValue(m.Value()),
	)
}

func (impl *repositoryImpl) ToSyncmapMessage(pack *pack.Pack) *syncmap.Message {
	if pack == nil {
		return &syncmap.Message{}
	}
	return syncmap.NewMessage(
		pack.Key().String(),
		pack.Value().String(),
	)
}
