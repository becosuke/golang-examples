package injector

import (
	"github.com/becosuke/golang-examples/rest/application/controller"
	"github.com/becosuke/golang-examples/rest/domain/processor"
	"github.com/becosuke/golang-examples/rest/domain/repository"
	"github.com/becosuke/golang-examples/rest/infrastructure/repositoryimpl"
	"github.com/becosuke/golang-examples/rest/infrastructure/syncmap"
	"github.com/becosuke/golang-examples/rest/registry/config"
	"github.com/becosuke/golang-examples/rest/usecase"
)

type Container struct {
	cache struct {
		Config     *config.Config
		Controller controller.Controller
		Usecase    usecase.Usecase
		Repository repository.Repository
		Processor  processor.Processor
		SyncMap    syncmap.SyncMap
	}
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) InjectConfig() *config.Config {
	if c.cache.Config == nil {
		c.cache.Config = &config.Config{
			Environment: "DEVELOPMENT",
			LogLevel:    "INFO",
			HttpPort:    "8080",
			Constant: map[string]string{
				config.ConstKey: config.ConstValue,
			},
		}
	}
	return c.cache.Config
}

func (c *Container) InjectController() controller.Controller {
	if c.cache.Controller == nil {
		c.cache.Controller = controller.NewController(c.InjectConfig(), c.InjectUsecase())
	}
	return c.cache.Controller
}

func (c *Container) InjectUsecase() usecase.Usecase {
	if c.cache.Usecase == nil {
		c.cache.Usecase = usecase.NewUsecase(c.InjectConfig(), c.InjectRepository(), c.InjectProcessor())
	}
	return c.cache.Usecase
}

func (c *Container) InjectRepository() repository.Repository {
	if c.cache.Repository == nil {
		c.cache.Repository = repositoryimpl.NewRepository(c.InjectConfig(), c.InjectSyncMap())
	}
	return c.cache.Repository
}

func (c *Container) InjectProcessor() processor.Processor {
	if c.cache.Processor == nil {
		c.cache.Processor = processor.NewProcessor(c.InjectConfig())
	}
	return c.cache.Processor
}

func (c *Container) InjectSyncMap() syncmap.SyncMap {
	if c.cache.SyncMap == nil {
		c.cache.SyncMap = syncmap.NewSyncMap()
	}
	return c.cache.SyncMap
}
