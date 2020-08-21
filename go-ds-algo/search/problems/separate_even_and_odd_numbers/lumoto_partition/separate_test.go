package lumoto_partition

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeparatOddEven(t *testing.T) {
	input := []int{12, 34, 45, 9, 8, 90, 3}
	want := []int{12, 34, 8, 90, 45, 9, 3}
	SeparateOddEven(input)
	assert.Equal(t, want, input)
}
