package exponential

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		input := []int{1, 3, 4, 5, 7 , 8, 9, 10}
		got := Search(input, 7)
		assert.Equal(t, 4, got)
	})

	t.Run("Not Found", func(t *testing.T){
		input := []int{1, 3, 4, 5, 7 , 8, 9, 10}
		got := Search(input, 11)
		assert.Equal(t, -1, got)

	})
}
