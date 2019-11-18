package linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	input := []int{1, 3, 4, 5, 7 , 8, 9, 10}
	got := Search(input, 7)
	assert.True(t, got)
}

func TestSearchSorted(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		input := []int{1, 3, 4, 5, 7 , 8, 9, 10}
		got := SearchSorted(input, 7)
		assert.True(t, got)
	})

	t.Run("Not Found", func(t *testing.T){
		input := []int{1, 3, 4, 5, 7 , 8, 9, 10}
		got := SearchSorted(input, 11)
		assert.False(t, got)

	})
}
