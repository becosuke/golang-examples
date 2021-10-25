package injector

import (
	"github.com/becosuke/golang-examples/rest/infrastructure/syncmap"
	"sync"

	"github.com/becosuke/golang-examples/rest/application/usecase"
	"github.com/becosuke/golang-examples/rest/domain/processor"
	"github.com/becosuke/golang-examples/rest/domain/repository"
	"github.com/becosuke/golang-examples/rest/infrastructure/memorymap"
	"github.com/becosuke/golang-examples/rest/infrastructure/repositoryimpl"
	"github.com/becosuke/golang-examples/rest/presentation/controller"
	"github.com/becosuke/golang-examples/rest/presentation/handler"
	"github.com/becosuke/golang-examples/rest/registry/config"
)

type Injector interface {
	InjectConfig() *config.Config
	InjectHandler() handler.Handler
	InjectController() controller.Controller
	InjectUsecase() usecase.Usecase
	InjectRepository() repository.Repository
	InjectProcessor() processor.Processor
	InjectSyncMap() memorymap.MemoryMap
}

func NewInjector() Injector {
	return &injectorImpl{}
}

type injectorImpl struct {
	container struct {
		Config     *config.Config
		Handler    handler.Handler
		Controller controller.Controller
		Usecase    usecase.Usecase
		Repository repository.Repository
		Processor processor.Processor
		SyncMap   memorymap.MemoryMap
	}
	store sync.Map
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

func (impl *injectorImpl) InjectController() controller.Controller {
	actual, _ := impl.store.LoadOrStore("controller", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Controller = controller.NewController(impl.InjectConfig(), impl.InjectUsecase())
		})
	}
	return impl.container.Controller
}

func (impl *injectorImpl) InjectUsecase() usecase.Usecase {
	actual, _ := impl.store.LoadOrStore("usecase", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Usecase = usecase.NewUsecase(impl.InjectConfig(), impl.InjectRepository(), impl.InjectProcessor())
		})
	}
	return impl.container.Usecase
}

func (impl *injectorImpl) InjectRepository() repository.Repository {
	actual, _ := impl.store.LoadOrStore("repository", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Repository = repositoryimpl.NewRepository(impl.InjectConfig(), impl.InjectSyncMap())
		})
	}
	return impl.container.Repository
}

func (impl *injectorImpl) InjectProcessor() processor.Processor {
	actual, _ := impl.store.LoadOrStore("processor", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Processor = processor.NewProcessor(impl.InjectConfig())
		})
	}
	return impl.container.Processor
}

func (impl *injectorImpl) InjectSyncMap() memorymap.MemoryMap {
	actual, _ := impl.store.LoadOrStore("syncmap", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.SyncMap = memorymap.NewMemoryMap(syncmap.NewSyncMap())
		})
	}
	return impl.container.SyncMap
}
