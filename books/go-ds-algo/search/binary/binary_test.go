package binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	testCases := []struct {
		dataSet []int
		lookFor int
		isFound bool
	} {
		{
			[]int{1, 2, 3, 4, 5, 6},
			4,
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			9,
			false,
		},
	}

	for _, tc := range testCases {
		got := Search(tc.dataSet, tc.lookFor)
		assert.Equal(t, tc.isFound, got)
	}
}

func TestSearchProducts(t *testing.T) {
	testCases := []struct{
		name    string
		input   []Keyer
		lookFor int
		want    Keyer
		isFound bool
	} {
		{
			name: "Sorted Items",
			input: []Keyer{
				&Product{ ID: 1, Name: "Computer"},
				&Product{ ID: 3, Name: "Mouse"},
				&Product{ ID: 4, Name: "Keyboard"},
				&Product{ ID: 6, Name: "Processor"},
			},
			lookFor: 4,
			want: &Product{
				ID:   4,
				Name: "Keyboard",
			},
			isFound: true,
		},
		{
			name: "Unsorted Items",
			input: []Keyer{
				&Product{ ID: 3, Name: "Mouse"},
				&Product{ ID: 6, Name: "Processor"},
				&Product{ ID: 1, Name: "Computer"},
				&Product{ ID: 4, Name: "Keyboard"},
			},
			lookFor: 4,
			want: &Product{
				ID:   4,
				Name: "Keyboard",
			},
			isFound: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			got, isFound := SearchProducts(tc.input, tc.lookFor)
			assert.Equal(t, tc.isFound, isFound)
			assert.Equal(t, tc.want, got)
		})
	}
}

func BenchmarkSearch(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		Search(data, 6)
	}
}

