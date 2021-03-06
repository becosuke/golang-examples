package b58uuid

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/google/uuid"
	"github.com/mr-tron/base58"
)

type Origin struct {
	v uuid.UUID
}

func (o *Origin) String() string {
	return o.v.String()
}

func (o *Origin) Bytes() []byte {
	bin, _ := o.v.MarshalBinary()
	return bin
}

func (o *Origin) Hex() string {
	return hex.EncodeToString(o.Bytes())
}

func (o *Origin) Base58() string {
	return base58.Encode(o.Bytes())
}

func (o *Origin) Sha512() []byte {
	s := sha512.Sum512(o.Bytes())
	return s[:]
}

func (o *Origin) Sha512Hex() string {
	return hex.EncodeToString(o.Sha512())
}

func (o *Origin) Sha512Base58() string {
	return base58.Encode(o.Sha512())
}
