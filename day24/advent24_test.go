package main

import (
	"testing"
)

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestParse(t *testing.T) {
	var tests = []struct {
		line     string
		expected []string
	}{
		{
			"esenee",
			[]string{"e", "se", "ne", "e"},
		},
		{
			"sesenwnenenewseeswwswswwnenewsewsw",
			[]string{"se", "se", "nw", "ne", "ne", "ne", "w", "se", "e", "sw", "w", "sw", "sw", "w", "ne", "ne", "w", "se", "w", "sw"},
		},
	}

	for _, test := range tests {
		if output := parse(test.line); !Equal(output, test.expected) {
			t.Errorf("Test Failed: %s inputted, %v expected, recieved: %v", test.line, test.expected, output)
		}
	}

}
