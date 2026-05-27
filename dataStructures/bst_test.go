package datastructures

import "testing"

func TestBinarySearchTreeInsertAndSearchInts(t *testing.T) {
	less := func(a, b int) bool { return a < b }

	var bst BinarySearchTree[int]
	for _, value := range []int{8, 3, 10, 1, 6, 14, 4, 7, 13} {
		bst.Insert(value, less)
	}

	for _, value := range []int{8, 1, 6, 13, 14} {
		if !bst.Search(value, less) {
			t.Fatalf("expected to find %d in tree", value)
		}
	}

	for _, value := range []int{0, 2, 5, 9, 15} {
		if bst.Search(value, less) {
			t.Fatalf("did not expect to find %d in tree", value)
		}
	}
}

func TestBinarySearchTreeSearchEmptyTree(t *testing.T) {
	less := func(a, b int) bool { return a < b }

	var bst BinarySearchTree[int]
	if bst.Search(42, less) {
		t.Fatalf("expected empty tree search to return false")
	}
}

func TestBinarySearchTreeSupportsNonComparableTypes(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}

	less := func(a, b person) bool {
		if a.Age != b.Age {
			return a.Age < b.Age
		}
		return a.Name < b.Name
	}

	var bst BinarySearchTree[person]
	ada := person{Name: "Ada", Age: 37}
	grace := person{Name: "Grace", Age: 45}
	linus := person{Name: "Linus", Age: 29}

	for _, value := range []person{ada, grace, linus} {
		bst.Insert(value, less)
	}

	if !bst.Search(person{Name: "Grace", Age: 45}, less) {
		t.Fatalf("expected to find matching struct value using ordering-based equality")
	}

	if bst.Search(person{Name: "Grace", Age: 46}, less) {
		t.Fatalf("did not expect to find non-matching struct value")
	}
}

func TestBinarySearchTreeHandlesDuplicates(t *testing.T) {
	less := func(a, b int) bool { return a < b }

	var bst BinarySearchTree[int]
	bst.Insert(5, less)
	bst.Insert(5, less)
	bst.Insert(5, less)

	if !bst.Search(5, less) {
		t.Fatalf("expected duplicate value to remain searchable")
	}

	if bst.Search(4, less) {
		t.Fatalf("did not expect unrelated value to be present")
	}
}

func TestBinarySearchTreeInorderTraversalVisitsSortedOrder(t *testing.T) {
	less := func(a, b int) bool { return a < b }

	var bst BinarySearchTree[int]
	for _, value := range []int{8, 3, 10, 1, 6, 14, 4, 7, 13} {
		bst.Insert(value, less)
	}

	var got []int
	bst.InorderTraversal(func(value int) {
		got = append(got, value)
	})

	want := []int{1, 3, 4, 6, 7, 8, 10, 13, 14}
	if len(got) != len(want) {
		t.Fatalf("got %d visited values, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("visit order mismatch at index %d: got %d, want %d", i, got[i], want[i])
		}
	}
}

func TestBinarySearchTreeInorderTraversalEmptyTree(t *testing.T) {
	var bst BinarySearchTree[int]
	visited := false

	bst.InorderTraversal(func(value int) {
		visited = true
	})

	if visited {
		t.Fatalf("expected empty tree traversal to visit no values")
	}
}
