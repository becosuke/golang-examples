package pack

import (
	"context"
)

type Repository interface {
	Read(context.Context, *Key) (*Pack, error)
	Create(context.Context, *Pack) error
	Update(context.Context, *Pack) error
	Delete(context.Context, *Key) error
}