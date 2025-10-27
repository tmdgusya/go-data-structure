package tree

import (
	"golang.org/x/exp/constraints"
)

type TreeNode[T constraints.Ordered] struct {
	value  T
	left   *TreeNode[T]
	right  *TreeNode[T]
	parent *TreeNode[T]
}

func (n *TreeNode[T]) InsertNode(value T) {
	if value < n.value {
		if n.left == nil {
			n.left = &TreeNode[T]{
				value:  value,
				parent: n,
			}
		} else {
			n.left.InsertNode(value)
		}
	} else {
		if n.right == nil {
			n.right = &TreeNode[T]{
				value:  value,
				parent: n,
			}
		} else {
			n.right.InsertNode(value)
		}
	}
}

func (n *TreeNode[T]) FindValue(target T) (T, bool) {
	curr := n
	for curr != nil && curr.value != target {
		if target > curr.value {
			curr = curr.right
		} else {
			curr = curr.left
		}
	}
	if curr == nil {
		var zero T
		return zero, false
	}
	return curr.value, true
}

type BinarySearchTree[T constraints.Ordered] struct {
	root *TreeNode[T]
}

func (bst *BinarySearchTree[T]) FindTreeNode(target T) (T, bool) {
	var zero T
	if bst.root == nil {
		return zero, false
	}

	return bst.root.FindValue(target)
}

func (bst *BinarySearchTree[T]) InsertTreeNode(value T) *BinarySearchTree[T] {
	if bst.root == nil {
		bst.root = &TreeNode[T]{
			value: value,
		}
	} else {
		bst.root.InsertNode(value)
	}

	return bst
}

func (bst *BinarySearchTree[T]) RemoveTreeNode(node *TreeNode[T]) {
	if bst.root == nil || node == nil {
		return
	}

	// if node 가 leaf node 일때
	if node.left == nil && node.right == nil {
		if node.parent == nil {
			bst.root = nil
		} else if node.parent.left == node {
			node.parent.left = nil
		} else if node.parent.right == node {
			node.parent.right = nil
		}

		return
	}

	// if node 가 하나의 자식을 가지고 있을때
	if node.left == nil || node.right == nil {
		child := node.left

		if node.left == nil {
			child = node.right
		}

		child.parent = node.parent

		if node.parent == nil {
			bst.root = child
		} else if node.parent.left == node {
			node.parent.left = child
		} else {
			node.parent.right = child
		}

		return
	}

	// if node 가 두개의 자식을 가지고 있을때
	successor := node.right
	for successor.left != nil {
		// 현재 노드보다는 크지만 제일 작은 값을 우측에서 찾음
		successor = successor.left
	}

	bst.RemoveTreeNode(successor)

	if node.parent == nil {
		bst.root = successor
	} else if node.parent.left == node {
		node.parent.left = successor
	} else {
		node.parent.right = successor
	}

	successor.parent = node.parent
	successor.left = node.left
	if node.left != nil {
		node.left.parent = successor
	}

	successor.right = node.right
	if node.right != nil {
		node.right.parent = successor
	}
}
