package bubble

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	input := []int{2, 3, 1, 4}
	SortInt(input)

	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(input, want) {
		t.Error("expected not match")
	}
}
