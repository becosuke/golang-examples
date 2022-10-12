package gateway

import (
	"context"
	"errors"
	domain "github.com/becosuke/golang-examples/kvstore/internal/domain/pack"
	syncmap2 "github.com/becosuke/golang-examples/kvstore/internal/pkg/syncmap"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
)

func NewRepository(config *config.Config, syncmap syncmap2.Syncmap, generator domain.Generator) domain.Repository {
	return &repositoryImpl{
		config:    config,
		store:     syncmap,
		generator: generator,
	}
}

type repositoryImpl struct {
	config    *config.Config
	store     syncmap2.Syncmap
	generator domain.Generator
}

func (impl *repositoryImpl) Read(_ context.Context, key *domain.Key) (*domain.Pack, error) {
	message, err := impl.store.Load(key.String())
	return impl.ToEntity(message), err
}

func (impl *repositoryImpl) Create(_ context.Context, value *domain.Value) (*domain.Key, error) {
	key := impl.generator.GenerateKey()
	_, loaded, err := impl.store.LoadOrStore(impl.ToSyncmapMessage(domain.NewPack(key, value)))
	if loaded {
		return nil, errors.New("already exists")
	}
	return key, err
}

func (impl *repositoryImpl) Update(_ context.Context, pack *domain.Pack) error {
	_, err := impl.store.Store(impl.ToSyncmapMessage(pack))
	return err
}

func (impl *repositoryImpl) Delete(_ context.Context, key *domain.Key) error {
	err := impl.store.Delete(key.String())
	return err
}

func (impl *repositoryImpl) ToEntity(m *syncmap2.Message) *domain.Pack {
	if m == nil {
		return &domain.Pack{}
	}
	return domain.NewPack(
		domain.NewKey(m.Key()),
		domain.NewValue(m.Value()),
	)
}

func (impl *repositoryImpl) ToSyncmapMessage(pack *domain.Pack) *syncmap2.Message {
	if pack == nil {
		return &syncmap2.Message{}
	}
	return syncmap2.NewMessage(
		pack.Key().String(),
		pack.Value().String(),
	)
}
