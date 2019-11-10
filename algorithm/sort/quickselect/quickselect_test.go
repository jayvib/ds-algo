package quickselect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSelect(t *testing.T) {
	input := []int{2, 3, 1, 4, 5}
	want := 2
	got := QuickSelect(input, 1,0, 5)
	assert.Equal(t, want, got)
}