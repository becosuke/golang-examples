package injection

import (
	"github.com/becosuke/golang-examples/hello/config"
)

func (c *Container) InjectConfig() *config.Config {
	if c.cache.Config == nil {
		cfg := config.Config{
			Environment: "DEVELOPMENT",
			LogLevel: "INFO",
		}
		c.cache.Config = &cfg
	}
	return c.cache.Config
}
