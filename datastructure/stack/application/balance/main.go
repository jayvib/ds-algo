package main

import "github.com/jayvib/ds-algo/stack"

func IsBalance(expn string) bool {
	s := stack.New()
	for _, c := range expn {
		switch c {
		case '{', '[', '(':
			s.Push(c)
		case '}':
			v := s.Pop()
			if v.(rune) != '{' {
				return false
			}
		case ']':
			v := s.Pop()
			if v.(rune) != '[' {
				return false
			}
		case ')':
			v := s.Pop()
			if v.(rune) != '(' {
				return false
			}
		}
	}
	return true
}
