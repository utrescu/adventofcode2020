package main

import (
	"testing"
)

func TestPlay(t *testing.T) {
	var tests = []struct {
		numbers  []int
		expected int
	}{
		{
			[]int{0, 3, 6},
			436,
		},
		{
			[]int{1, 3, 2},
			1,
		},
		{
			[]int{2, 1, 3},
			10,
		},
		{
			[]int{1, 2, 3},
			27,
		},
		{
			[]int{2, 3, 1},
			78,
		},
		{
			[]int{3, 2, 1},
			438,
		},
		{
			[]int{3, 1, 2},
			1836,
		},
	}

	for _, test := range tests {
		if output := soluciona1(test.numbers, 2020); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.numbers, test.expected, output)
		}
	}

}
