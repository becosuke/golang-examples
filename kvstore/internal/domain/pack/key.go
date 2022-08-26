package pack

import (
	"github.com/google/uuid"
)

type Key uuid.UUID

func NewKey(key uuid.UUID) *Key {
	k := Key(key)
	return &k
}

func (k *Key) UUID() uuid.UUID {
	if k == nil {
		return uuid.Nil
	}
	return uuid.UUID(*k)
}

func (k *Key) String() string {
	return k.UUID().String()
}
