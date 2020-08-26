package curly_braces_balancing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBalance(t *testing.T) {
	input := "( var x = { y: [1, 2, 3] } )"
	got := IsBalance(input)
	assert.True(t, got)
}
