package stacks

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRuneStack_Push(t *testing.T) {
	s := &RuneStack{}
	s.Push('L')
	s.Push('U')
	s.Push('F')
	s.Push('F')
	s.Push('Y')
	assert.Len(t, s.data, 5)
}

func TestRuneStack_Pop(t *testing.T) {
	s := &RuneStack{}
	s.Push('L')
	s.Push('U')
	s.Push('F')
	s.Push('F')
	s.Push('Y')
	r, err := s.Pop()
	require.NoError(t, err)
	assert.Len(t, s.data, 4)
	assert.Equal(t, 'Y', r)
}

