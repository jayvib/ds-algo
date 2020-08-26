package sort_and_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPair(t *testing.T) {
	t.Run("found", func(t *testing.T){
		setA := []int{0, 1, 2, 3, 4, 5}
		setB := []int{2, 3, 6,  8}
		value := 6
		want := [][2]int{
			{0, 6},
			{3, 3},
			{4, 2},
		}
		got, ok := FindPair(setA, setB, value)
		assert.True(t, ok)
		assert.Equal(t, want, got)
	})
	t.Run("not found", func(t *testing.T){
		setA := []int{0, 1, 2, 3, 4, 5}
		setB := []int{2, 3, 6,  8}
		value := 23
		_, ok := FindPair(setA, setB, value)
		assert.False(t, ok)
	})
}
