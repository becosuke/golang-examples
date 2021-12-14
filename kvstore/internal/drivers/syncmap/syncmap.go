package syncmap

import (
	"github.com/becosuke/golang-examples/kvstore/internal/registry/errors"
	"sync"
)

type Syncmap interface {
	LoadOrStore(*Message) (*Message, bool, error)
	Load(Key) (*Message, error)
	Store(*Message) (*Message, error)
	Delete(Key) error
}

func NewSyncmap() Syncmap {
	return &syncmapImpl{
		syncmap: &sync.Map{},
	}
}

type syncmapImpl struct {
	syncmap *sync.Map
}

func (si *syncmapImpl) LoadOrStore(message *Message) (*Message, bool, error) {
	if message == nil {
		return nil, false, errors.ErrSyncmapInvalidArgument
	}
	actual, loaded := si.syncmap.LoadOrStore(message.GetKey(), message.GetValue())
	if loaded {
		asserted, ok := actual.(Value)
		if !ok {
			return nil, loaded, errors.ErrSyncmapInvalidData
		}
		return &Message{Key: message.GetKey(), Value: asserted}, loaded, nil
	}
	return &Message{Key: message.GetKey(), Value: message.GetValue()}, loaded, nil
}

func (si *syncmapImpl) Load(key Key) (*Message, error) {
	if key == "" {
		return nil, errors.ErrSyncmapInvalidArgument
	}
	value, ok := si.syncmap.Load(key)
	if !ok {
		return nil, errors.ErrSyncmapNotFound
	}
	asserted, ok := value.(Value)
	if !ok {
		return nil, errors.ErrSyncmapInvalidData
	}
	return &Message{Key: key, Value: asserted}, nil
}

func (si *syncmapImpl) Store(message *Message) (*Message, error) {
	if message == nil {
		return nil, errors.ErrSyncmapInvalidArgument
	}
	si.syncmap.Store(message.GetKey(), message.GetValue())
	return &Message{Key: message.GetKey(), Value: message.GetValue()}, nil
}

func (si *syncmapImpl) Delete(key Key) error {
	if key == "" {
		return errors.ErrSyncmapInvalidArgument
	}
	si.syncmap.Delete(key)
	return nil
}
