package injection

import (
	"github.com/becosuke/golang-examples/hello/infrastructure/client"
)

func (c *Container) InjectHelloClient() client.Hello {
	if c.cache.HelloClient == nil {
		c.cache.HelloClient = client.NewHello()
	}
	return c.cache.HelloClient
}
