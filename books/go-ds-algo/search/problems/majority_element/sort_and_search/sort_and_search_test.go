package sort_and_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMajority(t *testing.T) {
	t.Run("Found", func(t *testing.T){
		input := []int{1, 2, 2, 2, 2, 3}
		want := 2
		got, ok := FindMajority(input)
		assert.True(t, ok)
		assert.Equal(t, want, got)
	})
	t.Run("Not Found", func(t *testing.T){
		input := []int{1, 2, 2, 3}
		want := 0
		got, ok := FindMajority(input)
		assert.False(t, ok)
		assert.Equal(t, want, got)
	})
}
