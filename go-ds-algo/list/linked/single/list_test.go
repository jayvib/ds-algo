package list

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushFront(t *testing.T) {
	t.Run("Push front on an empty list", func(t *testing.T) {
		l := New()
		l.PushFront("Luffy Monkey")
		assert.Equal(t, 1, l.size)
		assert.NotEmpty(t, l.head)
		assert.Equal(t, "Luffy Monkey", l.head.value)
	})

	t.Run("Push front on a non-empty list", func(t *testing.T) {
		l := New()
		l.PushFront("Luffy Monkey")
		l.PushFront("Sanji Vinsmoke")
		assert.Equal(t, 2, l.size)
		assert.Equal(t, "Sanji Vinsmoke", l.head.value)
	})
}

func TestPushBack(t *testing.T) {
	t.Run("Push back on an empty list", func(t *testing.T) {
		l := New()
		l.PushBack("Luffy Monkey")
		assert.Equal(t, 1, l.size)
	})

	t.Run("Push back on an non-empty list", func(t *testing.T) {
		l := New()
		l.PushBack("Luffy Monkey")
		l.PushBack("Sanji Vinsmoke")
		assert.Equal(t, 2, l.size)
		assert.Equal(t, "Sanji Vinsmoke", l.head.next.value)
	})
}

func TestFindNode(t *testing.T) {
	l := New()
	l.PushBack("Luffy Monkey")
	l.PushBack("Sanji Vinsmoke")
	node, found := l.FindNode("Sanji Vinsmoke")
	require.True(t, found)
	assert.Equal(t, "Sanji Vinsmoke", node.value)
}

func TestRemoveHead(t *testing.T) {
  t.Run("When tree not empty", func(t *testing.T){
		l := New()
		l.PushBack("Luffy Monkey")
		l.PushBack("Sanji Vinsmoke")

		value, ok := l.RemoveFront()
		assert.True(t, ok)
		assert.Equal(t, "Luffy Monkey", value)
	})

  t.Run("When tree is empty", func(t *testing.T){
  	l := New()
  	_, ok := l.RemoveFront()
  	assert.False(t, ok)
	})
}

func TestDeleteNode(t *testing.T) {
  t.Run("Delete node with matches the head, a tree with only 1 node", func(t *testing.T) {
		l := New()
		l.PushBack("Luffy Monkey")
		v, ok := l.DeleteNode("Luffy Monkey")
		assert.True(t, ok)
		assert.Equal(t, "Luffy Monkey", v)
		assert.Nil(t, l.head)
		assert.Nil(t, l.tail)
		assert.Equal(t, 0, l.size)
	})

	t.Run("Delete node with not matches the head a tree with only 1 node", func(t *testing.T) {
		l := New()
		l.PushBack("Luffy Monkey")
		_, ok := l.DeleteNode("Sanji")
		assert.False(t, ok)
		assert.NotNil(t, l.head)
		assert.NotNil(t, l.tail)
		assert.Equal(t, 1, l.size)
	})


	t.Run("Delete node somewhere in the body", func(t *testing.T){
		l := New()
		l.PushBack("Luffy Monkey")
		l.PushBack("Roronoa Zoro")
		l.PushBack("Sanji Vinsmoke")
		v, ok := l.DeleteNode("Sanji Vinsmoke")
		assert.True(t, ok)
		assert.Equal(t, "Sanji Vinsmoke", v)
		assert.Equal(t, 2, l.size)
	})
  t.Run("Delete node with no found", func(t *testing.T){
		l := New()
		l.PushBack("Luffy Monkey")
		l.PushBack("Roronoa Zoro")
		l.PushBack("Sanji Vinsmoke")
  	v, ok := l.DeleteNode("notexist")
  	assert.False(t, ok)
  	assert.Nil(t, v)
	})
  t.Run("Delete node in an empty tree", func(t *testing.T){
  	l := New()
  	v, ok := l.DeleteNode("notexist")
  	assert.False(t, ok)
  	assert.Nil(t, v)
	})
  t.Run("Deleted node is in the tail", func(t *testing.T){
		l := New()
		l.PushBack("Luffy Monkey")
		l.PushBack("Roronoa Zoro")
		l.PushBack("Sanji Vinsmoke")
		_, ok := l.DeleteNode("Sanji Vinsmoke")
		assert.True(t, ok)
		assert.Equal(t, "Roronoa Zoro", l.tail.value)
	})
}

func TestDeleteNodes(t *testing.T) {
	l := New()
	input := "Luffy Monkey"
	l.PushBack(input)
	l.PushBack(input)
	l.PushBack(input)
	l.PushBack(input)
	ok := l.DeleteNodes(input)
	assert.True(t, ok)
	assert.Equal(t, 0, l.size)
}

