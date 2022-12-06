package calendar

import (
	"io"
	"strings"
	"testing"
)

var input1 string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestDay1(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(input1),
			part:     1,
			expected: 24000,
		},
		{
			input:    strings.NewReader(input1),
			part:     2,
			expected: 45000,
		},
	}

	for _, tc := range cases {
		ans := Day1(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
