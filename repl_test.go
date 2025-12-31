package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello world some",
			expected: []string{"hello", "world", "some"},
		},
		{
			input:    "world",
			expected: []string{"world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected len: %v, got: %v", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if !reflect.DeepEqual(word, expectedWord) {
				t.Errorf("expected: %v, got %v", expectedWord, word)
			}
		}
	}
}
