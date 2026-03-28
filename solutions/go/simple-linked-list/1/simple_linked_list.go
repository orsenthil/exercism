package linkedlist

import "errors"

// Define the List and Element types here.

type Element struct {
	value int
	next  *Element
}

type List struct {
	head *Element
	size int
}

func New(elements []int) *List {
	list := &List{}
	for _, element := range elements {
		list.Push(element)
	}
	return list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	e := &Element{value: element}
	if l.head == nil {
		l.head = e
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = e
	}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("empty list")
	}
	if l.head.next == nil {
		value := l.head.value
		l.head = nil
		l.size--
		return value, nil
	}

	current := l.head
	for current.next.next != nil {
		current = current.next
	}

	value := current.next.value
	current.next = nil
	l.size--
	return value, nil
}

func (l *List) Array() []int {
	result := make([]int, 0, l.size)
	current := l.head
	for current != nil {
		result = append(result, current.value)
		current = current.next
	}
	return result
}

func (l *List) Reverse() *List {
	if l.head == nil {
		return New(nil)
	}
	var prev *Element
	current := l.head

	for current != nil {
		nextTemp := current.next
		current.next = prev
		prev = current
		current = nextTemp
	}

	return &List{head: prev, size: l.size}
}
