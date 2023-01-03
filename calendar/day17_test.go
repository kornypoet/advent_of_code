package calendar

import (
	"io"
	"strings"
	"testing"
)

var input17 string = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func TestDay17(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(input17),
			part:     1,
			expected: 1,
		},
		// {
		//	input:    strings.NewReader(input17),
		//	part:     2,
		//	expected: 2,
		// },
	}

	for _, tc := range cases {
		ans := Day17(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
