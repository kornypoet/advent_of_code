package calendar

import (
	"io"
	"strings"
	"testing"
)

func TestDay7(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(""),
			part:     1,
			expected: 1,
		},
		{
			input:    strings.NewReader(""),
			part:     2,
			expected: 2,
		},
	}

	for _, tc := range cases {
		ans := Day7(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
