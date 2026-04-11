package datastructures

import "errors"

// Node is a node in the singly linked list.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// SinglyLinkedList is a generic singly linked list.
type SinglyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// AddFirst inserts an item at the front.
func (l *SinglyLinkedList[T]) AddFirst(value T) {
	newNode := &Node[T]{Value: value}
	l.head = newNode
	if l.tail == nil { // was empty
		l.tail = newNode
	}
	l.size++
}

// AddLast inserts an item at the end.
func (l *SinglyLinkedList[T]) AddLast(value T) {
	newNode := &Node[T]{Value: value}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}
	l.size++
}

// RemoveFirst removes the item at the front and returns its value.
func (l *SinglyLinkedList[T]) RemoveFirst() (T, error) {
	var zero T
	if l.head == nil {
		return zero, errors.New("list is empty")
	}
	val := l.head.Value
	l.head = l.head.Next
	if l.head == nil { // list became empty
		l.tail = nil
	}
	l.size--
	return val, nil
}

// First returns the item at the front.
func (l *SinglyLinkedList[T]) First() (T, error) {
	var zero T
	if l.head == nil {
		return zero, errors.New("list is empty")
	}
	return l.head.Value, nil
}

// Last returns the item at the end.
func (l *SinglyLinkedList[T]) Last() (T, error) {
	var zero T
	if l.tail == nil {
		return zero, errors.New("list is empty")
	}
	return l.tail.Value, nil
}

// Len returns the number of items in the list.
func (l *SinglyLinkedList[T]) Len() int {
	return l.size
}

// IsEmpty checks if the list is empty.
func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Clear removes all elements from the list.
func (l *SinglyLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Iterate calls the given function for each element in order.
func (l *SinglyLinkedList[T]) Iterate(f func(T)) {
	for cur := l.head; cur != nil; cur = cur.Next {
		f(cur.Value)
	}
}