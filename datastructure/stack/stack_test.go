package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := New()
	s.Push("hello")
	s.Push("world")

	got := s.l.Front().Value
	assert.Equal(t, "world", got)
}


