package usecases

import (
	"context"

	"github.com/becosuke/golang-examples/rest/internal/entities"
	"github.com/becosuke/golang-examples/rest/internal/interfaces/repository"
	"github.com/becosuke/golang-examples/rest/internal/registry/config"
)

type Interactor interface {
	Create(context.Context, *entities.Node) (*entities.Node, error)
	Read(context.Context, *entities.NodeKey) (*entities.Node, error)
	Update(context.Context, *entities.Node) (*entities.Node, error)
	Delete(context.Context, *entities.NodeKey) error
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

func (impl *interactorImpl) Create(ctx context.Context, node *entities.Node) (*entities.Node, error) {
	m, err := impl.repository.Create(ctx, node)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (impl *interactorImpl) Read(ctx context.Context, key *entities.NodeKey) (*entities.Node, error) {
	m, err := impl.repository.Read(ctx, key)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (impl *interactorImpl) Update(ctx context.Context, node *entities.Node) (*entities.Node, error) {
	m, err := impl.repository.Update(ctx, node)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (impl *interactorImpl) Delete(ctx context.Context, key *entities.NodeKey) error {
	return impl.repository.Delete(ctx, key)
}
