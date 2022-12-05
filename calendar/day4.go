package calendar

import (
	"log"
	"strconv"
	"strings"

	"github.com/kornypoet/advent_of_code/util"
)

func Day4() {
	total := 0
	util.ProcessByLine("input/day4.txt", func(line string, num int) {
		log.Printf("line: %s num: %d", line, num)
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

		if sOneFirst <= sTwoFirst && sOneSecond >= sTwoFirst { // 1-5 2-7
			log.Printf("The first section overlaps the second: %v %v", sectionPairs[0], sectionPairs[1])
			total++
		} else if sTwoFirst <= sOneFirst && sTwoSecond >= sOneFirst { // 2-7 1-5
			log.Printf("The first section overlaps the second: %v %v", sectionPairs[0], sectionPairs[1])
			total++
		}
	})
	log.Printf("Total overlaps: %d", total)
}
