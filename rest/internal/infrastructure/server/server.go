package server

import (
	"net/http"

	"github.com/becosuke/golang-examples/rest/internal/infrastructure/server/handler"
	"github.com/becosuke/golang-examples/rest/internal/registry/config"
)

type Server interface {
	Serve()
}

func NewServer(config *config.Config, handler handler.Handler) Server {
	return &serverImpl{
		config:  config,
		handler: handler,
	}
}

type serverImpl struct {
	config  *config.Config
	handler handler.Handler
}

func (impl *serverImpl) Serve() {
	_ = http.ListenAndServe(impl.config.HttpPort, impl.handler.Route())
}
