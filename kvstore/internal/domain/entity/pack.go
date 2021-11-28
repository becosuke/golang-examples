package entity

type Pack struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewPack(key, value string) *Pack {
	return &Pack{
		Key:   key,
		Value: value,
	}
}

func (p *Pack) GetKey() string {
	return p.Key
}

func (p *Pack) GetValue() string {
	return p.Value
}
