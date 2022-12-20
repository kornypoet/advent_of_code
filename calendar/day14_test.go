package calendar

import (
	"io"
	"strings"
	"testing"
)

var input14 string = ``

func TestDay14(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(input14),
			part:     1,
			expected: 1,
		},
		{
			input:    strings.NewReader(input14),
			part:     2,
			expected: 2,
		},
	}

	for _, tc := range cases {
		ans := Day14(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
