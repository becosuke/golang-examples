package pack

import (
	"context"
)

type Usecase interface {
	Read(context.Context, *Key) (*Pack, error)
	Create(context.Context, *Value) (*Pack, error)
	Update(context.Context, *Pack) (*Pack, error)
	Delete(context.Context, *Key) error
}
