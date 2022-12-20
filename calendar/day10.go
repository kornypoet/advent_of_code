package calendar

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/kornypoet/advent_of_code/util"
)

func Day10(input io.Reader, part int) string {
	var total int

	var cycle int
	if part == 1 {		
		cycle = 1
	} else {
		cycle = 0
	}
	register := 1
	cmdRegexp := regexp.MustCompile(`addx ([\-\d]+)`)
	crt := ""

	util.ProcessByLine(input, func(line string, num int) {
		log.Printf("cycle %d", cycle)

		if part == 1 {
			// 20th, 60th, 100th, 140th, 180th, and 220th
			// fix this awful logic
			if cycle == 20 {
				log.Print("20th cycle")
				signal := 20 * register
				log.Printf("signal is %d", signal)
				total += signal
			} else if cycle == 60 {
				log.Print("60th cycle")
				signal := 60 * register
				log.Printf("signal is %d", signal)
				total += signal
			} else if cycle == 100 {
				log.Print("100th cycle")
				signal := 100 * register
				log.Printf("signal is %d", signal)
				total += signal
			} else if cycle == 140 {
				log.Print("140th cycle")
				signal := 140 * register
				log.Printf("signal is %d", signal)
				total += signal
			} else if cycle == 180 {
				log.Print("180th cycle")
				signal := 180 * register
				log.Printf("signal is %d", signal)
				total += signal
			} else if cycle == 220 {
				log.Print("220th cycle")
				signal := 220 * register
				log.Printf("signal is %d", signal)
				total += signal
			}

			if strings.HasPrefix(line, "noop") {
				log.Print("tick")
				cycle++
			} else {
				command := cmdRegexp.FindStringSubmatch(line)
				log.Printf("perform %s in two cycles", command[1])
				cycle++

				// 20th, 60th, 100th, 140th, 180th, and 220th
				// fix this awful logic
				if cycle == 20 {
					log.Print("20th cycle")
					signal := 20 * register
					log.Printf("signal is %d", signal)
					total += signal
				} else if cycle == 60 {
					log.Print("60th cycle")
					signal := 60 * register
					log.Printf("signal is %d", signal)
					total += signal
				} else if cycle == 100 {
					log.Print("100th cycle")
					signal := 100 * register
					log.Printf("signal is %d", signal)
					total += signal
				} else if cycle == 140 {
					log.Print("140th cycle")
					signal := 140 * register
					log.Printf("signal is %d", signal)
					total += signal
				} else if cycle == 180 {
					log.Print("180th cycle")
					signal := 180 * register
					log.Printf("signal is %d", signal)
					total += signal
				} else if cycle == 220 {
					log.Print("220th cycle")
					signal := 220 * register
					log.Printf("signal is %d", signal)
					total += signal
				}

				cycle++
				compute, _ := strconv.Atoi(command[1])
				register += compute
			}
			log.Printf("register %d", register)
		} else {
			// part 2 cycle starts at 0
			if cycle % 40 == 0 {
				cycle = 0
				crt += "\n"
			}			
			if cycle >= register - 1 && cycle <= register + 1 {
				crt += "#"
			} else {
				crt += "."
			}

			log.Print("\n", crt)
			if strings.HasPrefix(line, "noop") {
				log.Print("tick")
				cycle++

			} else {
				command := cmdRegexp.FindStringSubmatch(line)
				log.Printf("perform %s in two cycles", command[1])
				cycle++
				if cycle % 40 == 0 {
					cycle = 0
					crt += "\n"
				}				
				if cycle >= register - 1 && cycle <= register + 1 {
					crt += "#"
				} else {
					crt += "."
				}
			
				log.Print("\n", crt)				
				cycle++
				compute, _ := strconv.Atoi(command[1])
				register += compute
			}
		}
	})

	if part == 1 {
		log.Print(total)

		return fmt.Sprintf("%d", total)
	} else {
		return crt
	}
}
