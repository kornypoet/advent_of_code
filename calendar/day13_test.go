package calendar

import (
	"io"
	"strings"
	"testing"
)

var input13 string = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`

func TestDay13(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(input13),
			part:     1,
			expected: 13,
		},
		{
			input:    strings.NewReader(input13),
			part:     2,
			expected: 2,
		},
	}

	for _, tc := range cases {
		ans := Day13(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
