package double

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Cases to consider:
// 1. Null Values(Empty Tree)
// 2. Only Element(Existing node in the list is the head and tail which are the same node)
// 3. First Element
// 4. General Case
// 5. Last Element
func TestPushFront(t *testing.T) {
	t.Run("Empty Tree", func(t *testing.T) {
		l := New()
		input := "Luffy Monkey"
		l.PushFront(input)
		assert.Equal(t, 1, l.Size())
		gotH := l.head
		gotT := l.tail
		assert.Equal(t, input, gotH.Value())
		assert.Equal(t, input, gotT.Value())
	})

	t.Run("List that has only 1 element", func(t *testing.T) {
		l := New()
		input := "Luffy Monkey"
		input2 := "Sanji Vinsmoke"
		l.PushFront(input2)
		l.PushFront(input)
		assert.Equal(t, 2, l.Size())
		gotH := l.head
		gotT := l.tail
		assert.Equal(t, input, gotH.Value())
		assert.Equal(t, input2, gotT.Value())
	})

	t.Run("List that has an existing size >2", func(t *testing.T) {
		l := New()
		input := "Luffy Monkey"
		input2 := "Sanji Vinsmoke"
		input3 := "Roronoa Zorro"
		l.PushFront(input)
		l.PushFront(input2)
		l.PushFront(input3)
		assert.Equal(t, 3, l.Size())

		gotM := l.head.Next()
		gotH := l.head
		gotT := l.tail

		assert.Equal(t, input, gotT.Value())
		assert.Equal(t, input2, gotM.Value())
		assert.Equal(t, input3, gotH.Value())
		assert.Equal(t, input3, gotM.Previous().Value())
		assert.Equal(t, input2, gotT.Previous().Value())
		assert.Nil(t, gotT.Next())
		assert.Nil(t, gotH.Previous())
	})
}

func TestPushBack(t *testing.T) {
	input1 := "Luffy Monkey"
	input2 := "Sanji Vinsmoke"
	input3 := "Zoro Roronoa"
	t.Run("Empty Tree", func(t *testing.T) {
		l := New()
		l.PushBack(input1)
		gotH := l.head
		gotT := l.tail
		assert.Equal(t, 1, l.Size())
		assert.Equal(t, input1, gotH.Value())
		assert.Equal(t, input1, gotT.Value())
	})

	t.Run("List that has only 1 element", func(t *testing.T) {
		l := New()
		l.PushBack(input1)
		l.PushBack(input2)
		assert.Equal(t, 2, l.Size())
		gotH := l.head
		gotT := l.tail
		assert.Equal(t, input1, gotH.value)
		assert.Equal(t, input2, gotT.value)
		// Test the next and the previous value of the head and the tail
		assert.Equal(t, input2, gotH.next.value)
		assert.Equal(t, input1, gotT.prev.value)
	})

	t.Run("List that has an existing size >2", func(t *testing.T) {
		l := New()
		l.PushBack(input1)
		l.PushBack(input2)
		l.PushBack(input3)
		gotH := l.head
		gotM := l.head.next
		gotT := l.tail
		assert.Equal(t, 3, l.size)
		assert.Equal(t, input1, gotH.value)
		assert.Equal(t, input2, gotM.value)
		assert.Equal(t, input3, gotT.value)

		// Prev and next of the head
		assert.Equal(t, input2, gotH.next.value)
		assert.Nil(t, gotH.prev)

		// Prev and next of the middle
		assert.Equal(t, input3, gotM.next.value)
		assert.Equal(t, input1, gotM.prev.value)

		// Pre and next of the tail
		assert.Equal(t, input2, gotT.prev.value)
		assert.Nil(t, gotT.next)
	})
}

// Logic:
// Starting from the head. Check if the next node is not nil,
// then compare the value of the current node to the next node.
// If the value is the same then make the next's next node of
// the current node be the next node.
// Make the previous node's of the next's next node by the current
// node value.
func ExampleRemoveDuplicate() {
	l := New()
	l.PushFront("a")
	l.PushFront("b")
	l.PushFront("b")
	l.PushFront("b")
	l.PushFront("b")
	l.PushFront("c")
	l.RemoveDuplicate()
	l.Print()

	// Output:
	// c
	// b
	// a
}


func TestRemoveFront(t *testing.T) {
	t.Run("Empty Tree", func(t *testing.T) {
		l := New()
		v, ok := l.RemoveFront()
		assert.False(t, ok)
		assert.Nil(t, v)
	})

	t.Run("Only 1 Element", func(t *testing.T) {
		l := New()
		l.PushFront("a")
		v, ok := l.RemoveFront()
		assert.True(t, ok)
		assert.Equal(t, "a", v)
		assert.Nil(t, l.head)
		assert.Nil(t, l.tail)
		assert.Equal(t, 0, l.size)
	})

	t.Run("2 Elements", func(t *testing.T) {
		l := New()
		l.PushFront("b")
		l.PushFront("a")
		v, ok := l.RemoveFront()
		assert.True(t, ok)
		assert.Equal(t, "a", v)

		// Asserting the current head
		currentHead := l.head
		assert.Equal(t, "b", currentHead.value)
		assert.Nil(t, currentHead.prev)
		assert.Nil(t, currentHead.next)

		currentTail := l.tail
		assert.Equal(t, "b", currentTail.value)
		assert.Nil(t, currentTail.prev)
		assert.Nil(t, currentTail.next)
		assert.Equal(t, 1, l.size)
	})

	t.Run("3 elements", func(t *testing.T) {
		l := New()
		l.PushFront("c")
		l.PushFront("b")
		l.PushFront("a")

		v, ok := l.RemoveFront()
		assert.True(t, ok)

		assert.Equal(t, "a", v)
		assert.Equal(t, 2, l.size)

		gotH := l.head
		gotT := l.tail

		assert.Equal(t, "b", gotH.value)
		assert.Equal(t, "c", gotT.value)

		assert.Equal(t, "c", gotH.next.value)
		assert.Equal(t, "b", gotT.prev.value)
		assert.Nil(t, gotH.prev)
		assert.Nil(t, gotT.next)
	})
}

