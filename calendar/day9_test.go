package calendar

import (
	"io"
	"strings"
	"testing"
)

var input9 string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

var input9p2 string = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

func TestDay9(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(input9),
			part:     1,
			expected: 13,
		},
		{
			input:    strings.NewReader(input9p2),
			part:     2,
			expected: 36,
		},
	}

	for _, tc := range cases {
		ans := Day9(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
