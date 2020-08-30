package queue

import "fmt"

func Example() {
	q := Queue{}
	q.Push("Luffy")
	q.Push("Sanji")
	q.Push("Zoro")

	for !q.IsEmpty() {
		fmt.Println(q.Front())
	}

	// Output:
	// Luffy
	// Sanji
	// Zoro
}
