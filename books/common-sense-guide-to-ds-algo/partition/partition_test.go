package partition

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {
	input := []int{0, 5, 2, 1, 6, 3}
	want := []int{0, 1, 2, 3, 6, 5}

	got := Partition(input, 0, len(input)-1)
	assert.Equal(t, want, input	)
	assert.Equal(t, 3, got)
}
