package singly

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetFlags(log.Ldate | log.Ltime)
	if err := setup(); err != nil {
		log.Fatal(err)
	}
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func setup() error {
	return nil
}

func teardown() {
}

func TestNode(t *testing.T) {
	node1 := &Node{Data: 1}
	node2 := &Node{Data: 2, Next: node1}

	log.Printf("node1:%+v", node1)
	log.Printf("node2:%+v", node2)
}

func TestListPrepend(t *testing.T) {
	list := &List{}
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)

	node := list.first
	for node != nil {
		log.Printf("node:%+v", node)
		node = node.Next
	}
}

func TestListHead(t *testing.T) {
	list := &List{}
	list = list.Prepend(1)
	list = list.Prepend(2)
	list = list.Prepend(3)

	log.Printf("node1:%+v", list.Head())
}

func TestListTail(t *testing.T) {
	list1 := &List{}
	list1 = list1.Prepend(1)
	list1 = list1.Prepend(2)
	list1 = list1.Prepend(3)
	var list1node *Node = list1.first
	for list1node != nil {
		log.Printf("list1node:%+v", list1node)
		list1node = list1node.Next
	}

	list2 := list1.Tail()
	list2node := list2.first
	for list2node != nil {
		log.Printf("list2node:%+v", list2node)
		list2node = list2node.Next
	}
}
