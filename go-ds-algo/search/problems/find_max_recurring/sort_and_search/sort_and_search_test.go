package sort_and_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxRecurring(t *testing.T) {
	t.Run("Has more than 1 recurring", func(t *testing.T){
		input := []int{1, 2, 2, 2, 2, 3}
		want := 2
		got := FindMaxRecurring(input)
		assert.Equal(t, want, got)
	})

	t.Run("Has no recurring", func(t *testing.T){
		input := []int{2, 3, 1}
		want := 1
		got := FindMaxRecurring(input)
		assert.Equal(t, want, got)
	})
}

func benchmarkFindMaxRecurring(b *testing.B, input []int) {
	for i := 0; i < b.N; i++ {
		FindMaxRecurring(input)
	}
}

func BenchmarkFindMaxRecurring_With_20_Elements(b *testing.B) {
	input := []int{1, 2, 2, 2, 2, 3, 3, 4, 5, 23, 2, 2, 3}
	benchmarkFindMaxRecurring(b, input)
}

