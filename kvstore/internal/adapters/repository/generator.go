package repository

import (
	"github.com/becosuke/golang-examples/kvstore/internal/domain/pack"
	"github.com/google/uuid"
)

func NewGenerator() pack.Generator {
	return &generatorImpl{}
}

type generatorImpl struct{}

func (impl *generatorImpl) GenerateKey() *pack.Key {
	k := pack.Key(uuid.NewString())
	return &k
}
