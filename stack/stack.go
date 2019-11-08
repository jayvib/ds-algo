package stack

import "container/list"

func New() *Stack {
	return &Stack{l: list.New()}
}

type Stack struct {
	l *list.List
}

func (s *Stack) Pop() (value interface{}) {
	e := s.l.Front()
	return s.l.Remove(e)
}

func (s *Stack) Push(value interface{}) {
	s.l.PushFront(value)
}

func (s *Stack) Len() int {
	return s.l.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}


