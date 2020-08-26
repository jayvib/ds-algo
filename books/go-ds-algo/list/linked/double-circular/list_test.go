package double_circular

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushFront(t *testing.T) {
	t.Run("Empty list", func(t *testing.T){
		l := New()
		l.PushFront("a")
		assert.Equal(t, 1, l.size)
		assert.NotNil(t, l.head)
		assert.NotNil(t, l.tail)
	})

	t.Run("With 1 existing node", func(t *testing.T){
		l := New()
		l.PushFront("a")
		l.PushFront("b")
		assert.Equal(t, 2, l.size)

		gotH := l.head
		gotT := l.tail

		// Head value
		assert.Equal(t, "b", gotH.value)

		// Tail value
		assert.Equal(t, "a", gotT.value)

		// Tail value
		assert.Equal(t, "a", gotH.next.value)

		// Tail value
		assert.Equal(t, "a", gotH.prev.value)

		// Head value
		assert.Equal(t, "b", gotT.prev.value)
	})

	t.Run("General case", func(t *testing.T){
		l := New()
		l.PushFront("c")
		l.PushFront("b")
		l.PushFront("a")
		assert.Equal(t, 3, l.size)

		gotH := l.head
		gotB := l.head.next
		gotT := l.tail
		assert.Equal(t, "a", gotH.value)
		assert.Equal(t, "b", gotB.value)
		assert.Equal(t, "c", gotT.value)

		// Next and Previous assertion

		// Assert the next and previous of the head
		assert.Equal(t, "b", gotH.next.value)
		assert.Equal(t, "c", gotH.prev.value)

		// Assert the next and previous of the body
		assert.Equal(t, "c", gotB.next.value)
		assert.Equal(t, "a", gotB.prev.value)

		// Assert the next and previous of the tail
		assert.Equal(t, "a", gotT.next.value)
		assert.Equal(t, "b", gotT.prev.value)
	})
}

func TestIsPresent(t *testing.T) {
	t.Run("Found", func(t *testing.T){
		l := New()
		l.PushFront("a")
		l.PushFront("b")
		l.PushFront("c")
		l.PushFront("d")
		got := l.IsPresent("c")
		assert.True(t, got)
	})

	t.Run("Not Found", func(t *testing.T){
		l := New()
		l.PushFront("a")
		l.PushFront("b")
		l.PushFront("c")
		l.PushFront("d")
		got := l.IsPresent("x")
		assert.False(t, got)
	})
}

func TestPushBack(t *testing.T) {
	t.Run("Empty List", func(t *testing.T){
		l := New()
		l.PushBack("a")
		assert.Equal(t, 1, l.size)
		assert.NotNil(t, l.head)
		assert.NotNil(t, l.tail)

		assert.Equal(t, l.head.prev, l.tail)
		assert.Equal(t, l.tail.next, l.head)
	})

	t.Run("With 1 existing element", func(t *testing.T){
		l := New()
		l.PushBack("a")
		l.PushBack("b")
		assert.Equal(t, 2, l.size)

		gotH := l.head
		gotT := l.tail

		assert.Equal(t, "a", gotH.value)
		assert.Equal(t, "b", gotT.value)

		// Prev and Next assertion
		assert.Equal(t, gotT, gotH.next)
		assert.Equal(t, gotT, gotH.prev)
		assert.Equal(t, gotH, gotT.prev)
		assert.Equal(t, gotH, gotT.next)
	})

	t.Run("With existing element is >=2", func(t *testing.T){
		l := New()
		l.PushBack("a")
		l.PushBack("b")
		l.PushBack("c")
		assert.Equal(t, 3, l.size)

		gotH := l.head
		gotB := l.head.next
		gotT := l.tail

		assert.Equal(t, "a", gotH.value)
		assert.Equal(t, "b", gotB.value)
		assert.Equal(t, "c", gotT.value)

		// previous and next
		assert.Equal(t, "c", gotH.prev.value)
		assert.Equal(t, "b", gotH.next.value)

		assert.Equal(t, "a", gotB.prev.value)
		assert.Equal(t, "c", gotB.next.value)

		assert.Equal(t, "b", gotT.prev.value)
		assert.Equal(t, "a", gotT.next.value)
	})
}

func TestRemoveFront(t *testing.T) {
	t.Run("When empty list", func(t *testing.T){
		l := New()
		ok := l.RemoveFront()
		assert.False(t, ok)
	})

	t.Run("With only 1 node", func(t *testing.T){
		l := New()
		l.PushFront("a")
		ok := l.RemoveFront()
		assert.True(t, ok)

		assert.Nil(t, l.head)
		assert.Nil(t, l.tail)
	})

	t.Run("With 2 nodes", func(t *testing.T){
		l := New()
		l.PushBack("a")
		l.PushBack("b")

		ok := l.RemoveFront()
		assert.True(t, ok)

		gotH, gotT := l.head, l.tail

		assert.Equal(t, "b", gotH.value)
		assert.Equal(t, "b", gotT.value)
		assert.Equal(t, "b", gotH.prev.value)
		assert.Equal(t, "b", gotH.next.value)
		assert.Equal(t, "b", gotT.next.value)
		assert.Equal(t, "b", gotT.prev.value)
	})

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