package syncmap

type Key string
type Value string

type Message struct {
	Key
	Value
}

func NewMessage(key Key, value Value) *Message {
	return &Message{
		Key:   key,
		Value: value,
	}
}

func (m Message) GetKey() Key {
	return m.Key
}

func (m Message) GetValue() Value {
	return m.Value
}

func (k Key) String() string {
	return string(k)
}

func (v Value) String() string {
	return string(v)
}
