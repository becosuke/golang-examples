package b58uuid

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var origin Origin

func TestMain(m *testing.M) {
	log.SetFlags(log.Ldate | log.Ltime)
	if err := setup(); err != nil {
		log.Fatal(err)
	}
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func setup() error {
	origin = Origin{
		v: uuid.MustParse("85f72119-3127-4e2a-8f69-1d64d4062597"),
	}
	return nil
}

func teardown() {
}

func TestOrigin_String(t *testing.T) {
	assert.Equal(t, "85f72119-3127-4e2a-8f69-1d64d4062597", origin.String())
}

func TestOrigin_Bytes(t *testing.T) {
	assert.Equal(t, "85f7211931274e2a8f691d64d4062597", fmt.Sprintf("%x", origin.Bytes()))
}

func TestOrigin_Hex(t *testing.T) {
	assert.Equal(t, "85f7211931274e2a8f691d64d4062597", origin.Hex())
}

func TestOrigin_Base58(t *testing.T) {
	assert.Equal(t, "HYUMBz2JqfF2wnfbU72hYA", origin.Base58())
}

func TestOrigin_Sha512(t *testing.T) {
	assert.Equal(t, "913b13506dcd7bbfff09cb46202c24a8b0c7cc1824b1dec0250f5def573273e3c65a031a87930c0ec462fd629d991e5d020e570fb7a166f6aafccc2783672b26", fmt.Sprintf("%x", origin.Sha512()))
}

func TestOrigin_Sha512Hex(t *testing.T) {
	assert.Equal(t, "913b13506dcd7bbfff09cb46202c24a8b0c7cc1824b1dec0250f5def573273e3c65a031a87930c0ec462fd629d991e5d020e570fb7a166f6aafccc2783672b26", origin.Sha512Hex())
}

func TestOrigin_Sha512Base58(t *testing.T) {
	assert.Equal(t, "3uQoh7ZirdJhSWrE8Y3GxVyP4oLJgY6XPrviaurbsyEWmRLZHAKJjQ24HYyyY5mkzgBA81Kn1fAxjWjcchCGmaL5", origin.Sha512Base58())
}
