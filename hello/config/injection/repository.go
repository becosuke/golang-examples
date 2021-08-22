package injection

import (
	"github.com/becosuke/golang-examples/hello/domain/repository"
	"github.com/becosuke/golang-examples/hello/infrastructure/repositoryimpl"
)

func (c *Container) InjectHelloRepository() repository.HelloRepository {
	if c.cache.HelloRepository == nil {
		helloRepository := repositoryimpl.NewHelloRepository(c.InjectConfig(), c.InjectClient())
		c.cache.HelloRepository = helloRepository
	}
	return c.cache.HelloRepository
}
