package injection

import (
	"github.com/becosuke/golang-examples/hello/presentation/controller"
)

func (c *Container) InjectHelloController() controller.HelloController {
	if c.cache.HelloController == nil {
		helloController := controller.NewHelloController(c.InjectConfig(), c.InjectHelloService())
		c.cache.HelloController = helloController
	}
	return c.cache.HelloController
}
