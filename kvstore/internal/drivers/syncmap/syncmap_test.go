package syncmap

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

func TestNewSyncmap(t *testing.T) {
	syncmap := NewSyncmap()

	assert.Implements(t, (*Syncmap)(nil), syncmap)
}

func TestSyncmapImpl_LoadOrStore(t *testing.T) {
	syncmap := NewSyncmap()

	actual, loaded, err := syncmap.LoadOrStore(nil)
	assert.Nil(t, actual)
	assert.False(t, loaded)
	assert.EqualError(t, err, ErrSyncmapInvalidArgument.String())

	message1 := NewMessage("kkk", "vvv")
	actual, loaded, err = syncmap.LoadOrStore(message1)
	assert.Equal(t, message1.Key(), actual.Key())
	assert.Equal(t, message1.Value(), actual.Value())
	assert.False(t, loaded)
	assert.NoError(t, err)

	message2 := NewMessage("kkk", "bbb")
	actual, loaded, err = syncmap.LoadOrStore(message2)
	assert.Equal(t, message2.Key(), actual.Key())
	assert.Equal(t, message1.Value(), actual.Value()) // no change
	assert.True(t, loaded)
	assert.NoError(t, err)

	syncMap := &sync.Map{}
	syncMap.Store("kkk", 111)
	syncmap = &syncmapImpl{syncmap: syncMap}
	actual, loaded, err = syncmap.LoadOrStore(NewMessage("kkk", "vvv"))
	assert.Nil(t, actual)
	assert.True(t, loaded)
	assert.EqualError(t, err, ErrSyncmapInvalidData.String())
}

func TestSyncmapImpl_Load(t *testing.T) {
	syncmap := NewSyncmap()

	actual, err := syncmap.Load("")
	assert.Nil(t, actual)
	assert.EqualError(t, err, ErrSyncmapInvalidArgument.String())

	messageKey := "kkk"
	actual, err = syncmap.Load(messageKey)
	assert.Nil(t, actual)
	assert.EqualError(t, err, ErrSyncmapNotFound.String())

	message := NewMessage("kkk", "vvv")
	_, err = syncmap.Store(message)
	require.NoError(t, err)

	actual, err = syncmap.Load(messageKey)
	assert.Equal(t, message.Key(), actual.Key())
	assert.Equal(t, message.Value(), actual.Value())
	assert.NoError(t, err)

	syncMap := &sync.Map{}
	syncMap.Store("kkk", 111)
	syncmap = &syncmapImpl{syncmap: syncMap}
	actual, err = syncmap.Load("kkk")
	assert.Nil(t, actual)
	assert.EqualError(t, err, ErrSyncmapInvalidData.String())
}

func TestSyncmapImpl_Store(t *testing.T) {
	syncmap := NewSyncmap()

	actual, err := syncmap.Store(nil)
	assert.Nil(t, actual)
	assert.EqualError(t, err, ErrSyncmapInvalidArgument.String())

	message1 := NewMessage("kkk", "vvv")
	actual, err = syncmap.Store(message1)
	assert.Equal(t, message1.Key(), actual.Key())
	assert.Equal(t, message1.Value(), actual.Value())
	assert.NoError(t, err)

	message2 := NewMessage("kkk", "bbb")
	actual, err = syncmap.Store(message2)
	assert.Equal(t, message2.Key(), actual.Key())
	assert.Equal(t, message2.Value(), actual.Value())
	assert.NoError(t, err)
}

func TestSyncmapImpl_Delete(t *testing.T) {
	syncmap := NewSyncmap()

	err := syncmap.Delete("")
	assert.EqualError(t, err, ErrSyncmapInvalidArgument.String())

	message := NewMessage("kkk", "vvv")
	_, err = syncmap.Store(message)
	require.NoError(t, err)

	messageKey := "kkk"
	err = syncmap.Delete(messageKey)
	assert.NoError(t, err)

	messageKey = "kkk"
	err = syncmap.Delete(messageKey)
	assert.NoError(t, err)
}
