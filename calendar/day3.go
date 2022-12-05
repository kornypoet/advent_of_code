package calendar

import (
	"log"

	"github.com/kornypoet/advent_of_code/util"
)

func priorityMap() map[rune]int {
	priorityMap := make(map[rune]int)
	priority := 1
	var r rune
	for r = 'a'; r <= 'z'; r++ {
		priorityMap[r] = priority
		priority++
	}
	for r = 'A'; r <= 'Z'; r++ {
		priorityMap[r] = priority
		priority++
	}
	return priorityMap
}

func Day3() {
	total := 0
	pm := priorityMap()
	// util.ProcessByLine("input/day3.txt", func(line string) {
	//	duplicates := make(map[rune]bool)
	//	for i, char := range line {
	//		if i < len(line) / 2 { // compartment one
	//			duplicates[char] = true
	//		} else { // compartment two
	//			if duplicates[char] {
	//				log.Printf("Found the duplicate item: %c", char)
	//				log.Printf("The priority is: %d", pm[char])
	//				total += pm[char]
	//				break
	//			}
	//		}
	//	}
	// })
	// log.Printf("The total priority is %d", total)
	group := 1
	// groupPriority := make(map[int]int)
	duplicates := make(map[rune]int)
	util.ProcessByLine("input/day3.txt", func(line string, num int) {
		groupMember := num % 3
		log.Printf("Group %d, member %d", group, groupMember)
		for _, char := range line {
			if groupMember == 1 {
				// always add all of member 1
				duplicates[char] = 1
			} else if groupMember == 2 {
				if duplicates[char] == 1 {
					duplicates[char] = 2

				}
			} else if groupMember == 0 {
				if duplicates[char] == 2 {
					log.Printf("We found the group duplicate %c", char)
					total += pm[char]
					break
				}
			}
		}
		if groupMember == 0 {
			duplicates = make(map[rune]int)
			group++
		}
	})
	log.Printf("The total priority is %d", total)
}
