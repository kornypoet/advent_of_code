package calendar

import (
	"io"
	"strings"
	"testing"
)

var input3 string = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func TestDay3(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(input3),
			part:     1,
			expected: 157,
		},
		{
			input:    strings.NewReader(input3),
			part:     2,
			expected: 70,
		},
	}

	for _, tc := range cases {
		ans := Day3(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
