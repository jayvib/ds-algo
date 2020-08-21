package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("existing element", func(t *testing.T){
		slice := Slice{
			1, 2, 3, 4, 5,
		}
		i, found := slice.Search(3)
		assert.Equal(t, 2, i)
		assert.True(t, found)
	})
	t.Run("non-existing element", func(t *testing.T){
		slice := Slice{
			1, 2, 3, 4, 5,
		}
		i, found := slice.Search(10)
		assert.Equal(t, 0, i)
		assert.False(t, found)
	})
}

func TestUpdate(t *testing.T) {
}

func ExampleInsert() {
	slice := Slice{
		1, 2, 3, 4, 5,
	}
	slice.Insert(6, 2)

	slice.Print()
	// Output:
	// 1 2 6 3 4 5
}

func ExampleDelete() {
	slice := Slice{
		1, 2, 3, 4, 5,
	}
	slice.Delete(1)

	slice.Print()
	// Output:
	// 1 3 4 5
}

func ExampleTraverse() {
	arr := Slice{
		1, 3, 5, 7, 8, 9, 2, 4, 0,
	}
	arr.Print()

	// Output:
	// 1 3 5 7 8 9 2 4 0
}
