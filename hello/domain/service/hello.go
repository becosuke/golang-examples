package service

import (
	"context"
	"github.com/becosuke/golang-examples/hello/domain/entity"
)

type HelloService interface {
	GetHello(ctx context.Context) *entity.Hello
}
