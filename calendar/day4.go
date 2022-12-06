package calendar

import (
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/kornypoet/advent_of_code/util"
)

func Day4(input io.Reader, part int) int {
	total := 0
	util.ProcessByLine(input, func(line string, num int) {
		sectionPairs := make([][]string, 0)
		assignments := strings.Split(line, ",")
		for _, asst := range assignments {
			sections := strings.Split(asst, "-")
			sectionPairs = append(sectionPairs, sections)
		}
		sOneFirst, _ := strconv.Atoi(sectionPairs[0][0])
		sOneSecond, _ := strconv.Atoi(sectionPairs[0][1])
		sTwoFirst, _ := strconv.Atoi(sectionPairs[1][0])
		sTwoSecond, _ := strconv.Atoi(sectionPairs[1][1])

		if part == 1 {
			if sOneFirst <= sTwoFirst && sOneSecond >= sTwoSecond {
				log.Printf("The first section overlaps the second: %v %v", sectionPairs[0], sectionPairs[1])
				total++
			} else if sTwoFirst <= sOneFirst && sTwoSecond >= sOneSecond {
				log.Printf("The second section overlaps the first: %v %v", sectionPairs[0], sectionPairs[1])
				total++
			}
		} else {
			if sOneFirst <= sTwoFirst && sOneSecond >= sTwoFirst {
				log.Printf("The first section overlaps the second: %v %v", sectionPairs[0], sectionPairs[1])
				total++
			} else if sTwoFirst <= sOneFirst && sTwoSecond >= sOneFirst {
				log.Printf("The second section overlaps the first: %v %v", sectionPairs[0], sectionPairs[1])
				total++
			}
		}
	})
	log.Printf("Total overlaps: %d", total)
	return total
}
