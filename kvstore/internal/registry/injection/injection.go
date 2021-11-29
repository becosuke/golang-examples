package injection

import (
	"github.com/becosuke/golang-examples/kvstore/internal/adapters/boundary"
	"github.com/becosuke/golang-examples/kvstore/internal/adapters/controller"
	"github.com/becosuke/golang-examples/kvstore/internal/adapters/repository"
	"github.com/becosuke/golang-examples/kvstore/internal/drivers/grpcserver"
	"github.com/becosuke/golang-examples/kvstore/internal/drivers/logger"
	"github.com/becosuke/golang-examples/kvstore/internal/drivers/syncmap"
	"github.com/becosuke/golang-examples/kvstore/internal/pb"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
	"github.com/becosuke/golang-examples/kvstore/internal/usecases/interactor"
	"google.golang.org/grpc"
	"log"
	"sync"
)

type Injection interface {
	InjectConfig() *config.Config
	InjectLogger() logger.Logger
	InjectGrpcServer() *grpc.Server
	InjectController() pb.KVStoreServiceServer
	InjectUsecase() interactor.Interactor
	InjectBoundary() boundary.Boundary
	InjectRepository() repository.Repository
	InjectSyncmap() syncmap.Syncmap
}

func NewInjection(serviceName, version string) Injection {
	return &injectionImpl{serviceName: serviceName, version: version}
}

type injectionImpl struct {
	container struct {
		Config               *config.Config
		Logger               logger.Logger
		GrpcServer           *grpc.Server
		KVStoreServiceServer pb.KVStoreServiceServer
		Usecase              interactor.Interactor
		Boundary             boundary.Boundary
		Repository           repository.Repository
		Syncmap              syncmap.Syncmap
	}
	serviceName string
	version     string
	store       sync.Map
}

func (i *injectionImpl) InjectConfig() *config.Config {
	actual, _ := i.store.LoadOrStore("config", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			i.container.Config = config.NewConfig()
		})
	}
	return i.container.Config
}

func (i *injectionImpl) InjectLogger() logger.Logger {
	actual, _ := i.store.LoadOrStore("logger", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			l, err := logger.NewLogger(
				i.InjectConfig().LogLevel,
				i.serviceName,
				i.version,
				i.InjectConfig().Environment.String(),
			)
			if err != nil {
				log.Fatal(err)
			}
			i.container.Logger = l
		})
	}
	return i.container.Logger
}

func (i *injectionImpl) InjectGrpcServer() *grpc.Server {
	actual, _ := i.store.LoadOrStore("grpcServer", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			i.container.GrpcServer = grpcserver.NewGrpcServer()
		})
	}
	return i.container.GrpcServer
}

func (i *injectionImpl) InjectController() pb.KVStoreServiceServer {
	actual, _ := i.store.LoadOrStore("controller", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			i.container.KVStoreServiceServer = controller.NewKVStoreServiceServer(
				i.InjectConfig(),
				i.InjectUsecase(),
				i.InjectBoundary(),
			)
		})
	}
	return i.container.KVStoreServiceServer
}

func (i *injectionImpl) InjectUsecase() interactor.Interactor {
	actual, _ := i.store.LoadOrStore("interactor", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			i.container.Usecase = interactor.NewInteractor(i.InjectConfig(), i.InjectRepository())
		})
	}
	return i.container.Usecase
}

func (i *injectionImpl) InjectBoundary() boundary.Boundary {
	actual, _ := i.store.LoadOrStore("boundary", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			i.container.Boundary = boundary.NewBoundary()
		})
	}
	return i.container.Boundary
}

func (i *injectionImpl) InjectRepository() repository.Repository {
	actual, _ := i.store.LoadOrStore("repository", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			i.container.Repository = repository.NewRepository(i.InjectConfig(), i.InjectSyncmap())
		})
	}
	return i.container.Repository
}

func (i *injectionImpl) InjectSyncmap() syncmap.Syncmap {
	actual, _ := i.store.LoadOrStore("syncmap", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			i.container.Syncmap = syncmap.NewSyncMap()
		})
	}
	return i.container.Syncmap
}
