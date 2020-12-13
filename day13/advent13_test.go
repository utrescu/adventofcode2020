package main

import (
	"testing"
)

func TestFirstBus2(t *testing.T) {
	var tests = []struct {
		autobusos []autobus
		expected  uint64
	}{
		{
			[]autobus{{7, 0}, {13, 1}, {59, 4}, {31, 6}, {19, 7}},
			1068781,
		},
		// `67,7,59,61` first occurs at timestamp **754018**.
		{
			[]autobus{{67, 0}, {7, 1}, {59, 2}, {61, 3}},
			754018,
		},
		// 67,x,7,59,61` first occurs at timestamp **779210**.
		{
			[]autobus{{67, 0}, {7, 2}, {59, 3}, {61, 4}},
			779210,
		},
	}

	for _, test := range tests {
		if output := firstBus2(test.autobusos); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.autobusos, test.expected, output)
		}
	}

}
