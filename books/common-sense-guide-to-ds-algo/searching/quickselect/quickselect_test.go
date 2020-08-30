package quickselect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSelect(t *testing.T) {
	input := []int{2, 1, 3, 4, 5, 7, 8}
	want := 7

	got := QuickSelect(input, 5)
	assert.Equal(t, want, got)
}
