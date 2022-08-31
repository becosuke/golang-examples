package pack

import (
	"context"
)

type Repository interface {
	Read(context.Context, *Key) (*Pack, error)
	Create(context.Context, *Value) (*Key, error)
	Update(context.Context, *Pack) error
	Delete(context.Context, *Key) error
}
