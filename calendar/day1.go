package calendar

import (
	"log"
	"strconv"

	"github.com/kornypoet/advent_of_code/util"
)

var TopN = 3

func filterTopN(candidate int, largest []int) {
	for i, existing := range largest {
		if candidate > existing {
			largest[i] = candidate
			candidate = existing
		}
	}
}

func Day1() {
	total := 0
	largest := []int{0, 0, 0}

	util.ProcessByLine("input/day1.txt", func(line string, num int) {
		if line != "" {
			calories, _ := strconv.Atoi(line)
			total += calories
		} else {
			filterTopN(total, largest)
			total = 0
		}
	})

	log.Printf("Top %d calorie totals: %v", TopN, largest)

	totalCalories := 0
	for _, calories := range largest {
		totalCalories += calories
	}
	log.Printf("Sum of top %d calorie totals: %d", TopN, totalCalories)
}
