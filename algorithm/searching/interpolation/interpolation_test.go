package interpolation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("found", func(t *testing.T){
		input := []int{1, 3, 4, 5, 7 , 8, 9, 10}
		got := Search(input, 7)
		assert.True(t, got)
	})

	t.Run("not-found", func(t *testing.T){
		input := []int{1, 3, 4, 5, 7 , 8, 9, 10}
		got := Search(input, 100)
		assert.False(t, got)
	})
}
