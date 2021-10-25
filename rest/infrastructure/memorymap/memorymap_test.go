package memorymap

import (
	"github.com/becosuke/golang-examples/rest/domain/entity"
	"github.com/becosuke/golang-examples/rest/mock/syncmap"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	key   = "key"
	value = "value"
)

func TestNewMemoryMap(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := syncmap.NewMockSyncMap(ctrl)
	impl := NewMemoryMap(mock)

	assert.Implements(t, (*MemoryMap)(nil), impl)
}

func TestMemoryMapImpl_LoadOrStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := syncmap.NewMockSyncMap(ctrl)
	impl := memoryMapImpl{syncMap: mock}

	mock.EXPECT().LoadOrStore(gomock.Eq(key), gomock.Eq(value)).Return(value, true)
	message, loaded, err := impl.LoadOrStore(key, value)
	assert.Equal(t, &Message{Key: key, Value: value}, message)
	assert.True(t, loaded)
	assert.Nil(t, err)

	mock.EXPECT().LoadOrStore(gomock.Eq(key), gomock.Eq(value)).Return(value, false)
	message, loaded, err = impl.LoadOrStore(key, value)
	assert.Equal(t, &Message{Key: key, Value: value}, message)
	assert.False(t, loaded)
	assert.Nil(t, err)

	mock.EXPECT().LoadOrStore(gomock.Eq(key), gomock.Eq(value)).Return(struct{}{}, true)
	message, loaded, err = impl.LoadOrStore(key, value)
	assert.Equal(t, &Message{}, message)
	assert.True(t, loaded)
	assert.NotNil(t, err)
}

func TestMemoryMapImpl_Load(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := syncmap.NewMockSyncMap(ctrl)
	impl := memoryMapImpl{syncMap: mock}

	mock.EXPECT().Load(gomock.Eq(key)).Return(value, true)
	message, err := impl.Load(key)
	assert.Equal(t, &Message{Key: key, Value: value}, message)
	assert.Nil(t, err)

	mock.EXPECT().Load(gomock.Eq(key)).Return(value, false)
	message, err = impl.Load(key)
	assert.Equal(t, &Message{}, message)
	assert.NotNil(t, err)

	mock.EXPECT().Load(gomock.Eq(key)).Return(struct{}{}, true)
	message, err = impl.Load(key)
	assert.Equal(t, &Message{}, message)
	assert.NotNil(t, err)
}

func TestMemoryMapImpl_Store(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := syncmap.NewMockSyncMap(ctrl)
	impl := memoryMapImpl{syncMap: mock}

	mock.EXPECT().Store(gomock.Eq(key), gomock.Eq(value)).Return()
	message, err := impl.Store(key, value)
	assert.Equal(t, &Message{Key: key, Value: value}, message)
	assert.Nil(t, err)
}

func TestMemoryMapImpl_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := syncmap.NewMockSyncMap(ctrl)
	impl := memoryMapImpl{syncMap: mock}

	mock.EXPECT().Delete(gomock.Eq(key)).Return()
	err := impl.Delete(key)
	assert.Nil(t, err)
}

func TestNewMessage(t *testing.T) {
	message := NewMessage(key, value)
	assert.Equal(t, &Message{Key: key, Value: value}, message)
}

func TestMessage_ConvertToEntity(t *testing.T) {
	message := NewMessage(key, value)
	assert.Equal(t, &entity.Entity{Key: key, Value: value}, message.ConvertToEntity())
}
