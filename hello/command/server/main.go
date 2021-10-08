package main

import (
	"net/http"
	"sync"

	"github.com/becosuke/golang-examples/hello/application/controller"
	"github.com/becosuke/golang-examples/hello/config/injection"
)

var once sync.Once
var mux *http.ServeMux
var container *injection.Container
var helloController controller.Hello

func init() {
	container = &injection.Container{}
	helloController = container.InjectHelloController()
}

func main() {
	http.ListenAndServe(":8080", route())
}

func route() *http.ServeMux {
	once.Do(func() {
		mux = http.NewServeMux()
		mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				helloController.Get(w, r)
			default:
				http.NotFound(w, r)
			}
		})
	})
	return mux
}
