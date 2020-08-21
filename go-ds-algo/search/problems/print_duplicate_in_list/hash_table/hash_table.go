package hash_table

import (
	"fmt"
	"io"
	"os"
)

func WriteDuplicatesTo(w io.Writer, ints []int) {
	track := make(map[int]struct{})

	isFirstWriteDone := false
	for _, v := range ints {
		_, ok := track[v]
		if !ok {
			track[v] = struct{}{}
			continue
		}

		if isFirstWriteDone {
			fmt.Fprint(w, " ")
		}

		fmt.Fprintf(w, "%d", v)
		if !isFirstWriteDone {
			isFirstWriteDone = true
		}
	}
}

func PrintDuplicates(ints []int) {
	WriteDuplicatesTo(os.Stdout, ints)
}
