package datastructures

type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type BinarySearchTree[T any] struct {
	root *TreeNode[T]
}

func (bst *BinarySearchTree[T]) Insert(value T, less func(a, b T) bool) {
	if bst.root == nil {
		bst.root = &TreeNode[T]{Value: value}
		return
	}
	insertNode(bst.root, value, less)
}

func insertNode[T any](node *TreeNode[T], value T, less func(a, b T) bool) {
	if less(value, node.Value) {
		if node.Left == nil {
			node.Left = &TreeNode[T]{Value: value}
		} else {
			insertNode(node.Left, value, less)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode[T]{Value: value}
		} else {
			insertNode(node.Right, value, less)
		}
	}
}

func (bst *BinarySearchTree[T]) Search(value T, less func(a, b T) bool) bool {
	return searchNode(bst.root, value, less)
}

func (bst *BinarySearchTree[T]) InorderTraversal(visit func(T)) {
	inorderTraversal(bst.root, visit)
}

func inorderTraversal[T any](node *TreeNode[T], visit func(T)) {
	if node == nil {
		return
	}

	inorderTraversal(node.Left, visit)
	visit(node.Value)
	inorderTraversal(node.Right, visit)
}

func searchNode[T any](node *TreeNode[T], value T, less func(a, b T) bool) bool {
	if node == nil {
		return false
	}

	// Equality is defined by the ordering relation: values are equal when
	// neither one is less than the other.
	if !less(value, node.Value) && !less(node.Value, value) {
		return true
	}
	if less(value, node.Value) {
		return searchNode(node.Left, value, less)
	}
	return searchNode(node.Right, value, less)
}
