package gobin

import (
	"bytes"
	"encoding/gob"
)

type Hoge struct {
	C1 string
	C2 int
}

func Encv(hoge Hoge) []byte {
	buf := bytes.NewBuffer(nil)
	_ = gob.NewEncoder(buf).Encode(&hoge)
	return buf.Bytes()
}

func Encp(hoge *Hoge) []byte {
	buf := bytes.NewBuffer(nil)
	_ = gob.NewEncoder(buf).Encode(&hoge)
	return buf.Bytes()
}

func Dec(data []byte) *Hoge {
	var hoge Hoge
	buf := bytes.NewBuffer(data)
	_ = gob.NewDecoder(buf).Decode(&hoge)
	return &hoge
}
