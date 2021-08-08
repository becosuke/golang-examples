package main

import (
	"net/http"
	"sync"
)

var once sync.Once
var mux *http.ServeMux

func main() {
	http.ListenAndServe(":8080", route())
}

func route() *http.ServeMux {
	once.Do(func() {
		mux = http.NewServeMux()
		mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				GetHello(w, r)
			default:
				http.NotFound(w, r)
			}
		})
	})
	return mux
}
