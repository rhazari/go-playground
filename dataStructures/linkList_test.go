package datastructures

import (
	"testing"
)

func TestSinglyLinkedList_AddFirstAndAddLast(t *testing.T) {
	var list SinglyLinkedList[int]

	// Test AddFirst
	list.AddFirst(10) // 10
	list.AddFirst(5)  // 5->10
	if list.Len() != 2 {
		t.Errorf("Len() = %d, want 2", list.Len())
	}
	if v, _ := list.First(); v != 5 {
		t.Errorf("First() = %v, want 5", v)
	}
	if v, _ := list.Last(); v != 10 {
		t.Errorf("Last() = %v, want 10", v)
	}

	// Test AddLast
	list.AddLast(20) // 5->10->20
	if list.Len() != 3 {
		t.Errorf("Len() = %d, want 3", list.Len())
	}
	if v, _ := list.Last(); v != 20 {
		t.Errorf("Last() = %v, want 20", v)
	}
}

func TestSinglyLinkedList_RemoveFirst(t *testing.T) {
	var list SinglyLinkedList[string]
	list.AddLast("a")
	list.AddLast("b")
	list.AddLast("c") // a->b->c

	val, err := list.RemoveFirst()
	if err != nil || val != "a" {
		t.Errorf("RemoveFirst() got (%v, %v), want (a, nil)", val, err)
	}
	val, _ = list.RemoveFirst()
	if val != "b" {
		t.Errorf("RemoveFirst() got %v, want b", val)
	}
	val, _ = list.RemoveFirst()
	if val != "c" {
		t.Errorf("RemoveFirst() got %v, want c", val)
	}
	if !list.IsEmpty() {
		t.Errorf("List should be empty after removing all elements")
	}
	_, err = list.RemoveFirst()
	if err == nil {
		t.Errorf("RemoveFirst() should return error on empty list")
	}
}

func TestSinglyLinkedList_FirstLastOnEmpty(t *testing.T) {
	var list SinglyLinkedList[int]
	_, err := list.First()
	if err == nil {
		t.Errorf("First() should return error when list is empty")
	}
	_, err = list.Last()
	if err == nil {
		t.Errorf("Last() should return error when list is empty")
	}
}

func TestSinglyLinkedList_Clear(t *testing.T) {
	var list SinglyLinkedList[int]
	list.AddLast(1)
	list.AddLast(2)
	list.Clear()
	if !list.IsEmpty() || list.Len() != 0 {
		t.Errorf("Clear() didn't reset the list correctly")
	}
}

func TestSinglyLinkedList_Iterate(t *testing.T) {
	var list SinglyLinkedList[int]
	for i := 1; i <= 5; i++ {
		list.AddLast(i)
	}
	got := []int{}
	list.Iterate(func(x int) {
		got = append(got, x)
	})
	want := []int{1, 2, 3, 4, 5}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Iterate() at index %d: got %d, want %d", i, got[i], want[i])
		}
	}
}