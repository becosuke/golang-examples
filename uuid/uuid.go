package uuid

import (
	"github.com/google/uuid"
	"github.com/lytics/base62"
	"github.com/mr-tron/base58"

)

type Origin struct {
	v uuid.UUID
}

func NewOrigin() Origin {
	return Origin{
		v: uuid.New(),
	}
}

func (o *Origin) String() string {
	return o.v.String()
}

func (o *Origin) Base62() string {
	return base62.StdEncoding.EncodeToString([]byte(o.v.String()))
}

func (o *Origin) Base58() string {
	b, _ := o.v.MarshalBinary()
	return base58.Encode(b)
}
