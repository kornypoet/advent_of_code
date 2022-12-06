package calendar

import (
	"io"
	"strings"
	"testing"
)

var input2 string = `A Y
B X
C Z
`

func TestDay2(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(input2),
			part:     1,
			expected: 15,
		},
		{
			input:    strings.NewReader(input2),
			part:     2,
			expected: 12,
		},
	}

	for _, tc := range cases {
		ans := Day2(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
