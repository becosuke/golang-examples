package injection

import (
	"github.com/becosuke/golang-examples/hello/application/controller"
)

func (c *Container) InjectHelloController() controller.Hello {
	if c.cache.HelloController == nil {
		c.cache.HelloController = controller.NewHello(c.InjectConfig(), c.InjectHelloService())
	}
	return c.cache.HelloController
}
