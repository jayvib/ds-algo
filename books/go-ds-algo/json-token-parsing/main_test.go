package main

import (
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"strings"
	"testing"
)

var input = `{
	"products": [
		{
			"id": 1,
			"name": "Computer"
		},
		{
			"id": 2,
			"name": "Keyboard"
		},
		{
			"id": 3,
			"name": "Mouse"
		}
	]
}`

func TestJSONIterator(t *testing.T) {
	r := strings.NewReader(input)
	iter := &JSONIterator{
		decoder: json.NewDecoder(r),
		columns: make(map[string]int),
		startingKey: "products",
	}
	want := [][]*Match{
		{
			{Header: "id", Content: "1"},
			{Header: "name", Content: "Computer"},
		},
		{
			{Header: "id", Content: "2"},
			{Header: "name", Content: "Keyboard"},
		},
		{
			{Header: "id", Content: "3"},
			{Header: "name", Content: "Mouse"},
		},
	}

	got := make([][]*Match, 0)
	for {
		match, _, err := iter.Next()
		if err != nil {
			break
		}
		got = append(got, match)
	}
	assert.Equal(t, want, got)
}
