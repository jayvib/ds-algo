package jump_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("found", func(t *testing.T){
		input := []int{1, 2, 4, 5, 8, 9, 10}
		i, ok := Search(input, 8)
		assert.True(t, ok)
		assert.Equal(t, 4, i)
	})

	t.Run("not found", func(t *testing.T){
		input := []int{1, 2, 4, 5, 8, 9, 10}
		i, ok := Search(input, 22)
		assert.False(t, ok)
		assert.Equal(t, -1, i)
	})
}
