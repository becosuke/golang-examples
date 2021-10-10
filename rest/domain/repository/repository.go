package repository

import (
	"context"

	"github.com/becosuke/golang-examples/rest/domain/model"
)

type Repository interface {
	Create(context.Context, string, string) (*model.Model, error)
	Read(context.Context, string) (*model.Model, error)
	Update(context.Context, string, string) (*model.Model, error)
	Delete(context.Context, string) error
}
