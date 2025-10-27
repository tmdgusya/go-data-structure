package tree

import "testing"

func TestBinarySearchTree_InsertAndFind(t *testing.T) {
	bst := &BinarySearchTree[int]{}

	// Insert values
	bst.InsertTreeNode(10).
		InsertTreeNode(5).
		InsertTreeNode(15).
		InsertTreeNode(3).
		InsertTreeNode(7).
		InsertTreeNode(12).
		InsertTreeNode(20)

	// Test finding existing values
	tests := []struct {
		name   string
		target int
		want   int
		found  bool
	}{
		{"Find root", 10, 10, true},
		{"Find left child", 5, 5, true},
		{"Find right child", 15, 15, true},
		{"Find leaf left", 3, 3, true},
		{"Find leaf right", 20, 20, true},
		{"Find middle", 7, 7, true},
		{"Find non-existent", 100, 0, false},
		{"Find non-existent negative", -5, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := bst.FindTreeNode(tt.target)
			if found != tt.found {
				t.Errorf("FindTreeNode(%d) found = %v, want %v", tt.target, found, tt.found)
			}
			if found && got != tt.want {
				t.Errorf("FindTreeNode(%d) = %v, want %v", tt.target, got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_EmptyTree(t *testing.T) {
	bst := &BinarySearchTree[int]{}

	value, found := bst.FindTreeNode(10)
	if found {
		t.Errorf("Expected not to find value in empty tree, but found %v", value)
	}
}

func TestBinarySearchTree_SingleNode(t *testing.T) {
	bst := &BinarySearchTree[int]{}
	bst.InsertTreeNode(42)

	value, found := bst.FindTreeNode(42)
	if !found {
		t.Error("Expected to find value 42")
	}
	if value != 42 {
		t.Errorf("Expected value 42, got %v", value)
	}

	_, notFound := bst.FindTreeNode(100)
	if notFound {
		t.Error("Should not find non-existent value")
	}
}

func TestBinarySearchTree_RemoveLeafNode(t *testing.T) {
	bst := &BinarySearchTree[int]{}
	bst.InsertTreeNode(10).InsertTreeNode(5).InsertTreeNode(15).InsertTreeNode(3).InsertTreeNode(7)

	// Remove leaf node (3)
	node := bst.root.left.left // node with value 3
	bst.RemoveTreeNode(node)

	// Verify it's removed
	_, found := bst.FindTreeNode(3)
	if found {
		t.Error("Node 3 should be removed")
	}

	// Verify other nodes still exist
	if _, found := bst.FindTreeNode(5); !found {
		t.Error("Node 5 should still exist")
	}
	if _, found := bst.FindTreeNode(10); !found {
		t.Error("Node 10 should still exist")
	}
}

func TestBinarySearchTree_RemoveNodeWithOneChild(t *testing.T) {
	bst := &BinarySearchTree[int]{}
	bst.InsertTreeNode(10).InsertTreeNode(5).InsertTreeNode(15).InsertTreeNode(3)

	// Remove node with one child (5, which has left child 3)
	node := bst.root.left // node with value 5
	bst.RemoveTreeNode(node)

	// Verify 5 is removed but 3 is still accessible
	if _, found := bst.FindTreeNode(5); found {
		t.Error("Node 5 should be removed")
	}
	if _, found := bst.FindTreeNode(3); !found {
		t.Error("Node 3 should still be accessible")
	}
	if _, found := bst.FindTreeNode(10); !found {
		t.Error("Node 10 should still exist")
	}
}

func TestBinarySearchTree_RemoveNodeWithTwoChildren(t *testing.T) {
	bst := &BinarySearchTree[int]{}
	bst.InsertTreeNode(10).
		InsertTreeNode(5).
		InsertTreeNode(15).
		InsertTreeNode(3).
		InsertTreeNode(7).
		InsertTreeNode(12).
		InsertTreeNode(20)

	// Remove node with two children (5, which has children 3 and 7)
	node := bst.root.left // node with value 5
	bst.RemoveTreeNode(node)

	// Verify 5 is removed but its children are still accessible
	if _, found := bst.FindTreeNode(5); found {
		t.Error("Node 5 should be removed")
	}
	if _, found := bst.FindTreeNode(3); !found {
		t.Error("Node 3 should still be accessible")
	}
	if _, found := bst.FindTreeNode(7); !found {
		t.Error("Node 7 should still be accessible")
	}
	if _, found := bst.FindTreeNode(10); !found {
		t.Error("Root node 10 should still exist")
	}
}

func TestBinarySearchTree_RemoveRoot(t *testing.T) {
	bst := &BinarySearchTree[int]{}
	bst.InsertTreeNode(10).InsertTreeNode(5).InsertTreeNode(15)

	// Remove root
	bst.RemoveTreeNode(bst.root)

	// Verify root is replaced
	if bst.root == nil {
		t.Error("Tree should not be empty after removing root with children")
	}
	if _, found := bst.FindTreeNode(10); found {
		t.Error("Original root (10) should be removed")
	}
	if _, found := bst.FindTreeNode(5); !found {
		t.Error("Node 5 should still be accessible")
	}
	if _, found := bst.FindTreeNode(15); !found {
		t.Error("Node 15 should still be accessible")
	}
}

func TestBinarySearchTree_RemoveRootLeafNode(t *testing.T) {
	bst := &BinarySearchTree[int]{}
	bst.InsertTreeNode(10)

	// Remove root when it's the only node
	bst.RemoveTreeNode(bst.root)

	// Verify tree is empty
	if bst.root != nil {
		t.Error("Tree should be empty after removing single root node")
	}

	_, found := bst.FindTreeNode(10)
	if found {
		t.Error("Should not find any value in empty tree")
	}
}

func TestBinarySearchTree_RemoveNilNode(t *testing.T) {
	bst := &BinarySearchTree[int]{}
	bst.InsertTreeNode(10).InsertTreeNode(5).InsertTreeNode(15)

	// Should not panic when removing nil node
	bst.RemoveTreeNode(nil)

	// Verify tree is unchanged
	if _, found := bst.FindTreeNode(10); !found {
		t.Error("Tree should be unchanged")
	}
}

func TestBinarySearchTree_WithStrings(t *testing.T) {
	bst := &BinarySearchTree[string]{}

	bst.InsertTreeNode("dog").
		InsertTreeNode("cat").
		InsertTreeNode("elephant").
		InsertTreeNode("ant").
		InsertTreeNode("zebra")

	tests := []struct {
		target string
		want   bool
	}{
		{"dog", true},
		{"cat", true},
		{"elephant", true},
		{"ant", true},
		{"zebra", true},
		{"bear", false},
	}

	for _, tt := range tests {
		t.Run("Find_"+tt.target, func(t *testing.T) {
			_, found := bst.FindTreeNode(tt.target)
			if found != tt.want {
				t.Errorf("FindTreeNode(%s) found = %v, want %v", tt.target, found, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_WithFloat64(t *testing.T) {
	bst := &BinarySearchTree[float64]{}

	bst.InsertTreeNode(3.14).
		InsertTreeNode(2.71).
		InsertTreeNode(1.41).
		InsertTreeNode(9.81)

	value, found := bst.FindTreeNode(2.71)
	if !found {
		t.Error("Expected to find 2.71")
	}
	if value != 2.71 {
		t.Errorf("Expected 2.71, got %v", value)
	}

	_, notFound := bst.FindTreeNode(99.99)
	if notFound {
		t.Error("Should not find 99.99")
	}
}
