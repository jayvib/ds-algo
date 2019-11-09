package quicksort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	input := []int{3, 2, 7, 6, 1, 4}
	want := []int{1, 2, 3, 4, 6, 7}
	Sort(input)
	if !reflect.DeepEqual(input, want) {
		t.Error("not matched")
	}
}
