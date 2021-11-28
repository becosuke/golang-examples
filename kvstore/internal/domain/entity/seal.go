package entity

type Seal struct {
	Key string `json:"key"`
}

func NewSeal(key string) *Seal {
	return &Seal{Key: key}
}

func (p *Seal) GetKey() string {
	return p.Key
}
