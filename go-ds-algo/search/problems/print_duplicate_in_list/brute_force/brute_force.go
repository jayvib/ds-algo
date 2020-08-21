package brute_force

import (
	"fmt"
	"io"
	"os"
)

// Time Complexity:  O(n^2)
// Space Complexity: O(1)

// PrintDuplicate finds the duplicate entry from ints
// and display it in the console.
func PrintDuplicate(ints []int) {
	WriteDuplicatesTo(os.Stdout, ints)
}

// WriteDuplicatesTo searches for the duplicate entry
// from ints and write it to w.
func WriteDuplicatesTo(w io.Writer, ints []int) {
	isFirstWriteDone := false
	size := len(ints)
	for i := 0; i < size; i++ {
		for j := i+1; j < size; j++ {
			if ints[i] == ints[j] {
				// Duplicate
				if isFirstWriteDone {
					writeTo(w, " ")
				}
				writeTo(w, ints[i])
				if !isFirstWriteDone {
					isFirstWriteDone = true
				}
			}
		}
	}
}

func writeTo(w io.Writer, content interface{}) {
	_, _ = fmt.Fprint(w, content)
}
