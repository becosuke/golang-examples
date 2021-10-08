package serviceimpl

import (
	"context"

	"github.com/becosuke/golang-examples/hello/config"
	"github.com/becosuke/golang-examples/hello/domain/entity"
	"github.com/becosuke/golang-examples/hello/domain/repository"
	"github.com/becosuke/golang-examples/hello/domain/service"
)

func NewHello(cfg *config.Config, repository repository.Hello) service.Hello {
	return &helloImpl{
		config:     cfg,
		repository: repository,
	}
}

type helloImpl struct {
	config     *config.Config
	repository repository.Hello
}

func (si *helloImpl) Get(ctx context.Context) *entity.Hello {
	return si.repository.Get(ctx)
}
