package calendar

import (
	"io"
	"strings"
	"testing"
)

var input5 string = `move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func TestDay5(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected string
		stacks   map[string][]string
	}{
		{
			input:    strings.NewReader(input5),
			part:     1,
			expected: "CMZ",
			stacks: map[string][]string{
				"stack1": []string{"Z", "N"},
				"stack2": []string{"M", "C", "D"},
				"stack3": []string{"P"},
			},
		},
		{
			input:    strings.NewReader(input5),
			part:     2,
			expected: "MCD",
			stacks: map[string][]string{
				"stack1": []string{"Z", "N"},
				"stack2": []string{"M", "C", "D"},
				"stack3": []string{"P"},
			},
		},
	}

	for _, tc := range cases {
		ans := Day5(tc.input, tc.part, tc.stacks)
		if ans != tc.expected {
			t.Errorf("Got %s, expected %s", ans, tc.expected)
		}
	}
}
