package calendar

import (
	"io"
	"strings"
	"testing"
)

func TestDay6(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
			part:     1,
			expected: 7,
		},
		{
			input:    strings.NewReader("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			part:     2,
			expected: 26,
		},
	}

	for _, tc := range cases {
		ans := Day6(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
