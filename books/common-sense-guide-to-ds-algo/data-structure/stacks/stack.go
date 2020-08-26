package stacks

func NewRuneStack() *RuneStack {
	return &RuneStack{}
}

type RuneStack struct {
	data []rune
}

func (s *RuneStack) Push(r rune) {
	s.data = append(s.data, r)
}

func (s *RuneStack) Pop() (r rune) {
	r, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	return r
}

func (s *RuneStack) Size() int {
	return len(s.data)
}

func (s *RuneStack) IsEmpty() bool {
	return s.Size()==0
}

func (s *RuneStack) Peek() (r rune) {
	return s.data[s.Size()-1]
}
