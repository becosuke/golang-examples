package injector

import (
	"sync"

	"github.com/becosuke/golang-examples/rest/internal/infrastructure/server"
	"github.com/becosuke/golang-examples/rest/internal/infrastructure/server/handler"
	"github.com/becosuke/golang-examples/rest/internal/infrastructure/syncmap"
	"github.com/becosuke/golang-examples/rest/internal/interfaces/controllers"
	"github.com/becosuke/golang-examples/rest/internal/interfaces/repository"
	"github.com/becosuke/golang-examples/rest/internal/registry/config"
	"github.com/becosuke/golang-examples/rest/internal/usecases"
)

type Injector interface {
	InjectConfig() *config.Config
	InjectServer() server.Server
	InjectHandler() handler.Handler
	InjectController() controllers.Controller
	InjectUsecase() usecases.Interactor
	InjectRepository() repository.Repository
	InjectSyncMap() syncmap.SyncMap
}

func NewInjector() Injector {
	return &injectorImpl{}
}

type injectorImpl struct {
	container struct {
		Config     *config.Config
		Server     server.Server
		Handler    handler.Handler
		Controller controllers.Controller
		Usecase    usecases.Interactor
		Repository repository.Repository
		SyncMap    syncmap.SyncMap
	}
	store sync.Map
}

func (impl *injectorImpl) InjectServer() server.Server {
	actual, _ := impl.store.LoadOrStore("server", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Server = server.NewServer(impl.InjectConfig(), impl.InjectHandler())
		})
	}
	return impl.container.Server
}

func (impl *injectorImpl) InjectConfig() *config.Config {
	actual, _ := impl.store.LoadOrStore("config", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Config = config.NewConfig()
		})
	}
	return impl.container.Config
}

func (impl *injectorImpl) InjectHandler() handler.Handler {
	actual, _ := impl.store.LoadOrStore("handler", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Handler = handler.NewHandler(impl.InjectConfig(), impl.InjectController())
		})
	}
	return impl.container.Handler
}

func (impl *injectorImpl) InjectController() controllers.Controller {
	actual, _ := impl.store.LoadOrStore("controller", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Controller = controllers.NewController(impl.InjectConfig(), impl.InjectUsecase())
		})
	}
	return impl.container.Controller
}

func (impl *injectorImpl) InjectUsecase() usecases.Interactor {
	actual, _ := impl.store.LoadOrStore("interactor", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Usecase = usecases.NewInteractor(impl.InjectConfig(), impl.InjectRepository())
		})
	}
	return impl.container.Usecase
}

func (impl *injectorImpl) InjectRepository() repository.Repository {
	actual, _ := impl.store.LoadOrStore("repository", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Repository = repository.NewRepository(impl.InjectConfig(), impl.InjectSyncMap())
		})
	}
	return impl.container.Repository
}

func (impl *injectorImpl) InjectSyncMap() syncmap.SyncMap {
	actual, _ := impl.store.LoadOrStore("syncmap", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.SyncMap = syncmap.NewSyncMap()
		})
	}
	return impl.container.SyncMap
}
