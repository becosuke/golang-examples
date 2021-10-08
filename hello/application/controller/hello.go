package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/becosuke/golang-examples/hello/config"
	"github.com/becosuke/golang-examples/hello/domain/service"
)

func NewHello(cfg *config.Config, service service.Hello) Hello {
	return &helloImpl{
		config:  cfg,
		service: service,
	}
}

type Hello interface {
	Get(http.ResponseWriter, *http.Request)
}

type helloImpl struct {
	config  *config.Config
	service service.Hello
}

func (impl *helloImpl) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res := impl.service.Get(ctx)
	j, _ := json.Marshal(res)
	fmt.Fprint(w, string(j))
}
