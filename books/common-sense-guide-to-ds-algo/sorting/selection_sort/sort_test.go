package selection_sort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	data := []int{3, 1, 2, 10, 8, 7}
	input := append(make([]int, 0), data...)
	Sort(data)
	want := []int{1, 2, 3, 7, 8, 10}

	if !reflect.DeepEqual(data, want) {
		t.Errorf("want '%v' got '%v', given '%v'", want, data, input)
	}
}
