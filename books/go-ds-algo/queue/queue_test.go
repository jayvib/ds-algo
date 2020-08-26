package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Empty queue", func(t *testing.T){
		q := &Queue{}
		want := 100
		ok := q.Add(want)
		assert.True(t, ok)
		assert.Equal(t, 1, q.size)
		assert.Equal(t, 1, q.back)

		// TODO: Replace with the Remove method
		got := q.data[0]
		assert.Equal(t, want, got)
	})

	t.Run("Non-empty queue", func(t *testing.T){
		q := &Queue{}
		q.Add(100)
		want := 200
		ok := q.Add(want)
		assert.True(t, ok)
		assert.Equal(t, 2, q.size)
		assert.Equal(t, 2, q.back)
		got := q.data[1]
		assert.Equal(t, want, got)
	})

	t.Run("Full", func(t *testing.T){
		q := &Queue{}
		for i := 0; i < capacity; i++ {
			q.Add(i)
		}

		ok := q.Add(100)
		assert.False(t, ok)
		assert.Equal(t, 0, q.back)
		assert.Equal(t, 1000, q.size)
	})
}

