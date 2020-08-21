package main

import "fmt"

type Slice []int

func (s Slice) Print() {
	s.traverse(func(i, v int)bool{
		fmt.Printf("%d", s[i])
		if i != s.Size()-1 {
			fmt.Print(" ")
		}
		return false
	})
}

func (s Slice) Search(value int) (index int, found bool ) {
	return linearSearch(s, value)
}

func linearSearch(s Slice, value int) (int, bool) {
	for i := 0; i < s.Size(); i++ {
		if s[i] == value {
			return i, true
		}
	}
	return 0, false
}

func (s *Slice) Delete(index int) {
	*s = append((*s)[:index], (*s)[index+1:]...)
}

// 1 2 3 4
func (s *Slice) Insert(value, index int) {
	s.insert(index, value)
}

func (s *Slice) insert(index int, value int) {
	*s = append((*s)[:index+1], (*s)[index:]...)
	(*s)[index] = value
}

func (s Slice) traverse(fn func(i, v int)bool) {
	for i := range s {
		if fn(i, s[i]) {
			return
		}
	}
}

func (s Slice) Size() int {
	return len(s)
}

func main() {

}
