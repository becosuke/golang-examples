package handler

import (
	"sync"

	"net/http"

	"github.com/becosuke/golang-examples/kvstore/internal/adapter/controller"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
)

type Handler interface {
	Route() *http.ServeMux
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

func (impl *handlerImpl) Route() *http.ServeMux {
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
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		})
	})
	return impl.mux
}
