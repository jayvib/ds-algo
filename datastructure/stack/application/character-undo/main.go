package main

import (
	"fmt"
	"github.com/jayvib/ds-algo/datastructure/stack"
	"sort"
)

func New(str string) *characterProcessor {
	return &characterProcessor{
		s:          stack.New(),
		characters: []rune(str),
	}
}

type characterProcessor struct {
	s          *stack.Stack
	characters []rune
}

type memento struct {
	value rune
	index int
}

func (p *characterProcessor) Remove(r rune) {
	// Find r.... gonna apply binary search tree
	// gonna use the sort.Search
	index := sort.Search(len(p.characters), func(i int) bool{
		if p.characters[i] == r {
			return true
		}
		return false
	})

	p.s.Push(memento{value: p.characters[index], index: index})
	if index >= 0 {
		p.characters = append(p.characters[:index], p.characters[index+1:]...)
	}
}

func (p *characterProcessor) Undo() {
	val := p.s.Pop()
	mem := val.(memento)
	p.characters = append(p.characters, ' ')
	copy(p.characters[mem.index+1:], p.characters[mem.index:])
	p.characters[mem.index] = mem.value
	// l u f f y
	//     i
	// step 1: l u f f y "
	// step 2: l u f f y "
	// step 3: l u r f f y
}

func (p *characterProcessor) Get() []rune {
	return p.characters
}

func main() {
	c := New("Luffy")
	c.Remove('y')
	fmt.Println(string(c.Get()))
	c.Undo()
	fmt.Println(string(c.Get()))
}
