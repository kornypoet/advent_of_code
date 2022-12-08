package calendar

import (
	"io"
	"strings"
	"testing"
)

var input7 string = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

func TestDay7(t *testing.T) {
	cases := []struct {
		input    io.Reader
		part     int
		expected int
	}{
		{
			input:    strings.NewReader(input7),
			part:     1,
			expected: 95437,
		},
		// {
		// 	input:    strings.NewReader(input7),
		// 	part:     2,
		// 	expected: 24933642,
		// },
	}

	for _, tc := range cases {
		ans := Day7(tc.input, tc.part)
		if ans != tc.expected {
			t.Errorf("Got %d, expected %d", ans, tc.expected)
		}
	}
}
