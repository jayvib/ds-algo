package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPairBrute(t *testing.T) {
	input := []int{1, 2, 3, 4, 6, 2, 3, 3}
	want := [][2]int{
		{1, 3},
		{1, 3},
		{1, 3},
		{2, 2},
	}
	got := FindPairBrute(input, 4)
	assert.Equal(t, want, got)
}

func TestFindPairSorted(t *testing.T) {
	input := []int{1, 2, 3, 4, 6, 2, 3, 3}
	want := [][2]int{
		{1, 3},
		{1, 3},
		{1, 3},
		{2, 2},
	}
	got := FindPairSorted(input, 4)
	assert.Equal(t, want, got)
}
