package repository

import (
	"context"
	"errors"
	"github.com/becosuke/golang-examples/grpc/internal/entities/node"

	"github.com/becosuke/golang-examples/grpc/internal/infrastructure/syncmap"
	"github.com/becosuke/golang-examples/grpc/internal/registry/config"
)

type Repository interface {
	Create(context.Context, *node.KeyValueNode) (*node.KeyValueNode, error)
	Read(context.Context, *node.KeyNode) (*node.KeyValueNode, error)
	Update(context.Context, *node.KeyValueNode) (*node.KeyValueNode, error)
	Delete(context.Context, *node.KeyNode) error
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

func (s *repositoryImpl) Create(ctx context.Context, node *node.KeyValueNode) (*node.KeyValueNode, error) {
	message, loaded, err := s.syncMap.LoadOrStore(node.Key, node.Value)
	if loaded {
		return nil, errors.New("already exists")
	}
	return message.ConvertToEntity(), err
}

func (s *repositoryImpl) Read(ctx context.Context, nodeKey *node.KeyNode) (*node.KeyValueNode, error) {
	message, err := s.syncMap.Load(nodeKey.Key)
	return message.ConvertToEntity(), err
}

func (s *repositoryImpl) Update(ctx context.Context, node *node.KeyValueNode) (*node.KeyValueNode, error) {
	message, err := s.syncMap.Store(node.Key, node.Value)
	return message.ConvertToEntity(), err
}

func (s *repositoryImpl) Delete(ctx context.Context, nodeKey *node.KeyNode) error {
	err := s.syncMap.Delete(nodeKey.Key)
	return err
}
