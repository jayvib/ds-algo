package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/tree/avl"
)

func main() {
	immu := avl.NewImmutable()
	m1 := mockEntry(2)
	immu, _ = immu.Insert(m1)
	m2 := mockEntry(5)
	immu, _ = immu.Insert(m2)
	m3 := mockEntry(3)
	immu, _ = immu.Insert(m3)
	m4 := mockEntry(4)
	immu, _ = immu.Insert(m4)
	m5 := mockEntry(1)
	immu, _ = immu.Insert(m5)

	entries := immu.Get(m2, m1, m3, m4)
	for _, e := range entries {
		if v, ok := e.(mockEntry); ok {
			fmt.Println(v)
		}
	}
	entries = immu.Get(m5)
	for _, e := range entries {
		if v, ok := e.(mockEntry); ok {
			fmt.Println(v)
		}
	}

}

type mockEntry int

func (m mockEntry) Compare(other avl.Entry) int {
	otherMe := other.(mockEntry)

	if m > otherMe {
		return 1 // right side
	}

	if m < otherMe {
		return -1
	}
	return 0
}