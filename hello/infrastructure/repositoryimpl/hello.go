package repositoryimpl

import (
	"context"

	"github.com/becosuke/golang-examples/hello/config"
	"github.com/becosuke/golang-examples/hello/domain/entity"
	"github.com/becosuke/golang-examples/hello/domain/repository"
	"github.com/becosuke/golang-examples/hello/infrastructure/client"
)

func NewHello(cfg *config.Config, clt client.Hello) repository.Hello {
	return &helloImpl{
		config: cfg,
		client: clt,
	}
}

type helloImpl struct {
	config *config.Config
	client client.Hello
}

func (ri *helloImpl) Get(ctx context.Context) *entity.Hello {
	return &entity.Hello{
		Message: ri.client.Get(),
	}
}
