package single_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_PushFront(t *testing.T) {
	t.Run("on empty linked list", func(t *testing.T){
		l := LinkedList{}
		input := 1
		l.PushFront(input)
		assertLen(t, l, 1)
		assertNode(t, l, 0, input)
	})

	t.Run("on non-empty linked list", func(t *testing.T){
		l := LinkedList{}
		setupInitialData(&l, 1)

		l.PushFront(2)
		assertLen(t, l, 2)
		assertNode(t, l, 0, 2)
	})
}

func TestList_PushBack(t *testing.T) {
	t.Run("on empty linked list", func(t *testing.T){
		l := LinkedList{}
		input := 1
		l.PushBack(input)
		assertLen(t, l, 1)
		assertNode(t, l, 0, input)
	})

	t.Run("on non-empty linked list", func(t *testing.T){
		l := LinkedList{}
		setupInitialData(&l, 1, 2)
		l.PushBack(3)
		assertLen(t, l, 3)
		assertNode(t, l, 2, 3)
	})
}

func TestList_InsertAt(t *testing.T) {
	// Insert at index 0  -> Front
	// Insert at index -1 -> PushBack
	// Insert at index n  -> InsertAt
	t.Run("index is greater than the current size", func(t *testing.T){
		l := LinkedList{}
		input := 1
		err := l.InsertAt( 99, input)
		assert.Error(t, err)
		assert.Equal(t, ErrIndexBound, err)
	})

	t.Run("index is within the current size", func(t *testing.T){
		l := LinkedList{}
		setupInitialData(&l, 1, 2)
		err := l.InsertAt(1, 99)
		assert.NoError(t, err)
		assertLen(t, l, 3)
		assertNode(t, l, 1, 99)
	})

	t.Run("index is equal to the current size", func(t *testing.T){
		l := LinkedList{}
		err := l.InsertAt(0, 99)
		assert.NoError(t, err)
		assertLen(t, l, 1)
		assertNode(t, l, 0, 99)
	})

	t.Run("insert at front", func(t *testing.T){
		l := LinkedList{}
		setupInitialData(&l, 1, 2)
		err := l.InsertAt(0, 99)
		assert.NoError(t, err)
		assertLen(t, l, 3)
		assertNode(t, l, 0, 99)
	})

	t.Run("insert at end", func(t *testing.T){
		l := LinkedList{}
		setupInitialData(&l, 1, 2, 3)
		err := l.InsertAt(-1, 4)
		assert.NoError(t, err)
		assertLen(t, l, 4)
		assertNode(t, l, -1, 4)
	})
}

func TestList_Get(t *testing.T) {
	t.Run("get existing node", func(t *testing.T){
		l := LinkedList{}
		setupInitialData(&l, 1, 2)
		got, _ := l.Get(1)
		assert.Equal(t, 2, got)
	})

	t.Run("getting a node beyond the current size", func(t *testing.T){
		l := LinkedList{}
		got, err := l.Get(10)
		assert.Error(t, err)
		assert.Equal(t, 0, got)
	})
}

func TestList_Search(t *testing.T) {
	t.Run("existing node", func(t *testing.T){
		l := LinkedList{}
		setupInitialData(&l, 1, 10, 20)
		got := l.Search(10)
		assert.Equal(t, 1, got)
	})
}

// Step 1: Find the node before the node to be deleted
// Step 2: Get the next node of the node to be deleted
// Step 3: Set the next node fo the node before to be the new next node.
//
// Test Cases:
// 1. When the list is empty
// 2. When the node to be deleted is beyond the current size
// 3. When the node to be deleted is within the current size.
// 4. When the node to be deleted is the tail.
func TestList_Delete(t *testing.T) {
	t.Run("when the list is empty", func(t *testing.T){
		l := LinkedList{}
		err := l.Delete(2)
		assert.Error(t, err, "Expecting an error but nothing got")
		assert.Equal(t, ErrIndexBound, err)
	})
	t.Run("when the node to be deleted is beyond the current size", func(t *testing.T){
		l := LinkedList{}
		setupInitialData(&l, 3, 2)
		err := l.Delete(2)
		assert.Error(t, err)
		assert.Equal(t, ErrIndexBound, err)
	})
	t.Run("when the node to be deleted is within the current size", func(t *testing.T){
		l := LinkedList{}
		setupInitialData(&l, 3, 2, 1)
		err := l.Delete(1)
		assert.NoError(t, err)
		assertLen(t, l, 2)
		assertNode(t, l, 1, 1)
	})
	t.Run("when the node to be deleted is the tail", func(t *testing.T){
		l := LinkedList{}
		setupInitialData(&l, 3, 2, 1)
		err := l.Delete(2)
		assert.NoError(t, err)
		assertLen(t, l, 2)
		assert.Equal(t, 2, l.tail.data)
	})
}

func ExamplePrint() {
	l := LinkedList{}
	setupInitialData(&l, 1, 2, 3)

	l.Print()

	// Output:
	// 1 2 3
}

func assertNode(t *testing.T, l LinkedList, index,  want int) {
	node := l.head
	if index == -1 {
		node = l.tail
	}	else {
		for i := 0; i < index; i++ {
			node = node.next
		}
	}
	assert.Equal(t, want, node.data)
}

func assertLen(t *testing.T, l LinkedList, want int) {
	t.Helper()
	assert.Equal(t, want, l.Len())
}

func setupInitialData(l *LinkedList, value ...int) {
	for _, v := range value {
		l.PushBack(v)
	}
}