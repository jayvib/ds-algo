package sort_and_search

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func WriteDuplicatesTo(w io.Writer, ints []int) {
	// Is ints sorted?
	if !sort.IsSorted(sort.IntSlice(ints)) {
		sort.Ints(ints)
	}

	isFirstWriteDone := false
	isDuplicateSetWritten := false
	for i := 1; i < len(ints); i++ {
		if ints[i] == ints[i-1] {
			if isDuplicateSetWritten {
				continue
			}
			if isFirstWriteDone {
				writeTo(w, " ")
			}
			isDuplicateSetWritten = true
			writeTo(w, ints[i])
			if !isFirstWriteDone {
				isFirstWriteDone = true
			}
		} else {
			isDuplicateSetWritten = false
		}
	}
}

func writeTo(w io.Writer, v interface{}) {
	_, _ = fmt.Fprint(w, v)
}

func PrintDuplicate(ints []int) {
	WriteDuplicatesTo(os.Stdout, ints)
}
