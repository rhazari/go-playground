package datastructures

import "errors"

// Stack is a generic stack implementation using a slice.
type Stack[T any] struct {
	items []T
}

// Push puts an item on top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack.
// Returns an error if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.items) == 0 {
		return zero, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek returns (without removing) the item from the top of the stack.
func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if len(s.items) == 0 {
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// Size returns the number of items in the stack.
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// IsEmpty returns true if the stack has no items.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
