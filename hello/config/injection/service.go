package injection

import (
	"github.com/becosuke/golang-examples/hello/application/serviceimpl"
	"github.com/becosuke/golang-examples/hello/domain/service"
)

func (c *Container) InjectHelloService() service.HelloService {
	if c.cache.HelloService == nil {
		helloService := serviceimpl.NewHelloService(c.InjectConfig(), c.InjectHelloRepository())
		c.cache.HelloService = helloService
	}
	return c.cache.HelloService
}
