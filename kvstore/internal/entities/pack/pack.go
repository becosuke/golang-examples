package pack

type Pack struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Seal struct {
	Key string `json:"key"`
}

func NewPack(key, value string) *Pack {
	return &Pack{
		Key:   key,
		Value: value,
	}
}

func NewSeal(key string) *Seal {
	return &Seal{Key: key}
}

func (p *Pack) GetKey() string {
	return p.Key
}

func (p *Pack) GetValue() string {
	return p.Value
}
