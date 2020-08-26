package infix_to_postfix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	testCases := []struct {
		input string
		want string
	} {
		{
			"A * B + C * D",
			"A B * C D * +",
		},
		{
			"( A + B ) * C - ( D - E ) * ( F + G )",
			"A B + C * D E - F G + * -",
		},
	}

	for _, tc := range testCases {
		got := InfixToPostfix(tc.input)
		assert.Equal(t, tc.want, got, "given input: %v", tc.input)
	}
}

