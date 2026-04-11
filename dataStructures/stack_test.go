package datastructures

import (
	"errors"
	"testing"
)

func TestStackInt(t *testing.T) {
	stack := Stack[int]{}

	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
	if stack.Size() != 0 {
		t.Errorf("Expected stack size to be 0; got %d", stack.Size())
	}

	// Test Push
	stack.Push(42)
	stack.Push(99)

	if stack.IsEmpty() {
		t.Errorf("Expected stack to be non-empty after push")
	}
	if stack.Size() != 2 {
		t.Errorf("Expected stack size to be 2; got %d", stack.Size())
	}

	// Test Peek
	val, err := stack.Peek()
	if err != nil {
		t.Errorf("Unexpected error from Peek: %v", err)
	}
	if val != 99 {
		t.Errorf("Peek returned wrong value: got %d, want %d", val, 99)
	}

	// Test Pop
	val, err = stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error from Pop: %v", err)
	}
	if val != 99 {
		t.Errorf("Pop returned wrong value: got %d, want %d", val, 99)
	}
	if stack.Size() != 1 {
		t.Errorf("Expected stack size to be 1 after Pop; got %d", stack.Size())
	}

	val, err = stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error from Pop: %v", err)
	}
	if val != 42 {
		t.Errorf("Pop returned wrong value: got %d, want %d", val, 42)
	}

	// Test Pop/Peek on empty stack
	_, err = stack.Pop()
	if !errors.Is(err, errors.New("stack is empty")) && err == nil {
		t.Errorf("Expected error on Pop from empty stack")
	}
	_, err = stack.Peek()
	if !errors.Is(err, errors.New("stack is empty")) && err == nil {
		t.Errorf("Expected error on Peek from empty stack")
	}
}

func TestStackString(t *testing.T) {
	stack := Stack[string]{}
	stack.Push("a")
	stack.Push("b")

	val, err := stack.Peek()
	if err != nil {
		t.Errorf("Unexpected error from Peek: %v", err)
	}
	if val != "b" {
		t.Errorf("Peek returned wrong value: got %s, want %s", val, "b")
	}

	val, err = stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error from Pop: %v", err)
	}
	if val != "b" {
		t.Errorf("Pop returned wrong value: got %s, want %s", val, "b")
	}
	val, err = stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error from Pop: %v", err)
	}
	if val != "a" {
		t.Errorf("Pop returned wrong value: got %s, want %s", val, "a")
	}
}