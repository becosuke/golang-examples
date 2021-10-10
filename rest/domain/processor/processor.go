package processor

import (
	"github.com/becosuke/golang-examples/rest/registry/config"
)

type Processor interface {
	Const(string) (string, bool)
}

func NewProcessor(conf *config.Config) Processor {
	return &processorImpl{config: conf}
}

type processorImpl struct {
	config *config.Config
}

func (impl *processorImpl) Const(key string) (string, bool) {
	val, ok := impl.config.Constant[key]
	return val, ok
}
