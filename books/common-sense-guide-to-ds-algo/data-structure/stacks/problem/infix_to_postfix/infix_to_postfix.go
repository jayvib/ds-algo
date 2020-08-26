package infix_to_postfix

import (
	"github.com/jayvib/ds-algo/books/common-sense-guide-to-ds-algo/data-structure/stacks"
	"strings"
)

// Preprocess:
// Prep 1: Create a stack and list, for operator and output, respectively.

// Scan token from left to right and do the following condition:
// Step 1: If the token is an operand, append it to the end of the output list
// Step 2: If the token is an operator push it on the stack. However, first
//         remove any operators already on the stack that have higher or equal
//         precedence and append them to the output list.
// Step 3: If the token is a left parenthesis, push it on the stack.
// Step 4: If the token is a right parenthesis, pop all the operators in the
//         stack until '(' is returned, then add the operators to the output list

func InfixToPostfix(input string) string {
	opStack := stacks.NewRuneStack()
	output := make([]string, 0)

	for _, r := range input {
		if r == ' ' {
			continue
		}

		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			output = append(output, string(r))
		}  else if r == '(' {
			opStack.Push(r)
		} else if r == ')' {
			topOperator := opStack.Pop()
			for topOperator != '(' {
				// Append the operators within the parenthesis.
				output = append(output, string(topOperator))
				topOperator = opStack.Pop()
			}
		} else {
			// Operator
			// Remove first the operators that are greater or equal to with  the current
			// operator.
			for !opStack.IsEmpty() && precedence(opStack.Peek()) >= precedence(r) {
				output = append(output, string(opStack.Pop()))
			}

			// Push the operator
			opStack.Push(r)
		}
	}

	for !opStack.IsEmpty() {
		output = append(output, string(opStack.Pop()))
	}

	return strings.Join(output, " ")
}

func precedence(op rune) int {
	ops := map[rune]int{
		'*': 3,
		'/': 3,
		'+': 2,
		'-': 2,
		'(': 1,
	}
	return ops[op]
}

