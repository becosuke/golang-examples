package singly

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	first *Node
}

func (l *List) IsEmpty() bool {
	return l == nil || l.first == nil
}

func (l *List) Head() *Node {
	if l.IsEmpty() {
		return nil
	}

	return l.first
}

func (l *List) Tail() *List {
	if l.IsEmpty() {
		return nil
	}

	return &List{first: l.first.Next}
}

func (l *List) Prepend(data interface{}) *List {
	return &List{first: &Node{Data: data, Next: l.first}}
}
