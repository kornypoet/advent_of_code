package calendar

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"

	"github.com/kornypoet/advent_of_code/util"
)

var tailPositions map[string]bool

func recordTailPosition(tail []int) {
	key := fmt.Sprintf("%d:%d", tail[0], tail[1])
	tailPositions[key] = true
}

func moveKnot(head []int, tail []int) (position []int) {
	xPos := tail[0]
	yPos := tail[1]
	xDelta := head[0] - xPos
	yDelta := head[1] - yPos
	if xDelta < -1 || xDelta > 1 || yDelta < -1 || yDelta > 1 {
		if xDelta < 0 {
			xPos--
		} else if xDelta > 0 {
			xPos++
		}

		if yDelta < 0 {
			yPos--
		} else if yDelta > 0 {
			yPos++
		}
	}
	position = []int{xPos, yPos}
	return
}

func moveRope(rope [][]int) {
	head := rope[0]
	for i, knot := range rope[1:] {
		next := moveKnot(head, knot)
		rope[i+1] = next
		log.Print(rope)
		head = next
	}
	recordTailPosition(rope[len(rope)-1])
}

func Day9(input io.Reader, part int) int {
	var total int
	var rope [][]int
	tailPositions = make(map[string]bool)

	if part == 1 {
		rope = [][]int{
			make([]int, 2),
			make([]int, 2),
		}
	} else {
		rope = [][]int{
			make([]int, 2),
			make([]int, 2),
			make([]int, 2),
			make([]int, 2),
			make([]int, 2),
			make([]int, 2),
			make([]int, 2),
			make([]int, 2),
			make([]int, 2),
			make([]int, 2),
		}
	}

	util.ProcessByLine(input, func(line string, num int) {
		cmdRexp := regexp.MustCompile(`(\w) (\d+)`)
		command := cmdRexp.FindStringSubmatch(line)
		direction := command[1]
		spaces, _ := strconv.Atoi(command[2])
		switch direction {
		case "U":
			log.Printf("moving %s %d spaces", direction, spaces)
			for i := 1; i <= spaces; i++ {
				log.Print(rope)
				log.Print("1 space")
				head := &rope[0]
				(*head)[1]++
				moveRope(rope)
			}

		case "D":
			log.Printf("moving %s %d spaces", direction, spaces)
			for i := 1; i <= spaces; i++ {
				log.Print(rope)
				log.Print("1 space")
				head := &rope[0]
				(*head)[1]--
				moveRope(rope)
			}

		case "L":
			log.Printf("moving %s %d spaces", direction, spaces)
			for i := 1; i <= spaces; i++ {
				log.Print(rope)
				log.Print("1 space")
				head := &rope[0]
				(*head)[0]--
				moveRope(rope)
			}

		case "R":
			log.Printf("moving %s %d spaces", direction, spaces)
			for i := 1; i <= spaces; i++ {
				log.Print(rope)
				log.Print("1 space")
				head := &rope[0]
				(*head)[0]++
				moveRope(rope)
			}
		}
	})

	total = len(tailPositions)

	log.Printf("total amount of tail positions is %d", total)
	return total
}
