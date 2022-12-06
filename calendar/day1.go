package calendar

import (
	"io"
	"log"
	"strconv"

	"github.com/kornypoet/advent_of_code/util"
)

func filterTopN(candidate int, largest []int) {
	for i, existing := range largest {
		if candidate > existing {
			largest[i] = candidate
			candidate = existing
		}
	}
}

func Day1(input io.Reader, part int) int {
	total := 0
	var topN int
	if part == 1 {
		topN = 1
	} else {
		topN = 3
	}
	largest := make([]int, topN)

	util.ProcessByLine(input, func(line string, num int) {
		if line != "" {
			calories, _ := strconv.Atoi(line)
			total += calories
		} else {
			filterTopN(total, largest)
			total = 0
		}
	})
	filterTopN(total, largest) // last line in file

	log.Printf("Top %d calorie totals: %v", topN, largest)

	totalCalories := 0
	for _, calories := range largest {
		totalCalories += calories
	}
	log.Printf("Sum of top %d calorie totals: %d", topN, totalCalories)
	return totalCalories
}
