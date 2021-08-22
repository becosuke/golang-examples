package serviceimpl

import (
	"context"

	"github.com/becosuke/golang-examples/hello/config"
	"github.com/becosuke/golang-examples/hello/domain/entity"
	"github.com/becosuke/golang-examples/hello/domain/repository"
	"github.com/becosuke/golang-examples/hello/domain/service"
)

func NewHelloService(cfg *config.Config, helloRepository repository.HelloRepository) service.HelloService {
	return &helloServiceImpl{
		config:          cfg,
		helloRepository: helloRepository,
	}
}

type helloServiceImpl struct {
	config          *config.Config
	helloRepository repository.HelloRepository
}

func (si *helloServiceImpl) GetHello(ctx context.Context) *entity.Hello {
	return si.helloRepository.GetHello(ctx)
}
