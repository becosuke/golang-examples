package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/becosuke/golang-examples/hello/config"
	"github.com/becosuke/golang-examples/hello/domain/service"
)

func NewHelloController(cfg *config.Config, helloService service.HelloService) HelloController {
	return &helloControllerImpl{
		config:       cfg,
		helloService: helloService,
	}
}

type HelloController interface {
	GetHello(http.ResponseWriter, *http.Request)
}

type helloControllerImpl struct {
	config       *config.Config
	helloService service.HelloService
}

func (impl *helloControllerImpl) GetHello(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()
	res := impl.helloService.GetHello(ctx)
	j, _ := json.Marshal(res)
	_, _ = fmt.Fprint(w, string(j))
}
