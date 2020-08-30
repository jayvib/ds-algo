package infix_to_postfix

import (
	"github.com/jayvib/ds-algo/books/common-sense-guide-to-ds-algo/data-structure/stacks"
	"strings"
)

type output []string

func (o *output) push(r rune) {
	*o = append(*o, string(r))
}

func (o *output) String() string {
	return strings.Join(*o, " ")
}

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
// Step 5: Drain the remaining operators in the stack and add it to the output list

func InfixToPostfix(input string) string {
	opStack := stacks.NewRuneStack()
	o := make(output, 0)

	for _, r := range input {

		if r == ' ' {
			continue
		}

		// Step 1: If the current character is an operand
		//         add it to the output list
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			// Append to the output list
			o.push(r)
		} else if r == '(' {
			// Step 3: if the character is a '(' then add it to the stack
			opStack.Push(r)
		} else if r == ')' {
			// Step 4: if the character is a ')' then pop all the operators
			//         from the stack until '(' is reach and append it to the
			//         output list
			topOp := opStack.Pop()
			for topOp != '(' {
				o.push(topOp)
				topOp = opStack.Pop()
			}
		} else {
			// Step 2: If the character is an operator,
			//        Cond 1: If the stack is empty or the precedence is lower
			//       					then in the current stack, push it to the stack
			//        Cond 2: If not, then pop the operators that has a higher
			//                precedence and add it to the output list then
			//                push the current operator
			for !opStack.IsEmpty() && precedence(opStack.Peek()) >= precedence(r) {
				// Cond 2
				o.push(opStack.Pop())
			}

			// Cond 1:
			opStack.Push(r)
		}
	}

	// Step 5: Drain the remaining
	for !opStack.IsEmpty() {
		o.push(opStack.Pop())
	}

	return o.String()
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

