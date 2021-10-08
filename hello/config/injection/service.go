package injection

import (
	"github.com/becosuke/golang-examples/hello/application/serviceimpl"
	"github.com/becosuke/golang-examples/hello/domain/service"
)

func (c *Container) InjectHelloService() service.Hello {
	if c.cache.HelloService == nil {
		c.cache.HelloService = serviceimpl.NewHello(c.InjectConfig(), c.InjectHelloRepository())
	}
	return c.cache.HelloService
}
