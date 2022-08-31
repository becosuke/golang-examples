package pack

type Key string

func NewKey(key string) *Key {
	k := Key(key)
	return &k
}

func (k *Key) String() string {
	if k == nil {
		return ""
	}
	return string(*k)
}
