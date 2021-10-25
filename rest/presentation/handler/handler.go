package handler

import (
	"net/http"
	"sync"

	"github.com/becosuke/golang-examples/rest/presentation/controller"
	"github.com/becosuke/golang-examples/rest/registry/config"
)

type Handler interface {
	Serve()
}

func NewHandler(config *config.Config, controller controller.Controller) Handler {
	return &handlerImpl{
		config:     config,
		controller: controller,
	}
}

type handlerImpl struct {
	config     *config.Config
	controller controller.Controller
	once       sync.Once
	mux        *http.ServeMux
}

func (impl *handlerImpl) Serve() {
	_ = http.ListenAndServe(impl.config.HttpPort, impl.route())
}

func (impl *handlerImpl) route() *http.ServeMux {
	impl.once.Do(func() {
		impl.mux = http.NewServeMux()
		impl.mux.HandleFunc(config.Endpoint, func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				impl.controller.Post(w, r)
			case http.MethodGet:
				impl.controller.Get(w, r)
			case http.MethodPut:
				impl.controller.Put(w, r)
			case http.MethodDelete:
				impl.controller.Delete(w, r)
			default:
				http.NotFound(w, r)
			}
		})
	})
	return impl.mux
}
