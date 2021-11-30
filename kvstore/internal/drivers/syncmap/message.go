package syncmap

type Message struct {
	Key   string
	Value string
}

func NewMessage(key, value string) *Message {
	return &Message{Key: key, Value: value}
}

func (m *Message) GetKey() string {
	return m.Key
}

func (m *Message) GetValue() string {
	return m.Value
}

type MessageKey struct {
	Key string
}

func NewMessageKey(key string) *MessageKey {
	return &MessageKey{Key: key}
}

func (m *MessageKey) GetKey() string {
	return m.Key
}
