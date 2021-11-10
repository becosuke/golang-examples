package repository

import (
	"context"
	"errors"

	"github.com/becosuke/golang-examples/rest/internal/entities"
	"github.com/becosuke/golang-examples/rest/internal/infrastructure/syncmap"
	"github.com/becosuke/golang-examples/rest/internal/registry/config"
)

type Repository interface {
	Create(context.Context, *entities.Node) (*entities.Node, error)
	Read(context.Context, *entities.NodeKey) (*entities.Node, error)
	Update(context.Context, *entities.Node) (*entities.Node, error)
	Delete(context.Context, *entities.NodeKey) error
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

func (s *repositoryImpl) Create(ctx context.Context, node *entities.Node) (*entities.Node, error) {
	message, loaded, err := s.syncMap.LoadOrStore(node.Key, node.Value)
	if loaded {
		return nil, errors.New("already exists")
	}
	return message.ConvertToEntity(), err
}

func (s *repositoryImpl) Read(ctx context.Context, nodeKey *entities.NodeKey) (*entities.Node, error) {
	message, err := s.syncMap.Load(nodeKey.Key)
	return message.ConvertToEntity(), err
}

func (s *repositoryImpl) Update(ctx context.Context, node *entities.Node) (*entities.Node, error) {
	message, err := s.syncMap.Store(node.Key, node.Value)
	return message.ConvertToEntity(), err
}

func (s *repositoryImpl) Delete(ctx context.Context, nodeKey *entities.NodeKey) error {
	err := s.syncMap.Delete(nodeKey.Key)
	return err
}
