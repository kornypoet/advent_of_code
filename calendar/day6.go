package calendar

import (
	"io"
	"log"

	"github.com/kornypoet/advent_of_code/util"
)

func Day6(input io.Reader, part int) int {
	count := 0
	var bufSize int
	if part == 1 {
		bufSize = 4
	} else {
		bufSize = 14
	}
	util.ProcessByLine(input, func(line string, num int) {
		buffer := make([]rune, 0)
		for index, char := range line {
			for i, c := range buffer {
				if c == char {
					// buffer has this character already, remove it and preceding
					buffer = buffer[i+1:]
				}
			}
			buffer = append(buffer, char)
			if len(buffer) == bufSize {
				log.Printf("buffer %#v", buffer)
				count = index + 1
				break
			}
		}
	})
	log.Printf("Count is %d", count)
	return count
}
