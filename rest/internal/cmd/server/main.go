package main

import (
	"github.com/becosuke/golang-examples/rest/internal/registry/injector"
)

func main() {
	in := injector.NewInjector()
	in.InjectServer().Serve()
}
