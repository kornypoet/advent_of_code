package calendar

import (
	"io"
	"strings"
	"testing"
)

var input8 string = `30373
25512
65332
33549
35390`

func TestDay8(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(input8),
			part:     1,
			expected: 21,
		},
		{
			input:    strings.NewReader(input8),
			part:     2,
			expected: 8,
		},
	}

	for _, tc := range cases {
		ans := Day8(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
