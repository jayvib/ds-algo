package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeInsert(t *testing.T) {
	t.Run("inserting in empty tree", func(t *testing.T) {
		tree := Tree{}
		tree.Insert(45)
		assertSize(t, &tree, 1)
		assert.Equal(t, 45, tree.root.value)
	})

	t.Run("inserting a non empty tree which is lesser than the parent", func(t *testing.T) {
		tree := Tree{}
		tree.Insert(3)
		tree.Insert(2)
		assertSize(t, &tree, 2)
		assertLeftMost(t, &tree, 2)
	})
}

func TestBinaryTreeSearch(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		tree := Tree{}
		tree.Insert(6)
		tree.Insert(5)
		tree.Insert(10)
		tree.Insert(11)

		node := tree.Search(10)
		assert.Equal(t, 10, node.value)
	})

	t.Run("not-found", func(t *testing.T) {
		tree := Tree{}
		tree.Insert(6)
		tree.Insert(5)
		tree.Insert(10)
		tree.Insert(11)

		node := tree.Search(100)
		assert.Nil(t, node)
	})
}

func TestBinaryTreeDelete(t *testing.T) {
	// Cases:
	// Case 1: If the node being deleted has no children, simple delete it.
	// Case 2: If the node being deleted has one child, delete it and plug
	//         the child into the spot where the deleted node was.
	t.Run("Case 1", func(t *testing.T){
		tree := Tree{}
		fillValues(&tree, 6, 5, 7)
		tree.Delete(5)
		assertNotExistingNode(t, tree, 5)
	})

	t.Run("Case 2", func(t *testing.T){
		t.Run("has left child", func(t *testing.T){
			tree := Tree{}
			fillValues(&tree, 6, 5, 7, 4)
			tree.Delete(5)
			assertNotExistingNode(t, tree, 5)
			assertLeft(t, &tree, 6, 4)
		})

		t.Run("has right child", func(t *testing.T){
			tree := Tree{}
			fillValues(&tree, 8, 5, 9, 6)
			tree.Delete(5)
			assertNotExistingNode(t, tree, 5)
			assertLeft(t, &tree, 8, 6)
		})
	})

	t.Run("Case 3", func(t *testing.T){

	})
}

func assertLeft(t *testing.T, tree *Tree, of int, is int) {
	node := tree.Search(of)
	assertNodeValue(t, node.left, is)
}

func assertNodeValue(t *testing.T, node *Node, want int) bool {
	return assert.Equal(t, want, node.value)
}

func fillValues(tree *Tree, values ...int) {
	for _, v := range values {
		tree.Insert(v)
	}
}

func assertNotExistingNode(t *testing.T, tree Tree, x int) {
	got := tree.Search(x)
	assert.Nil(t, got)
}

func assertSize(t *testing.T, tree *Tree, want int) {
	t.Helper()
	assert.Equal(t, want, tree.Size())
}

func assertLeftMost(t *testing.T, tree *Tree, want int) {
	t.Helper()
	n := tree.root
	for n.left != nil {
		n = n.left
	}
	assert.Equal(t, want, n.value)
}
