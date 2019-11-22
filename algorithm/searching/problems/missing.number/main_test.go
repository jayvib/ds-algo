package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMissingElement(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
		want := 5
		got, _ := FindMissingElementBrute(input)
		assert.Equal(t, want, got)
	})

	t.Run("Not found", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		_, got := FindMissingElementBrute(input)
		assert.Equal(t, ErrNotFound, got)
	})
}

