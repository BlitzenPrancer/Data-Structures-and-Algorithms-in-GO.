package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	next  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	n := l.First()
	for {
		fmt.Print(n.value, "->")
		n = n.Next()
		if n == nil {
			fmt.Print(nil)
			break
		}
	}
}
