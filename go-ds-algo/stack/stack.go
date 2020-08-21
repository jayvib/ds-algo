package stack

import "container/list"

func New() *Stack {
	return &Stack{
		data: list.New(),
	}
}

type Stack struct {
	data *list.List
}

func (s *Stack) Push(d interface{})  {
	s.data.PushBack(d)
}
func (s *Stack) Pop() (d interface{}) {
	elem := s.data.Back()
	d = s.data.Remove(elem)
	return
}
func (s *Stack) Top() (d interface{}) {
	elem := s.data.Front()
	return elem.Value
}
func (s *Stack) Size() int { return s.data.Len() }
func (s *Stack) IsEmpty() bool { return s.Size()== 0 }