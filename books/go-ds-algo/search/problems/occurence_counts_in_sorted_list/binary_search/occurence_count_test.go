package binary_search

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestOccurenceCount(t *testing.T) {
	arr := []int{1, 2, 3, 3, 4, 4, 4, 4, 5}
	count := OccurenceCount(arr, 4)
	assert.Equal(t, 4, count)
}
