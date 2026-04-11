package datastructures

import (
	"errors"
	"testing"
)

// Helper to check queue emptiness and size
func checkQueueState[T any](t *testing.T, q *Queue[T], wantLen int, wantEmpty bool) {
	t.Helper()
	if q.Len() != wantLen {
		t.Errorf("Len() = %d; want %d", q.Len(), wantLen)
	}
	if q.IsEmpty() != wantEmpty {
		t.Errorf("IsEmpty() = %v; want %v", q.IsEmpty(), wantEmpty)
	}
}

func TestQueueInt(t *testing.T) {
	q := Queue[int]{}

	checkQueueState(t, &q, 0, true)

	// Enqueue some items
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	checkQueueState(t, &q, 3, false)

	// Test Front
	front, err := q.Front()
	if err != nil {
		t.Errorf("Front() unexpected error: %v", err)
	}
	if front != 1 {
		t.Errorf("Front() = %d, want 1", front)
	}

	// DeQueue items and check order
	for i, want := range []int{1, 2, 3} {
		v, err := q.DeQueue()
		if err != nil {
			t.Fatalf("DeQueue #%d returned error: %v", i, err)
		}
		if v != want {
			t.Errorf("DeQueue #%d got %d, want %d", i, v, want)
		}
	}
	checkQueueState(t, &q, 0, true)

	// DeQueue and Front on empty queue
	_, err = q.DeQueue()
	if err == nil {
		t.Errorf("Want error on DeQueue() from empty queue")
	}
	_, err = q.Front()
	if err == nil {
		t.Errorf("Want error on Front() from empty queue")
	}
}

func TestQueueString(t *testing.T) {
	q := Queue[string]{}

	q.EnQueue("a")
	q.EnQueue("b")

	val, err := q.Front()
	if err != nil {
		t.Fatalf("Front() error: %v", err)
	}
	if val != "a" {
		t.Errorf("Front() got %q, want %q", val, "a")
	}

	val, err = q.DeQueue()
	if err != nil {
		t.Fatalf("DeQueue() error: %v", err)
	}
	if val != "a" {
		t.Errorf("DeQueue() got %q, want %q", val, "a")
	}
	val, err = q.DeQueue()
	if err != nil {
		t.Fatalf("DeQueue() error: %v", err)
	}
	if val != "b" {
		t.Errorf("DeQueue() got %q, want %q", val, "b")
	}
}

func TestQueueEmptyErrors(t *testing.T) {
	q := Queue[float64]{}

	_, err := q.DeQueue()
	if err == nil {
		t.Errorf("Expected error from DeQueue() on empty queue")
	}
	_, err = q.Front()
	if err == nil {
		t.Errorf("Expected error from Front() on empty queue")
	}

	if !errors.Is(err, errors.New("Queue is empty")) && err != nil {
		// Optional: If you want to check exact error string:
		if err.Error() != "Queue is empty" {
			t.Errorf("Expected 'Queue is empty' error, got: %v", err)
		}
	}
}
