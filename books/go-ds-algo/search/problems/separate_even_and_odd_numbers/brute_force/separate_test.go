package brute_force

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeparatOddEven(t *testing.T) {
	input := []int{12, 34, 45, 9, 8, 90, 3}
	want := []int{12, 34, 90, 8, 9, 45, 3 }
	SeparateOddEven(input)
	assert.Equal(t, want, input)
}
