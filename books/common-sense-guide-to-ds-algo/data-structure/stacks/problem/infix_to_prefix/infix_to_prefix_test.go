package infix_to_prefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInfixToPrefix(t *testing.T) {
	input := "A + B * C + D"
	want := "* + A B + C D"
	got := InfixToPrefix(input)
	assert.Equal(t, want, got)
}
