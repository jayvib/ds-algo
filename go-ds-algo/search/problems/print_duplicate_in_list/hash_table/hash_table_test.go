package hash_table

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteDuplicatesTo(t *testing.T) {
	t.Run("Has Duplicate", func(t *testing.T){
		input := []int{1, 1, 2, 2, 3, 3}
		buff := &bytes.Buffer{}
		WriteDuplicatesTo(buff, input)
		want := "1 2 3"
		assert.Equal(t, want, buff.String())
	})

	t.Run("No Duplicate", func(t *testing.T){
		input := []int{1, 2, 3}
		buff := &bytes.Buffer{}
		WriteDuplicatesTo(buff, input)
		want := ""
		assert.Equal(t, want, buff.String())
	})
}
