package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushFront(t *testing.T) {
	t.Run("Adding on first empty tree", func(t *testing.T) {
		l := New()
		l.PushFront("Luffy Monkey")
		assert.NotNil(t, l.head)
		assert.Equal(t, 1, l.size)
	})

	t.Run("Adding value with existing node", func(t *testing.T) {
		l := New()
		l.PushFront("Luffy Monkey")
		l.PushFront("Sanji Vinsmoke")
		assert.Equal(t, 2, l.size)

		currentHeadNode := l.head
		nextNode := currentHeadNode.next
		assert.Equal(t, "Sanji Vinsmoke", currentHeadNode.value)
		assert.Equal(t, "Luffy Monkey", nextNode.value)
	})
}

func TestPushBack(t *testing.T) {
	t.Run("Adding on first empty tree", func(t *testing.T) {
		l := New()
		l.PushBack("Luffy Monkey")
		assert.Equal(t, 1, l.size)
		assert.Equal(t, "Luffy Monkey", l.head.value)
	})

	t.Run("Adding value to the tree with existing node", func(t *testing.T) {
		l := New()
		l.PushBack("Luffy Monkey")
		l.PushBack("Sanji Vinsmoke")
		assert.Equal(t, 2, l.size)

		lastNode := l.head.next
		assert.Equal(t, "Sanji Vinsmoke", lastNode.value)
	})
}