func TestCompare(t *testing.T) {
	t.Run("Equal", func(t *testing.T){
		l1 := New()
		l2 := New()
		input := []string{"a", "b", "c", "d"}
		for _, v := range input {
			l1.PushBack(v)
			l2.PushBack(v)
		}

		assert.True(t, l1.Compare(l2))
	})

	t.Run("Not Equal", func(t *testing.T){
		l1 := New()
		l2 := New()
		input := []string{"a", "b", "c", "d"}
		for _, v := range input {
			l1.PushBack(v)
		}
		l2.PushBack("3")
		assert.False(t, l1.Compare(l2))
	})
}

//                  -b--c--d-
//                /     ^     \
// --a--b--c--d--e      |      e
//         ^      \   detect  /
//         |        -d--c--b-
//       detect
func TestLoopDetect(t *testing.T) {
	t.Run("Found Hash Map", func(t *testing.T){
		l1 := New()
		l1.PushBack("a")
		l1.PushBack("b")
		l1.PushBack("c")
		l1.PushBack("d")
		l1.PushBack("e")

		// To Create a loop
		node, _ := l1.FindNode("b")
		lastNode := l1.Tail()
		lastNode.next = node

		got := l1.LoopDetect()
		assert.True(t, got)
	})

	t.Run("Floyd Cycle", func(t *testing.T) {
		l1 := New()
		l1.PushBack("a")
		l1.PushBack("b")
		l1.PushBack("c")
		l1.PushBack("x")
		l1.PushBack("d")
		l1.PushBack("e")

		// To Create a loop
		node, _ := l1.FindNode("b")
		lastNode := l1.Tail()
		lastNode.next = node

		got := l1.LoopDetectFloydCycle()
		assert.True(t, got)
	})

	t.Run("Not Found Hash Map", func(t *testing.T){
		l1 := New()
		l1.PushBack("a")
		l1.PushBack("b")
		l1.PushBack("c")
		l1.PushBack("d")

		got := l1.LoopDetect()
		assert.False(t, got)
	})

	t.Run("Reverse", func(t *testing.T){
		l1 := New()
		l1.PushBack("a")
		l1.PushBack("b")
		l1.PushBack("c")
		l1.PushBack("x")
		l1.PushBack("d")
		l1.PushBack("e")

		// To Create a loop
		node, _ := l1.FindNode("b")
		lastNode := l1.Tail()
		lastNode.next = node
		got := l1.LoopDetectReverse()
		assert.True(t, got)
		l1.Print()
	})
}

func TestRemoveLoop(t *testing.T) {
	l1 := New()
	l1.PushBack("a")
	l1.PushBack("b")
	l1.PushBack("c")
	l1.PushBack("x")
	l1.PushBack("d")
	l1.PushBack("e")

	// To Create a loop
	node, _ := l1.FindNode("b")
	lastNode := l1.Tail()
	lastNode.next = node

	ok := l1.RemoveLoop()
	assert.True(t, ok)

	foundLoop := l1.LoopDetectFloydCycle()
	assert.False(t, foundLoop)
}

func ExampleCopy() {
	l := New()
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.PushBack("d")

	nl := l.Copy()
	nl.Print()

	// Output:
	// a
	// b
	// c
	// d
}

func ExampleCopyListReversed() {
	l := New()
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.PushBack("d")

	nl := l.CopyListReversed()

	nl.Print()

	// Output:
	// d
	// c
	// b
	// a
}

func ExampleRemoveDuplicateNodes() {
	l := New()
	l.PushBack("a")
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("b")
	l.RemoveDuplicateNodes()
	l.Print()

	// Output:
	// a
	// b
}

func ExampleReverse() {
	l := New()
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.Reverse()
	l.Print()

	// Output:
	// c
	// b
	// a
}

func ExampleReverse2() {
	l := New()
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.reverse()
	l.Print()

	// Output:
	// c
	// b
	// a
}

func ExampleRecursiveReverse() {
	l := New()
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.recursiveReverse()
	l.Print()

	// Output:
	// c
	// b
	// a
}

func BenchmarkPushBack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := New()
		l.PushBack("Luffy Monkey")
		l.PushBack("Luffy Monkey")
		l.PushBack("Luffy Monkey")
		l.PushBack("Luffy Monkey")
		l.PushBack("Luffy Monkey")
		l.PushBack("Luffy Monkey")
		l.PushBack("Sanji Vinsmoke")
	}
}
