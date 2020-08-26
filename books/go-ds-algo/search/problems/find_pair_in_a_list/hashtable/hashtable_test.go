package hashtable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPair(t *testing.T) {
	t.Run("first test", func(t *testing.T){
		input := []int{1, 2, 3, 4, 5}
		value := 5
		want := [2]int{2, 3}
		got, ok := FindPair(input, value)
		assert.True(t, ok)
		assert.Equal(t, want, got)
	})
	t.Run("should not found", func(t *testing.T){
		input := []int{1, 2, 3, 4, 5}
		value := 10
		_, ok := FindPair(input, value)
		assert.False(t, ok)
	})
}
