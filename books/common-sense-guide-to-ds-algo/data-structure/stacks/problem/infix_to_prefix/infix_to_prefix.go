package infix_to_prefix

import "github.com/jayvib/ds-algo/books/common-sense-guide-to-ds-algo/data-structure/stacks/problem/infix_to_postfix"

// Step 1.1: Reverse the string
// Step 1.2: Replace the parenthesis
// Step 2: Do Infix to Postfix
// Step 3: Reverse the string

func InfixToPrefix(input string) string {

	// Step 1: Reverse the input string
	input = reverString(input)

	// Step 2: Do infix to postfix
	output := infix_to_postfix.InfixToPostfix(input)

	// Step 3: Reverse the string
	return reverString(output)
}

func reverString(input string) string {
	inputRunes := []rune(input)
	low := 0
	high := len(inputRunes) - 1
	for low < high {
		inputRunes[low], inputRunes[high] = inputRunes[high], inputRunes[low]
		low++
		high--
	}
	input = string(inputRunes)
	return input
}
