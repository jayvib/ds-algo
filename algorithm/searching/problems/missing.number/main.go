package main

import (
	"errors"
)

var ErrNotFound = errors.New("missing element not found")

// Problem: Given a list of n-1 elements, which are in the range of 1 to n.
// There are no duplicates in the list. One of the integer is missing.
func FindMissingElementBrute(elements []int) (missing int, err error) {
	// for each element, find if the selected element
	// exists in another loop.
	// This method is called Brute Force searching
	for i := 1; i < len(elements); i++ { // select the number to compare
		found := false
		for j := 0; j < len(elements); j++ { // this is the comparing occur
			if elements[j] == i {
				found = true
				break
			}
		}

		if !found {
			return i, nil
		}
	}
	return 0, ErrNotFound
}

func main() {

}

func FindMissingElementBruteReview(elements []int) (missing int) {
	for i := 1; i < len(elements); i++ {
		found := false
		for j := 0; i < len(elements); j++ {
			if elements[j] == i {
				found	= true
				break
			}
		}
		if !found {
			return i
		}
	}
	return
}
