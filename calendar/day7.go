package calendar

import (
	"io"
	"log"

	"github.com/kornypoet/advent_of_code/util"
)

func Day7(input io.Reader, part int) int {
	var total int
	if part == 1 {
		total = 1
	} else {
		total = 2
	}

	util.ProcessByLine(input, func(line string, num int) {
		log.Printf("%s %d", line, num)
	})
	log.Printf("Total is %d", total)
	return total
}
