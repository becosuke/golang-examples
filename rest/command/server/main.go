package main

import (
	"github.com/becosuke/golang-examples/rest/presentation/handler"
	"github.com/becosuke/golang-examples/rest/registry/injector"
)

func main() {
	in := injector.NewInjector()
	handler.NewHandler(in.InjectConfig(), in.InjectController()).Serve()
}
