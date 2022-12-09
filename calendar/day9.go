package calendar

import (
	"fmt"
	"io"
	"log"
	"math"
	"regexp"
	"strconv"

	"github.com/kornypoet/advent_of_code/util"
)

func moveTail(tail []int, head []int) []int {
	var position []int
	tailX := tail[0]
	tailY := tail[1]
	headX := head[0]
	headY := head[1]
	xDelta := headX - tailX
	yDelta := headY - tailY
	if math.Abs(float64(xDelta)) > 1 || math.Abs(float64(yDelta)) > 1 {
		// tail must move
		// diagonal means >1 on x or y axis, other
		// either axis even, means only move on the other axis
		if xDelta == 0 {
			// move y axis
			// [0,1] [0,3] d = -2
			if math.Signbit(float64(yDelta)) {
				// delta is negative, move Y axis down
				position = []int{tailX, (tailY - 1)}
			} else {
				// delta is positive, move Y axis up
				position = []int{tailX, (tailY + 1)}
			}
		} else if yDelta == 0 {
			// move x axis
			if math.Signbit(float64(xDelta)) {
				// delta is negative, move X axis down
				position = []int{(tailX - 1), tailY}
			} else {
				// delta is positive, move X axis up
				position = []int{(tailX + 1), tailY}
			}
		} else {
			// move diagonally
			var xPos int
			var yPos int
			if math.Signbit(float64(xDelta)) {
				// delta is negative, move X axis down
				xPos = tailX - 1
			} else {
				// delta is positive, move X axis up
				xPos = tailX + 1
			}
			if math.Signbit(float64(yDelta)) {
				// delta is negative, move Y axis down
				yPos = tailY - 1
			} else {
				// delta is positive, move Y axis up
				yPos = tailY + 1
			}
			position = []int{xPos, yPos}
		}
	} else {
		position = tail
	}
	return position
}

func Day9(input io.Reader, part int) int {
	var total int
	var rope [][]int
	if part == 1 {
		rope = [][]int{ []int{0,0}, []int{0,0} }
	} else {
		rope = [][]int{
			[]int{0,0},
			[]int{0,0},
			[]int{0,0},
			[]int{0,0},
			[]int{0,0},
			[]int{0,0},
			[]int{0,0},
			[]int{0,0},
			[]int{0,0},
			[]int{0,0},
		}
	}
	tailSpots := make(map[string]bool)

	util.ProcessByLine(input, func(line string, num int) {
		dirRexp := regexp.MustCompile(`(\w) (\d+)`)
		directions := dirRexp.FindStringSubmatch(line)
		spaces, _ := strconv.Atoi(directions[2])
		switch directions[1] {
		case "U":
			log.Printf("moving up %d spaces", spaces)
			for i := 1; i <= spaces; i++ {
				log.Print(rope)
				log.Print("1 space")
				head := rope[0]
				head[1]++
				rope[0] = head
				for i, knot := range rope[1:] {
					next := moveTail(knot, head)
					rope[i+1] = next
					log.Print(rope)
					head = next
				}
				tailSpots[fmt.Sprintf("%d:%d", rope[len(rope) - 1][0], rope[len(rope) - 1][1])] = true
			}
		case "D":
			log.Printf("moving down %d spaces", spaces)
			for i := 1; i <= spaces; i++ {
				log.Print(rope)
				log.Print("1 space")
				head := rope[0]
				head[1]--
				rope[0] = head
				for i, knot := range rope[1:] {
					next := moveTail(knot, head)
					rope[i+1] = next
					log.Print(rope)
					head = next
				}
				tailSpots[fmt.Sprintf("%d:%d", rope[len(rope) - 1][0], rope[len(rope) - 1][1])] = true

			}
		case "L":
			log.Printf("moving left %d spaces", spaces)
			for i := 1; i <= spaces; i++ {
				log.Print(rope)
				log.Print("1 space")
				head := rope[0]
				head[0]--
				rope[0] = head
				for i, knot := range rope[1:] {
					next := moveTail(knot, head)
					rope[i+1] = next
					log.Print(rope)
					head = next
				}
				tailSpots[fmt.Sprintf("%d:%d", rope[len(rope) - 1][0], rope[len(rope) - 1][1])] = true
			}
		case "R":
			log.Printf("moving right %d spaces", spaces)
			for i := 1; i <= spaces; i++ {
				log.Print(rope)
				log.Print("1 space")
				head := rope[0]
				head[0]++
				rope[0] = head
				for i, knot := range rope[1:] {
					next := moveTail(knot, head)
					rope[i+1] = next
					log.Print(rope)
					head = next
				}
				tailSpots[fmt.Sprintf("%d:%d", rope[len(rope) - 1][0], rope[len(rope) - 1][1])] = true
			}
		}
	})

	total = len(tailSpots)

	log.Printf("total amount of tail positions is %d", total)
	return total
}
