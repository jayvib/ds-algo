package curly_braces_balancing

import "github.com/jayvib/ds-algo/books/common-sense-guide-to-ds-algo/data-structure/stacks"

func IsBalance(input string) bool {
	s := stacks.NewRuneStack()

	// For each element do the following
	// Step 1: Check if the rune is '{', '[', '('
	// Step 2: Add to stack
	// Step 3: Check if the rune is '}', ']', ')'
	// Step 4: Pop from the stock
	// Step 5: Check if the pop item is its opening brace
	// Step 6: Return true if it is, otherwise false.

	for _, r := range input {
		switch r {
		case '{', '[', '(':
			s.Push(r)
		case '}':
			ok := popFromStackAndCheck(s, '{')
			if !ok {
				return false
			}
		case ']':
			if ok := popFromStackAndCheck(s, '['); !ok {
				return false
			}
		case ')':
			if ok := popFromStackAndCheck(s, '('); !ok {
				return false
			}
		}
	}

	return true
}

func popFromStackAndCheck(s *stacks.RuneStack, r rune) bool {
	d, err := s.Pop()
	if err != nil {
		return false
	}
	if d != r {
		return false
	}
	return true
}
