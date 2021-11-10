package entities

type Node struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type NodeKey struct {
	Key string `json:"key"`
}

func NewNode(key, value string) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

func NewNodeKey(key string) *NodeKey {
	return &NodeKey{Key: key}
}

func (s *Node) GetKey() string {
	return s.Key
}

func (s *Node) GetValue() string {
	return s.Value
}
