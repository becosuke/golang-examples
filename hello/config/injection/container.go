package injection

import (
	"github.com/becosuke/golang-examples/hello/application/controller"
	"github.com/becosuke/golang-examples/hello/config"
	"github.com/becosuke/golang-examples/hello/domain/repository"
	"github.com/becosuke/golang-examples/hello/domain/service"
	"github.com/becosuke/golang-examples/hello/infrastructure/client"
)

type Container struct {
	cache struct {
		Config          *config.Config
		HelloClient     client.Hello
		HelloRepository repository.Hello
		HelloService    service.Hello
		HelloController controller.Hello
	}
}
