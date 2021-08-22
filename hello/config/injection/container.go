package injection

import (
	"github.com/becosuke/golang-examples/hello/config"
	"github.com/becosuke/golang-examples/hello/domain/repository"
	"github.com/becosuke/golang-examples/hello/domain/service"
	"github.com/becosuke/golang-examples/hello/infrastructure/client"
	"github.com/becosuke/golang-examples/hello/presentation/controller"
)

type Container struct {
	cache struct {
		Client          *client.Client
		Config          *config.Config
		HelloRepository repository.HelloRepository
		HelloService    service.HelloService
		HelloController controller.HelloController
	}
}
