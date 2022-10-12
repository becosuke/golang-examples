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
	newSyncmap := NewSyncmap()

	actual, loaded, err := newSyncmap.LoadOrStore(nil)
	assert.Nil(t, actual)
	assert.False(t, loaded)
	assert.EqualError(t, err, ErrSyncmapInvalidArgument.String())

	keyString1 := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString1 := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	message1 := NewMessage(keyString1, valueString1)
	actual, loaded, err = newSyncmap.LoadOrStore(message1)
	assert.Equal(t, keyString1, actual.Key())
	assert.Equal(t, valueString1, actual.Value())
	assert.False(t, loaded)
	assert.NoError(t, err)

	valueString2 :=  "d1ee96c5-e6de-42ff-a168-1272586c2c70"
	message2 := NewMessage(keyString1, valueString2)
	actual, loaded, err = newSyncmap.LoadOrStore(message2)
	assert.Equal(t, keyString1, actual.Key())
	assert.Equal(t, valueString1, actual.Value()) // no change
	assert.True(t, loaded)
	assert.NoError(t, err)

	syncmap := &sync.Map{}
	syncmap.Store(keyString1, 111)
	newSyncmap = &syncmapImpl{syncmap: syncmap}
	actual, loaded, err = newSyncmap.LoadOrStore(NewMessage(keyString1, valueString1))
	assert.Nil(t, actual)
	assert.True(t, loaded)
	assert.EqualError(t, err, ErrSyncmapInvalidData.String())
}

func TestSyncmapImpl_Load(t *testing.T) {
	newSyncmap := NewSyncmap()

	actual, err := newSyncmap.Load("")
	assert.Nil(t, actual)
	assert.EqualError(t, err, ErrSyncmapInvalidArgument.String())

	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	actual, err = newSyncmap.Load(keyString)
	assert.Nil(t, actual)
	assert.EqualError(t, err, ErrSyncmapNotFound.String())

	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	message := NewMessage(keyString, valueString)
	_, err = newSyncmap.Store(message)
	require.NoError(t, err)

	actual, err = newSyncmap.Load(keyString)
	assert.Equal(t, keyString, actual.Key())
	assert.Equal(t, valueString, actual.Value())
	assert.NoError(t, err)

	syncmap := &sync.Map{}
	syncmap.Store(keyString, 111)
	newSyncmap = &syncmapImpl{syncmap: syncmap}
	actual, err = newSyncmap.Load(keyString)
	assert.Nil(t, actual)
	assert.EqualError(t, err, ErrSyncmapInvalidData.String())
}

func TestSyncmapImpl_Store(t *testing.T) {
	newSyncmap := NewSyncmap()

	actual, err := newSyncmap.Store(nil)
	assert.Nil(t, actual)
	assert.EqualError(t, err, ErrSyncmapInvalidArgument.String())

	keyString1 := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString1 := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	message1 := NewMessage(keyString1, valueString1)
	actual, err = newSyncmap.Store(message1)
	assert.Equal(t, keyString1, actual.Key())
	assert.Equal(t, valueString1, actual.Value())
	assert.NoError(t, err)

	valueString2 :=  "d1ee96c5-e6de-42ff-a168-1272586c2c70"
	message2 := NewMessage(keyString1, valueString2)
	actual, err = newSyncmap.Store(message2)
	assert.Equal(t, keyString1, actual.Key())
	assert.Equal(t, valueString2, actual.Value())
	assert.NoError(t, err)
}

func TestSyncmapImpl_Delete(t *testing.T) {
	syncmap := NewSyncmap()

	err := syncmap.Delete("")
	assert.EqualError(t, err, ErrSyncmapInvalidArgument.String())

	keyString := "b3b8500d-3502-4420-a600-49081c68d24b"
	valueString := "25338aef-9462-4c0e-bc8d-e701d3f66cc3"
	message := NewMessage(keyString, valueString)
	_, err = syncmap.Store(message)
	require.NoError(t, err)

	err = syncmap.Delete(keyString)
	assert.NoError(t, err)

	err = syncmap.Delete(keyString)
	assert.NoError(t, err)
}
