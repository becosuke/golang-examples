package node

type KeyValueNode struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type KeyNode struct {
	Key string `json:"key"`
}

func NewKeyValueNode(key, value string) *KeyValueNode {
	return &KeyValueNode{
		Key:   key,
		Value: value,
	}
}

func NewKeyNode(key string) *KeyNode {
	return &KeyNode{Key: key}
}

func (s *KeyValueNode) GetKey() string {
	return s.Key
}

func (s *KeyValueNode) GetValue() string {
	return s.Value
}
