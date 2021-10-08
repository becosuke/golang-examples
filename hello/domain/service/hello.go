package service

import (
	"context"
	"github.com/becosuke/golang-examples/hello/domain/entity"
)

type Hello interface {
	Get(ctx context.Context) *entity.Hello
}
