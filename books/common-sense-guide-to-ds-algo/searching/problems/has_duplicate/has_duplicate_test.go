package has_duplicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasDuplicate(t *testing.T) {
	t.Run("found", func(t *testing.T){
		data := []int{ 1, 3, 2, 4, 1, 5}
		got := HasDuplicate(data)
		assert.True(t, got)
	})

	t.Run("not found", func(t *testing.T){
		data := []int{ 1, 3, 2, 4, 5}
		got := HasDuplicate(data)
		assert.False(t, got)
	})
}
