package calendar

import (
	"io"
	"strings"
	"testing"
)

var input12 string = `
`


func TestDay12(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(input12),
			part:     1,
			expected: 1,
		},
		{
			input:    strings.NewReader(input12),
			part:     2,
			expected: 2,
		},
	}

	for _, tc := range cases {
		ans := Day12(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
