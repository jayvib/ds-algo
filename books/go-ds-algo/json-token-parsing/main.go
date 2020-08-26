package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Match struct {
	Header string
	Content string
}

type JSONIterator struct {
	decoder *json.Decoder
	startingKey string
	line        int
	start       int
	columns     map[string]int
	paths []string
}

func (j *JSONIterator) Next() ([]*Match, int, error) {
	matches := make([]*Match, 0)
	startingKey := j.startingKey
	isFoundKey := false
	for {
		token, err := j.decoder.Token()
		if err != nil {
			return matches, j.line, err
		}

		switch tv := token.(type) {
		case json.Delim:
			if !isFoundKey {
				continue
			}

			switch tv.String() {
			case "[":
				// Extract values
			}


		case string:
			j.paths = append(j.paths, tv)
			path := toPath(j.paths)
			if path == startingKey {
				isFoundKey = true
			}
		}
	}
	return nil, 0, errors.New("not implemented")
}

func toPath(v []string) string {
	return strings.Join(v, "/")
}

func main() {
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
	d := json.NewDecoder(strings.NewReader(input))
	for {
		token, err := d.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		fmt.Println(token)
	}
}
