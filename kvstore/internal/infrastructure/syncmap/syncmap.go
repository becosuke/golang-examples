package syncmap

import (
	"errors"
	"sync"
)

type SyncMap interface {
	LoadOrStore(string, string) (*Message, bool, error)
	Load(string) (*Message, error)
	Store(string, string) (*Message, error)
	Delete(string) error
}

func NewSyncMap() SyncMap {
	return &syncMapImpl{}
}

type syncMapImpl struct {
	syncMap sync.Map
}

func (impl *syncMapImpl) LoadOrStore(key, value string) (*Message, bool, error) {
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

func (impl *syncMapImpl) Load(key string) (*Message, error) {
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

func (impl *syncMapImpl) Store(key, value string) (*Message, error) {
	impl.syncMap.Store(key, value)
	return &Message{Key: key, Value: value}, nil
}

func (impl *syncMapImpl) Delete(key string) error {
	impl.syncMap.Delete(key)
	return nil
}
