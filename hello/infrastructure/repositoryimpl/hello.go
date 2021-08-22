package repositoryimpl

import (
	"context"

	"github.com/becosuke/golang-examples/hello/config"
	"github.com/becosuke/golang-examples/hello/domain/entity"
	"github.com/becosuke/golang-examples/hello/domain/repository"
	"github.com/becosuke/golang-examples/hello/infrastructure/client"
)

func NewHelloRepository(cfg *config.Config, clt *client.Client) repository.HelloRepository {
	return &helloRepositoryImpl{
		config: cfg,
		client: clt,
	}
}

type helloRepositoryImpl struct {
	config *config.Config
	client *client.Client
}

func (ri *helloRepositoryImpl) GetHello(ctx context.Context) *entity.Hello {
	ri.client.Get()
	return &entity.Hello{
		Message: ri.client.Get(),
	}
}