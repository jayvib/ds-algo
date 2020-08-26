package find_intersection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	arr1 := []int{3, 1, 4, 2}
	arr2 := []int{4, 5, 3, 4, 6 }
	want := []int{3, 4}

	got := FindIntersection(arr1, arr2)
	assert.Equal(t, want, got)
}
