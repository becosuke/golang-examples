package uuid

import (
	"github.com/google/uuid"
	"github.com/mr-tron/base58"
	"log"
	"os"
	"testing"
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
	return nil
}

func teardown() {
}

func TestNewOrigin(t *testing.T) {
	origin = NewOrigin()
	t.Log(origin)
}

func TestOrigin_String(t *testing.T) {
	t.Log(origin.String())
}

func TestOrigin_Base58(t *testing.T) {
	ori := origin.Base58()
	t.Log(ori)
	dec, _ := base58.Decode(string(ori))
	uu, _ := uuid.FromBytes(dec)
	t.Log(uu.String())

	src := uuid.MustParse("00000000-9058-4498-8002-2b6500000000")
	bin2, _ := src.MarshalBinary()
	ori2 := base58.Encode(bin2)
	t.Log(ori2)
	dec2, _ := base58.Decode(string(ori2))
	uu2, _ := uuid.FromBytes(dec2)
	t.Log(uu2.String())
}

func TestOrigin_Base62(t *testing.T) {
	t.Log(origin.Base62())
}