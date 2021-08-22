package main

import (
	"github.com/becosuke/golang-examples/hello/config/injection"
	"github.com/becosuke/golang-examples/hello/presentation/controller"
	"net/http"
	"sync"
)

var once sync.Once
var mux *http.ServeMux
var container *injection.Container
var helloController controller.HelloController

func main() {
	_ = http.ListenAndServe(":8080", route())
}

func init() {
	container = &injection.Container{}
	helloController = container.InjectHelloController()
}

func route() *http.ServeMux {
	once.Do(func() {
		mux = http.NewServeMux()
		mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				helloController.GetHello(w, r)
			default:
				http.NotFound(w, r)
			}
		})
	})
	return mux
}
