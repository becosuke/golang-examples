package main

import (
	"github.com/becosuke/golang-examples/rest/registry/config"
	"github.com/becosuke/golang-examples/rest/registry/injector"
	"net/http"
	"sync"
)

type server struct {
	once sync.Once
	mux *http.ServeMux
	container *injector.Container
}

var s server


func init() {
	s = server{
		container: injector.NewContainer(),
	}
	_ = s.container.InjectConfig()

}

func main() {
	s.Serve()
}

func (s *server) Serve() {
	_ = http.ListenAndServe(s.container.InjectConfig().HttpPort, s.Route())
}

func (s *server) Route() *http.ServeMux {
	s.once.Do(func() {
		controller := s.container.InjectController()
		s.mux = http.NewServeMux()
		s.mux.HandleFunc(config.Endpoint, func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				controller.Post(w, r)
			case http.MethodGet:
				controller.Get(w, r)
			case http.MethodPut:
				controller.Put(w, r)
			case http.MethodDelete:
				controller.Delete(w, r)
			default:
				http.NotFound(w, r)
			}
		})
	})
	return s.mux
}
