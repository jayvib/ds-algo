package linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	input := []int{1, 3, 4, 5, 7 , 8, 9, 10}
	got := Search(input, 7)
	assert.True(t, got)
}
