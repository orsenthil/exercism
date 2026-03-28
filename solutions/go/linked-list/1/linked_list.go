package linkedlist

import "errors"

type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

type List struct {
	head *Node
	tail *Node
}

func NewList(elements ...interface{}) *List {
	if len(elements) == 0 {
		return &List{}
	}
	list := &List{}
	list.head = &Node{Value: elements[0]}
	list.tail = list.head

	// Add the remaining elements
	for _, element := range elements[1:] {
		newNode := &Node{
			Value: element,
			prev:  list.tail,
		}
		list.tail.next = newNode
		list.tail = newNode
	}
	return list
}

func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
}

func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("list is empty")
	}
	value := l.head.Value
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}
	return value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("List is empty")
	}
	value := l.tail.Value
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}

	return value, nil
}

func (l *List) Reverse() {
	if l.head == nil || l.head == l.tail {
		return
	}
	current := l.head
	for current != nil {
		current.next, current.prev = current.prev, current.next
		current = current.prev
	}
	l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
