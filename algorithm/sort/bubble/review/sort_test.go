package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortInt(t *testing.T) {
	input := []int{2, 1, 3, 4, 5, 6}
	want := []int{1, 2, 3, 4, 5, 6}

	SortInt(input, func(i, j int) bool {
		return i > j
	})

	assert.Equal(t, want , input)
}
