package calendar

import (
	"io"
	"strings"
	"testing"
)

var input14 string = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`

func TestDay14(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(input14),
			part:     1,
			expected: 24,
		},
		{
			input:    strings.NewReader(input14),
			part:     2,
			expected: 93,
		},
	}

	for _, tc := range cases {
		ans := Day14(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
