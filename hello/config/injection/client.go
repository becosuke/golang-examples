package injection

import (
	"github.com/becosuke/golang-examples/hello/infrastructure/client"
)

func (c *Container) InjectClient() *client.Client {
	if c.cache.Client == nil {
		clt := client.Client{}
		c.cache.Client = &clt
	}
	return c.cache.Client
}