package selection_sort

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSort(t *testing.T) {
	input := []int{3, 1, 2, 0, 4}
	want := []int{0, 1, 2, 3, 4}
	Sort(input)
	assert.Equal(t, want, input)
}
