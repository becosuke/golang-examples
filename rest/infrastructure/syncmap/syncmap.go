package syncmap

import "sync"

type SyncMap interface {
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	Load(key interface{}) (value interface{}, ok bool)
	Store(key, value interface{})
	Delete(key interface{})
}

func NewSyncMap() SyncMap {
	return &sync.Map{}
}
