package binary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("found", func(t *testing.T){
		data := []int{1, 2, 3, 4 ,5, 7, 8, 9}
		got := Search(data, 8)
		want := 6
		if got != want {
			t.Errorf("want '%d' got '%d'", want, got)
		}
	})

	t.Run("not exist", func(t *testing.T){
		data := []int{1, 2, 3, 4 ,5, 7, 8, 9}
		got := Search(data, 6)
		want := -1
		if got != want {
			t.Errorf("want '%d' got '%d'", want, got)
		}

	})
}
