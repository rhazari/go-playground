package datastructures

import "sync/atomic"

// SPSCQueue is a bounded lock-free ring buffer for exactly one producer
// goroutine and one consumer goroutine.
type SPSCQueue[T any] struct {
	_        [56]byte
	head     atomic.Uint64
	_        [56]byte
	tail     atomic.Uint64
	_        [56]byte
	buffer   []T
	capacity uint64
}

// NewSPSCQueue creates a bounded single-producer/single-consumer queue.
func NewSPSCQueue[T any](capacity int) *SPSCQueue[T] {
	if capacity <= 0 {
		panic("spsc queue capacity must be greater than zero")
	}

	return &SPSCQueue[T]{
		buffer:   make([]T, capacity),
		capacity: uint64(capacity),
	}
}

// Enqueue adds an item to the queue.
// It returns false when the queue is full.
func (q *SPSCQueue[T]) Enqueue(value T) bool {
	tail := q.tail.Load()
	head := q.head.Load()
	if tail-head >= q.capacity {
		return false
	}

	q.buffer[tail%q.capacity] = value
	q.tail.Store(tail + 1)
	return true
}

// Dequeue removes and returns the next item from the queue.
// The bool result is false when the queue is empty.
func (q *SPSCQueue[T]) Dequeue() (T, bool) {
	head := q.head.Load()
	tail := q.tail.Load()
	if tail == head {
		var zero T
		return zero, false
	}

	index := head % q.capacity
	value := q.buffer[index]

	// Clear the slot after the consumer has taken ownership of the value.
	var zero T
	q.buffer[index] = zero
	q.head.Store(head + 1)
	return value, true
}

// Len returns the current number of items in the queue.
func (q *SPSCQueue[T]) Len() int {
	tail := q.tail.Load()
	head := q.head.Load()
	return int(tail - head)
}

// Cap returns the queue capacity.
func (q *SPSCQueue[T]) Cap() int {
	return int(q.capacity)
}

// IsEmpty reports whether the queue currently has no items.
func (q *SPSCQueue[T]) IsEmpty() bool {
	return q.Len() == 0
}

// IsFull reports whether the queue currently cannot accept more items.
func (q *SPSCQueue[T]) IsFull() bool {
	return q.Len() == int(q.capacity)
}
