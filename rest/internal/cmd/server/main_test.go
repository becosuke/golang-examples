package main

import (
	"github.com/becosuke/golang-examples/rest/internal/registry/injector"
	"testing"
)

func Test_main(t *testing.T) {
	in := injector.NewInjector()
	in.InjectServer().Serve()
}
