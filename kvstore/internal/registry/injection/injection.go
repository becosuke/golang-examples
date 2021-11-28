package injection

import (
	"github.com/becosuke/golang-examples/kvstore/internal/adapters/boundary"
	"github.com/becosuke/golang-examples/kvstore/internal/adapters/controller"
	"github.com/becosuke/golang-examples/kvstore/internal/adapters/repository"
	"github.com/becosuke/golang-examples/kvstore/internal/drivers/syncmap"
	"github.com/becosuke/golang-examples/kvstore/internal/pb"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
	"github.com/becosuke/golang-examples/kvstore/internal/usecase"
	"google.golang.org/grpc"
	"sync"
)

type Injection interface {
	InjectConfig() *config.Config
	InjectGrpcServer() *grpc.Server
	InjectKVStoreServiceServer() pb.KVStoreServiceServer
	InjectUsecase() usecase.Interactor
	InjectBoundary() boundary.Boundary
	InjectRepository() repository.Repository
	InjectSyncmap() syncmap.Syncmap
}

func NewInjection() Injection {
	return &injectionImpl{}
}

type injectionImpl struct {
	container struct {
		Config               *config.Config
		GrpcServer           *grpc.Server
		KVStoreServiceServer pb.KVStoreServiceServer
		Usecase              usecase.Interactor
		Boundary             boundary.Boundary
		Repository           repository.Repository
		Syncmap              syncmap.Syncmap
	}
	store sync.Map
}

func (impl *injectionImpl) InjectGrpcServer() *grpc.Server {
	actual, _ := impl.store.LoadOrStore("grpcserver", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.GrpcServer = controller.NewGrpcServer()
		})
	}
	return impl.container.GrpcServer
}

func (impl *injectionImpl) InjectKVStoreServiceServer() pb.KVStoreServiceServer {
	actual, _ := impl.store.LoadOrStore("kvstoreserviceserver", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.KVStoreServiceServer = controller.NewKVStoreServiceServer(
				impl.InjectConfig(),
				impl.InjectUsecase(),
				impl.InjectBoundary(),
			)
		})
	}
	return impl.container.KVStoreServiceServer
}

func (impl *injectionImpl) InjectConfig() *config.Config {
	actual, _ := impl.store.LoadOrStore("config", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Config = config.NewConfig()
		})
	}
	return impl.container.Config
}

func (impl *injectionImpl) InjectUsecase() usecase.Interactor {
	actual, _ := impl.store.LoadOrStore("interactor", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Usecase = usecase.NewInteractor(impl.InjectConfig(), impl.InjectRepository())
		})
	}
	return impl.container.Usecase
}

func (impl *injectionImpl) InjectBoundary() boundary.Boundary {
	actual, _ := impl.store.LoadOrStore("boundary", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Boundary = boundary.NewBoundary()
		})
	}
	return impl.container.Boundary
}

func (impl *injectionImpl) InjectRepository() repository.Repository {
	actual, _ := impl.store.LoadOrStore("repository", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Repository = repository.NewRepository(impl.InjectConfig(), impl.InjectSyncmap())
		})
	}
	return impl.container.Repository
}

func (impl *injectionImpl) InjectSyncmap() syncmap.Syncmap {
	actual, _ := impl.store.LoadOrStore("syncmap", &sync.Once{})
	once, ok := actual.(*sync.Once)
	if ok {
		once.Do(func() {
			impl.container.Syncmap = syncmap.NewSyncMap()
		})
	}
	return impl.container.Syncmap
}
