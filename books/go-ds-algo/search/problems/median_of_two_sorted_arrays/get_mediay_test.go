package median_of_two_sorted_arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMedian(t *testing.T) {
	arr1 := []int{1, 12, 15, 26, 38}
	arr2 := []int{2, 13, 17, 30, 45}
	got := GetMedian(arr1, arr2, len(arr1))
	want := 16
	assert.Equal(t, want, got)
}
