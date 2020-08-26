package sort_and_search

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteDuplicatesTo(t *testing.T) {
	t.Run("has duplicate", func(t *testing.T){
		buff := &bytes.Buffer{}
		input := []int{1, 2, 3, 3, 2, 5, 6, 8}
		want := "2 3"
		WriteDuplicatesTo(buff, input)
		got := buff.String()
		assert.Equal(t, want, got)
	})

	t.Run("three duplicate entries", func(t *testing.T){
		buff := &bytes.Buffer{}
		input := []int{1, 2, 3, 3, 3, 2, 5, 5, 5, 5, 6, 8}
		want := "2 3 5"
		WriteDuplicatesTo(buff, input)
		got := buff.String()
		assert.Equal(t, want, got)
	})

	t.Run("has no duplicate", func(t *testing.T){
		buff := &bytes.Buffer{}
		input := []int{1, 2, 3, 5, 6, 8}
		WriteDuplicatesTo(buff, input)
		got := buff.String()
		assert.Empty(t, got)
	})
}

func ExamplePrintDuplicate() {
	input := []int{1, 2, 3, 3, 2, 5, 6, 8}
	PrintDuplicate(input)

	// Output:
	// 2 3
}
