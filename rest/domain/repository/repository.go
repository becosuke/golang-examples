package repository

import (
	"context"

	"github.com/becosuke/golang-examples/rest/domain/entity"
)

type Repository interface {
	Create(context.Context, string, string) (*entity.Entity, error)
	Read(context.Context, string) (*entity.Entity, error)
	Update(context.Context, string, string) (*entity.Entity, error)
	Delete(context.Context, string) error
}
