package insertion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortInt(t *testing.T) {
	input := []int{1, 3, 2, 8, 5,4}
	want := []int{1, 2, 3, 4, 5, 8}
	SortInt(input)
	assert.Equal(t, want, input)
}
