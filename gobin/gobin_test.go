package gobin

import (
	"testing"
)

var h = &Hoge{C1: "asdf", C2: 3}

func TestEncv(t *testing.T) {
	t.Log(string(Encv(*h)))
}

func TestEncp(t *testing.T) {
	t.Log(Encp(h))
}
