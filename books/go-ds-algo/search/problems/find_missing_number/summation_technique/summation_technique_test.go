package summation_technique

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMissingElement(t *testing.T) {
	t.Run("Missing is 5", func(t *testing.T){
		input := []int{
			1, 2, 4, 6, 3, 7, 8,
		}
		want := 5
		got, found := FindMissingElement(input)
		assert.True(t, found)
		assert.Equal(t, want, got)
	})
	t.Run("No Missing", func(t *testing.T){
		input := []int{
			1, 2, 4, 5, 6, 3, 7, 8, 9,
		}
		_, found := FindMissingElement(input)
		assert.False(t, found)
	})
}

