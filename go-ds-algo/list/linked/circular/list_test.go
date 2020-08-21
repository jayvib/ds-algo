package circular

import (
	"github.com/jayvib/golog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushFront(t *testing.T) {
	t.Run("When list is empty", func(t *testing.T){
		l := New()
		l.PushFront("a")
		assert.Equal(t, 1, l.size)
		assert.NotNil(t, l.head)
		assert.NotNil(t, l.head.next)
	})

	t.Run("When list has 1 element", func(t *testing.T){
		l := New()
		l.PushFront("a")
		l.PushFront("b")
		gotH := l.head
		assert.Equal(t, 2, l.size)
		assert.Equal(t, "b", gotH.value)
		assert.Equal(t, "a", gotH.next.value)
	})

	t.Run("When list has >2 elements", func(t *testing.T){
		l := New()
		l.PushFront("a")
		l.PushFront("b")
		l.PushFront("c")

		assert.Equal(t, 3, l.size)
		gotH := l.head
		gotT := l.tail
		assert.Equal(t, "c", gotH.value)
		assert.Equal(t, "a", gotT.value)
		assert.Equal(t, gotH, gotT.next)
	})
}

func TestPushBack(t *testing.T) {
	t.Run("When list is empty", func(t *testing.T){
		l := New()
		l.PushBack("a")
		assert.Equal(t, 1, l.size)
		assert.Equal(t, "a", l.head.value)
		assert.Equal(t, l.head, l.head.next)
		assert.Equal(t, l.head, l.tail)
	})

	t.Run("When list as existing 1 elemet", func(t *testing.T){
		l := New()
		l.PushBack("a")
		l.PushBack("b")
		assert.Equal(t, 2, l.size)
		gotT := l.tail
		assert.Equal(t, "b", gotT.value)
		assert.Equal(t, l.head, gotT.next)
	})
}

func TestRemoveHead(t *testing.T) {
	t.Run("Empty List", func(t *testing.T){
		l := New()
		ok := l.RemoveHead()
		assert.False(t, ok)
	})

	t.Run("1 element in a list", func(t *testing.T){
		l := New()
		l.PushBack("a")
		ok := l.RemoveHead()
		assert.True(t, ok)
	})

	t.Run("Removing with >1 elements in a list", func(t *testing.T){
		l := New()
		l.PushBack("a")
		l.PushBack("b")
		ok := l.RemoveHead()
		assert.True(t, ok)
		gotH := l.head
		gotT := l.tail
		assert.Equal(t, "b", gotH.value)
		assert.Equal(t, "b", gotT.value)
		assert.Equal(t, gotT.next, gotH)
	})
}

func TestIsPresent(t *testing.T) {
	t.Run("Found", func(t *testing.T){
		l := New()
		l.PushBack("a")
		l.PushBack("b")
		found := l.IsExists("b")
		assert.True(t, found)
	})

	t.Run("Not Found", func(t *testing.T){
		l := New()
		l.PushBack("a")
		l.PushBack("b")
		found := l.IsExists("d")
		assert.False(t, found)
	})
}

func TestRemoveNode(t *testing.T) {
	t.Run("Empty List", func(t *testing.T){
		l := New()
		ok := l.RemoveNode("any")
		assert.False(t, ok)
	})

	t.Run("With a single node list", func(t *testing.T){
		t.Run("and is found", func(t *testing.T){
			l := New()
			l.PushFront("a")
			ok := l.RemoveNode("a")
			assert.True(t, ok)

			assert.Nil(t, l.head)
			assert.Nil(t, l.tail)
			assert.Empty(t, l.size)
		})
		t.Run("and is not found", func(t *testing.T){
			l := New()
			l.PushFront("a")
			ok := l.RemoveNode("b")
			assert.False(t, ok)
			assert.NotNil(t, l.head)
			assert.NotNil(t, l.tail)
			assert.NotEmpty(t, l.size)
		})
	})

	t.Run("When the node to be remove is at the tail", func(t *testing.T){
		golog.SetLevel(golog.DebugLevel)
		l := New()
		l.PushBack("c")
		l.PushBack("b")
		l.PushBack("a")
		ok := l.RemoveNode("a")
		assert.True(t, ok)
		gotH := l.head
		gotT := l.tail
		assert.Equal(t, "b", gotT.value)
		assert.Equal(t, gotH, gotT.next)
	})

	t.Run("With >1 node list", func(t *testing.T){
		t.Run("is found", func(t *testing.T){
			l := New()
			l.PushFront("a")
			l.PushFront("b")
			l.PushFront("c")
			ok := l.RemoveNode("b")
			assert.True(t, ok)
			assert.Equal(t, 2, l.size)
			gotH := l.head
			gotT := l.tail

			assert.Equal(t, "c", gotH.value)
			assert.Equal(t, "a", gotH.next.value)
			assert.Equal(t, gotH.next, gotT)
		})

		t.Run("is not found", func(t *testing.T){
			l := New()
			l.PushFront("a")
			l.PushFront("b")
			l.PushFront("c")
			ok := l.RemoveNode("d")
			assert.False(t, ok)
		})
	})
}

func TestRemoveDuplicate(t *testing.T) {

	t.Run("Empty List", func(t *testing.T){
		l := New()
		ok := l.RemoveDuplicate()
		assert.False(t, ok)
	})

	t.Run("With 1 Node", func(t *testing.T){
		l := New()
		l.PushFront("a")
		ok := l.RemoveDuplicate()
		assert.False(t, ok)
	})

	t.Run("The duplicate is in the tail", func(t *testing.T){
		l := New()
		l.PushFront("a")
		l.PushFront("a")
		ok := l.RemoveDuplicate()
		assert.True(t, ok)
		assert.Equal(t, 1, l.size)

		gotT := l.tail
		gotH := l.head

		assert.Equal(t, "a", gotT.value)
		if gotT != gotH {
			t.Error("Expecting same pointer address but is not")
		}
	})

	t.Run("The duplicate is in the body", func(t *testing.T){
		 l := New()
		 l.PushBack("a")
		 l.PushBack("b")
		 l.PushBack("b")
		 l.PushBack("b")
		 l.PushBack("c")

		 ok := l.RemoveDuplicate()
		 assert.True(t, ok)
		 assert.Equal(t, 3, l.size)

		 gotH := l.head
		 gotT := l.tail
		 gotB := gotH.next

		 assert.Equal(t, "b", gotH.next.value)
		 assert.Equal(t, "b", gotB.value)
		 assert.Equal(t, "c", gotB.next.value)
		 assert.Equal(t, "a", gotT.next.value)
	})

	t.Run("The duplicate is in the body 2", func(t *testing.T){
		l := New()
		l.PushBack("a")
		l.PushBack("b")
		l.PushBack("b")
		l.PushBack("b")
		l.PushBack("c")
		l.PushBack("c")

		ok := l.RemoveDuplicate()
		assert.True(t, ok)
		assert.Equal(t, 3, l.size)

		gotH := l.head
		gotT := l.tail
		gotB := gotH.next

		assert.Equal(t, "b", gotH.next.value)
		assert.Equal(t, "b", gotB.value)
		assert.Equal(t, "c", gotB.next.value)
		assert.Equal(t, "a", gotT.next.value)
		assert.Equal(t, "c", gotT.value)
	})
}

func ExamplePrint() {
	l := New()
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.Print()

	// Output:
	// a
	// b
	// c
}

func ExampleReverse() {
	l := New()
	l.PushFront("a")
	l.PushFront("b")
	l.PushFront("c")
	l.Reverse()
	l.Print()

	// Output:
	// a
	// b
	// c
}
