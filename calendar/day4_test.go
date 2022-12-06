package calendar

import (
	"io"
	"strings"
	"testing"
)

var input4 string = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

func TestDay4(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(input4),
			part:     1,
			expected: 2,
		},
		{
			input:    strings.NewReader(input4),
			part:     2,
			expected: 4,
		},
	}

	for _, tc := range cases {
		ans := Day4(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
