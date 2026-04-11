package datastructures

import (
	"fmt"
	"sync"
	"time"
)

// ConcurrentQueue is a thread-safe FIFO queue.
type ConcurrentQueue struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

// NewConcurrentQueue creates a new queue.
func NewConcurrentQueue() *ConcurrentQueue {
	q := &ConcurrentQueue{}
	q.cond = sync.NewCond(&q.mu)
	return q
}

// Enqueue adds an item to the end of the queue.
func (q *ConcurrentQueue) Enqueue(item int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
	q.cond.Signal() // Notify one waiting goroutine
}

// Dequeue removes and returns the first item. Blocks if empty.
func (q *ConcurrentQueue) Dequeue() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	for len(q.items) == 0 {
		q.cond.Wait() // Wait for an item to be enqueued
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Len returns the current length (thread-safe).
func (q *ConcurrentQueue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}

// Producer adds items to the queue.
func Producer(q *ConcurrentQueue, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		q.Enqueue(i)
		fmt.Printf("Produced: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

// Consumer removes items from the queue.
func Consumer(q *ConcurrentQueue, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ { // Consume exactly 5 items
		item := q.Dequeue()
		fmt.Printf("Consumed: %d\n", item)
		time.Sleep(200 * time.Millisecond)
	}
}