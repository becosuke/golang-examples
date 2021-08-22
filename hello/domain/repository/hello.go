package repository

import (
	"context"
	"github.com/becosuke/golang-examples/hello/domain/entity"
)

type HelloRepository interface {
	GetHello(ctx context.Context) *entity.Hello
}
