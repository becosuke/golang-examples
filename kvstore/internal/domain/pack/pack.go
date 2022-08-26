package pack

type Pack struct {
	key   *Key
	value *Value
}

func NewPack(key *Key, value *Value) *Pack {
	return &Pack{
		key:   key,
		value: value,
	}
}

func (p *Pack) Key() *Key {
	return p.key
}

func (p *Pack) Value() *Value {
	return p.value
}
