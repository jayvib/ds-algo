package brute_force

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPair(t *testing.T) {
	t.Run("first test", func(t *testing.T){
		input := []int{1, 2, 3, 4, 5}
		value := 3
		ok := FindPair(input, value)
		assert.True(t, ok)
	})
	t.Run("should not found", func(t *testing.T){
		input := []int{1, 2, 3, 4, 5}
		value := 10
		ok := FindPair(input, value)
		assert.False(t, ok)
	})
}
