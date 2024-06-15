package main

import (
	"errors"
	"fmt"
)

type node struct {
	val  int
	next *node
	prev *node
}

type LinkedList struct {
	head *node
	tail *node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) InsertFront(number int) {
	newNode := node{val: number}
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
	} else {
		newNode.next = l.head
		l.head.prev = &newNode
		l.head = &newNode
	}
	l.size++
}

func (l *LinkedList) InsertBack(number int) {
	newNode := node{val: number}
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
	} else {
		l.tail.next = &newNode
		newNode.prev = l.tail
		l.tail = &newNode
	}
	l.size++
}

func (l *LinkedList) InsertAt(idx, number int) error {
	if idx < 0 || idx > l.size {
		return errors.New("index out of bounds")
	}

	if idx == 0 {
		l.InsertFront(number)
		return nil
	}

	if idx == l.size {
		l.InsertBack(number)
		return nil
	}

	newNode := node{val: number}
	current := l.head
	for i := 0; i < idx; i++ {
		current = current.next
	}

	newNode.prev = current.prev
	newNode.next = current
	current.prev.next = &newNode
	current.prev = &newNode

	l.size++
	return nil
}

func (ll *LinkedList) RemoveFront() (*int, error) {
	if ll.head == nil {
		return nil, errors.New("list is empty")
	}
	value := ll.head.val
	ll.head = ll.head.next
	if ll.head != nil {
		ll.head.prev = nil
	} else {
		ll.tail = nil
	}
	ll.size--
	return &value, nil
}

func (ll *LinkedList) RemoveBack() (*int, error) {
	if ll.head == nil {
		return nil, errors.New("list is empty")
	}
	value := ll.tail.val
	ll.tail = ll.tail.prev
	if ll.tail != nil {
		ll.tail.next = nil
	} else {
		ll.head = nil
	}
	ll.size--
	return &value, nil
}

func (ll *LinkedList) RemoveAt(index int) (*int, error) {
	if index < 0 || index >= ll.size {
		return nil, errors.New("index out of bounds")
	}
	if index == 0 {
		return ll.RemoveFront()
	}
	if index == ll.size-1 {
		return ll.RemoveBack()
	}

	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	value := current.val
	current.prev.next = current.next
	current.next.prev = current.prev
	ll.size--
	return &value, nil
}

func (ll *LinkedList) Get(index int) (*int, error) {
	if index < 0 || index >= ll.size {
		return nil, errors.New("index out of bounds")
	}
	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return &current.val, nil
}

func main() {
	ll := NewLinkedList()
	ll.InsertFront(1)
	ll.InsertBack(2)
	ll.InsertAt(1, 3)
	ele, _ := ll.Get(0)
	fmt.Println(*ele) // Output: 1
	ele, _ = ll.Get(1)
	fmt.Println(*ele) // Output: 3
	ele, _ = ll.Get(2)
	fmt.Println(*ele) // Output: 2

	// ll.Set(1, 4)
	// fmt.Println(ll.Get(1)) // Output: 4

	ll.RemoveFront()
	ll.RemoveBack()
	ll.RemoveAt(0)
	fmt.Println(ll.Size()) // Output: 0
}
