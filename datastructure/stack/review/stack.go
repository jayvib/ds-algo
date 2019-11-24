package review

import "container/list"

type Stack struct {
	l *list.List
}

func (s *Stack) Pop() (v interface{}) {
	e := s.l.Back()
	return s.l.Remove(e)
}

func (s *Stack) Push(v interface{}) {
	s.l.PushBack(v)
}
