package syncmap

import (
	"errors"
	"sync"

	"github.com/becosuke/golang-examples/rest/domain/entity"
)

func NewSyncMap() SyncMap {
	return &syncMapImpl{
		syncMap: sync.Map{},
	}
}

type SyncMap interface {
	LoadOrStore(string, string) (*Message, error)
	Load(string) (*Message, error)
	Store(string, string) (*Message, error)
	Delete(string) error
}

type syncMapImpl struct {
	syncMap sync.Map
}

func (impl *syncMapImpl) LoadOrStore(key, value string) (*Message, error) {
	actual, loaded := impl.syncMap.LoadOrStore(key, value)
	message := &Message{Key: key, Value: actual.(string)}
	if loaded {
		return message, errors.New("already exists")
	}
	return message, nil
}

func (impl *syncMapImpl) Load(key string) (*Message, error) {
	value, ok := impl.syncMap.Load(key)
	if !ok {
		return &Message{}, errors.New("not exists")
	}
	return &Message{Key: key, Value: value.(string)}, nil
}

func (impl *syncMapImpl) Store(key, value string) (*Message, error) {
	impl.syncMap.Store(key, value)
	return &Message{Key: key, Value: value}, nil
}

func (impl *syncMapImpl) Delete(key string) error {
	impl.syncMap.Delete(key)
	return nil
}

type Message struct {
	Key   string
	Value string
}

func (m *Message) ConvertToEntity() *entity.Entity {
	return &entity.Entity{
		Key:   m.Key,
		Value: m.Value,
	}
}
