package injection

import (
	"github.com/becosuke/golang-examples/hello/domain/repository"
	"github.com/becosuke/golang-examples/hello/infrastructure/repositoryimpl"
)

func (c *Container) InjectHelloRepository() repository.Hello {
	if c.cache.HelloRepository == nil {
		c.cache.HelloRepository = repositoryimpl.NewHello(c.InjectConfig(), c.InjectHelloClient())
	}
	return c.cache.HelloRepository
}
