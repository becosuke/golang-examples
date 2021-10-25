package syncmap

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSyncMap(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	SUT := NewSyncMap()
	assert.Implements(t, (*SyncMap)(nil), SUT)
}
