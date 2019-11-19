package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMaxAppearingElementBruteForce(t *testing.T) {
	input := []int{1, 2, 3, 3, 3, 4, 4, 2, 2, 2,2, 2}
	want := 2
	got := GetMaxAppearingElementBruteForce(input)
	assert.Equal(t, want, got)
}

func TestGetMaxAppearingElementSorted(t *testing.T) {
	t.Run("expecting 2", func(t *testing.T){
		input := []int{1, 2, 3, 3, 3, 4, 4, 2, 2, 2,2, 2}
		want := 2
		got := GetMaxAppearingElementSorted(input)
		assert.Equal(t, want, got)
	})

	t.Run("expecting 5", func(t *testing.T){
		input := []int{1, 5, 3, 3, 3, 4, 4, 5, 5, 5,2, 5}
		want := 5
		got := GetMaxAppearingElementSorted(input)
		assert.Equal(t, want, got)
	})
}
