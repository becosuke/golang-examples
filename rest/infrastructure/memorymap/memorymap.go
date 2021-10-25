package memorymap

import (
	"errors"
	"github.com/becosuke/golang-examples/rest/domain/entity"
	"github.com/becosuke/golang-examples/rest/infrastructure/syncmap"
)

type MemoryMap interface {
	LoadOrStore(string, string) (*Message, bool, error)
	Load(string) (*Message, error)
	Store(string, string) (*Message, error)
	Delete(string) error
}

func NewMemoryMap(syncMap syncmap.SyncMap) MemoryMap {
	return &memoryMapImpl{
		syncMap: syncMap,
	}
}

type memoryMapImpl struct {
	syncMap syncmap.SyncMap
}

func (impl *memoryMapImpl) LoadOrStore(key, value string) (*Message, bool, error) {
	actual, loaded := impl.syncMap.LoadOrStore(key, value)
	if loaded {
		asserted, ok := actual.(string)
		if !ok {
			return &Message{}, loaded, errors.New("invalid data was returned")
		}
		return &Message{Key: key, Value: asserted}, loaded, nil
	}
	return &Message{Key: key, Value: value}, loaded, nil
}

func (impl *memoryMapImpl) Load(key string) (*Message, error) {
	value, ok := impl.syncMap.Load(key)
	if !ok {
		return &Message{}, errors.New("not exists")
	}
	asserted, ok := value.(string)
	if !ok {
		return &Message{}, errors.New("invalid data was returned")
	}
	return &Message{Key: key, Value: asserted}, nil
}

func (impl *memoryMapImpl) Store(key, value string) (*Message, error) {
	impl.syncMap.Store(key, value)
	return &Message{Key: key, Value: value}, nil
}

func (impl *memoryMapImpl) Delete(key string) error {
	impl.syncMap.Delete(key)
	return nil
}

type Message struct {
	Key   string
	Value string
}

func NewMessage(key, value string) *Message {
	return &Message{Key: key, Value: value}
}

func (m *Message) ConvertToEntity() *entity.Entity {
	return &entity.Entity{
		Key:   m.Key,
		Value: m.Value,
	}
}