func TestRemoveNode(t *testing.T) {
	t.Run("Empty List", func(t *testing.T) {
		l := New()
		ok := l.RemoveNode("a")
		assert.False(t, ok)
	})

	t.Run("With only 1 node", func(t *testing.T) {
		l := New()
		l.PushFront("a")
		ok := l.RemoveNode("a")
		assert.True(t, ok)
		assert.Equal(t, 0, l.size)
		assert.Nil(t, l.head)
		assert.Nil(t, l.tail)
	})

	t.Run("The node to remove is at the tail", func(t *testing.T) {
		l := New()
		l.PushFront("b")
		l.PushFront("a")

		ok := l.RemoveNode("b")
		assert.True(t, ok)
		assert.Equal(t, 1, l.size)

		gotH := l.head
		gotT := l.tail

		assert.NotNil(t, gotH)
		assert.Nil(t, gotH.next)
		assert.NotNil(t, gotT)
		assert.Equal(t, "a", gotT.value)
	})

	t.Run("The node to remove is somewhere in the body", func(t *testing.T) {
		l := New()
		l.PushFront("c")
		l.PushFront("b")
		l.PushFront("a")

		ok := l.RemoveNode("b")
		assert.True(t, ok)
		assert.Equal(t, 2, l.size)

		gotH := l.head
		gotT := l.tail

		assert.Equal(t, "c", gotH.next.value)
		assert.Equal(t, "a", gotT.prev.value)
	})

	t.Run("When not found in a list that has 1 element", func(t *testing.T) {
		l := New()
		l.PushFront("a")
		ok := l.RemoveNode("b")
		assert.False(t, ok)
		assert.Equal(t, 1, l.size)
	})

	t.Run("When not found in a list", func(t *testing.T) {
		l := New()
		l.PushFront("d")
		l.PushFront("c")
		l.PushFront("b")
		l.PushFront("a")

		ok := l.RemoveNode("xy")
		assert.False(t, ok)
	})
}

func ExampleReverseList() {
	l := New()
	l.PushFront("c")
	l.PushFront("b")
	l.PushFront("a")
	l.ReverseList()
	l.Print()

	// Output:
	// c
	// b
	// a
}

func ExampleReverseList2() {
	l := New()
	l.PushFront("c")
	l.PushFront("b")
	l.PushFront("a")
	l.reverseList()
	l.Print()

	// Output:
	// c
	// b
	// a
}

func TestIsExists(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		l := New()
		l.PushFront("d")
		l.PushFront("c")
		l.PushFront("b")
		l.PushFront("a")

		found := l.IsExists("d")
		assert.True(t, found)
	})

	t.Run("Not Exists", func(t *testing.T) {
		l := New()
		l.PushFront("d")
		l.PushFront("c")
		l.PushFront("b")
		l.PushFront("a")

		found := l.IsExists("e")
		assert.False(t, found)
	})
}

func TestInsertNode(t *testing.T) {
	t.Run("Insert with the body", func(t *testing.T){
		l := New()
		l.PushBack("a")
		l.PushBack("b")
		l.PushBack("c")
		l.PushBack("d")

		want := "a b x c d"
		ok := l.InsertNode(2, "x")
		assert.True(t, ok)
		assertInsertNode(t, l, want)
	})

	t.Run("Inserting on the tail", func(t *testing.T){
		l := New()
		l.PushBack("a")
		l.PushBack("b")
		l.PushBack("c")
		l.PushBack("d")
		ok := l.InsertNode(l.size, "e")
		assert.True(t, ok)

		want := "a b c d e"
		assertInsertNode(t, l, want)
	})

	t.Run("Inserting on the head", func(t *testing.T) {
		l := New()
		l.PushBack("a")
		l.PushBack("b")
		l.PushBack("c")
		l.PushBack("d")
		ok := l.InsertNode(0, "x")
		assert.True(t, ok)
		want := "x a b c d"
		assertInsertNode(t, l, want)
	})
}

func assertInsertNode(t *testing.T, l *List, want string) {
	t.Helper()

	var gotSlice []string
	l.iterate(func(n *Node)bool{
		gotSlice = append(gotSlice, n.value.(string))
		return false
	})

	got := strings.Join(gotSlice, " ")
	assert.Equal(t, want, got)
}

func ExamplePrint() {
	l := New()
	l.PushFront("a")
	l.PushFront("b")
	l.PushFront("c")
	l.PushFront("d")
	l.Print()

	// Output:
	// d
	// c
	// b
	// a
}

