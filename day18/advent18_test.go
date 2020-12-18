package main

import (
	"testing"
)

func TestCalcula(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{

		{
			"1 + 2 * 3 + 4 * 5 + 6",
			71,
		},
		{
			"1 + (2 * 3) + (4 * (5 + 6))",
			51,
		},
		{
			"2 * 3 + (4 * 5)",
			26,
		},
		{
			"5 + (8 * 3 + 9 + 3 * 4 * 3)",
			437,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			12240,
		},
		{
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			13632,
		},
	}

	for _, test := range tests {
		if output := calcula(test.input, false); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}

}

func TestCalculaAdvanced(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{

		{
			"1 + 2 * 3 + 4 * 5 + 6",
			231,
		},
		{
			"1 + (2 * 3) + (4 * (5 + 6))",
			51,
		},
		{
			"2 * 3 + (4 * 5)",
			46,
		},
		{
			"5 + (8 * 3 + 9 + 3 * 4 * 3)",
			1445,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			669060,
		},
		{
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			23340,
		},
	}

	for _, test := range tests {
		if output := calcula(test.input, true); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}

}
