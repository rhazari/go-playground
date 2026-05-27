package datastructures

import (
	"runtime"
	"testing"
)

func TestSPSCQueueFIFO(t *testing.T) {
	q := NewSPSCQueue[int](4)

	if !q.Enqueue(10) || !q.Enqueue(20) || !q.Enqueue(30) {
		t.Fatalf("expected enqueue operations to succeed")
	}

	for _, want := range []int{10, 20, 30} {
		got, ok := q.Dequeue()
		if !ok {
			t.Fatalf("expected value %d to be available", want)
		}
		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	}

	if !q.IsEmpty() {
		t.Fatalf("expected queue to be empty after consuming all items")
	}
}

func TestSPSCQueueFullAndEmpty(t *testing.T) {
	q := NewSPSCQueue[int](2)

	if _, ok := q.Dequeue(); ok {
		t.Fatalf("expected dequeue on empty queue to fail")
	}

	if !q.Enqueue(1) || !q.Enqueue(2) {
		t.Fatalf("expected initial enqueue operations to succeed")
	}

	if q.Enqueue(3) {
		t.Fatalf("expected enqueue on full queue to fail")
	}

	if !q.IsFull() {
		t.Fatalf("expected queue to report full")
	}
}

func TestSPSCQueueConcurrentProducerConsumer(t *testing.T) {
	const total = 10000

	q := NewSPSCQueue[int](128)
	done := make(chan struct{})
	results := make([]int, 0, total)

	go func() {
		for i := 0; i < total; {
			if q.Enqueue(i) {
				i++
				continue
			}
			runtime.Gosched()
		}
		close(done)
	}()

	for len(results) < total {
		if value, ok := q.Dequeue(); ok {
			results = append(results, value)
			continue
		}
		select {
		case <-done:
			if q.IsEmpty() {
				continue
			}
		default:
		}
		runtime.Gosched()
	}

	for i, got := range results {
		if got != i {
			t.Fatalf("out of order value at index %d: got %d, want %d", i, got, i)
		}
	}

	if !q.IsEmpty() {
		t.Fatalf("expected queue to be empty after test")
	}
}

func TestNewSPSCQueuePanicsOnInvalidCapacity(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatalf("expected invalid capacity to panic")
		}
	}()

	NewSPSCQueue[int](0)
}
