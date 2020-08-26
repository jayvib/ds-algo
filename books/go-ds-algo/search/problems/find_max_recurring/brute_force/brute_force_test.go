package brute_force

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxRecurring(t *testing.T) {
	input := []int{1, 2, 2, 2, 2, 3}
	want := 2
	got := FindMaxRecurring(input)
	assert.Equal(t, want, got)
}

func BenchmarkFindMaxRecurring(b *testing.B) {
	input := []int{1, 2, 2, 2, 2, 3, 3, 4, 5, 23, 2, 2, 3}
	for i := 0; i < b.N; i++ {
		FindMaxRecurring(input)
	}
}