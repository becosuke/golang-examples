package usecases

import (
	"context"
	"github.com/becosuke/golang-examples/grpc/internal/entities/node"

	"github.com/becosuke/golang-examples/grpc/internal/adapters/repository"
	"github.com/becosuke/golang-examples/grpc/internal/registry/config"
)

type Interactor interface {
	Create(context.Context, *node.KeyValueNode) (*node.KeyValueNode, error)
	Read(context.Context, *node.KeyNode) (*node.KeyValueNode, error)
	Update(context.Context, *node.KeyValueNode) (*node.KeyValueNode, error)
	Delete(context.Context, *node.KeyNode) error
}

func NewInteractor(config *config.Config, repository repository.Repository) Interactor {
	return &interactorImpl{
		config:     config,
		repository: repository,
	}
}

type interactorImpl struct {
	config     *config.Config
	repository repository.Repository
}

func (impl *interactorImpl) Create(ctx context.Context, node *node.KeyValueNode) (*node.KeyValueNode, error) {
	m, err := impl.repository.Create(ctx, node)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (impl *interactorImpl) Read(ctx context.Context, key *node.KeyNode) (*node.KeyValueNode, error) {
	m, err := impl.repository.Read(ctx, key)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (impl *interactorImpl) Update(ctx context.Context, node *node.KeyValueNode) (*node.KeyValueNode, error) {
	m, err := impl.repository.Update(ctx, node)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (impl *interactorImpl) Delete(ctx context.Context, key *node.KeyNode) error {
	return impl.repository.Delete(ctx, key)
}
