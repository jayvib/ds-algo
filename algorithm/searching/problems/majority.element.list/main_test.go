package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMajorityElementBruteForce(t *testing.T) {
	input := []int{1, 2, 2, 2, 2 , 3, 4}
	want := 2
	got, found := FindMajorityElementBruteForce(input)
	assert.True(t, found)
	assert.Equal(t, want, got)
}

func TestFindMajorityElementSorted(t *testing.T) {
	input := []int{1, 2, 2, 2, 2 , 3, 4}
	want := 2
	got, found := FindMajorityElementSorted(input)
	assert.True(t, found)
	assert.Equal(t, want, got)
}
