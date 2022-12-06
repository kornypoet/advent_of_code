package calendar

import (
	"io"
	"log"
	"strings"

	"github.com/kornypoet/advent_of_code/util"
)

var rock = 1
var paper = 2
var scissors = 3
var loss = 0
var draw = 3
var win = 6

var strategy1 = map[string]map[string]int{
	"A": { // rock
		"X": rock + draw,
		"Y": paper + win,
		"Z": scissors + loss,
	},
	"B": { // paper
		"X": rock + loss,
		"Y": paper + draw,
		"Z": scissors + win,
	},
	"C": { // scissors
		"X": rock + win,
		"Y": paper + loss,
		"Z": scissors + draw,
	},
}

var strategy2 = map[string]map[string]int{
	"A": { // rock
		"X": scissors + loss,
		"Y": rock + draw,
		"Z": paper + win,
	},
	"B": { // paper
		"X": rock + loss,
		"Y": paper + draw,
		"Z": scissors + win,
	},
	"C": { // scissors
		"X": paper + loss,
		"Y": scissors + draw,
		"Z": rock + win,
	},
}

func Day2(input io.Reader, part int) int {
	score := 0
	var strategy map[string]map[string]int
	if part == 1 {
		strategy = strategy1
	} else {
		strategy = strategy2
	}
	util.ProcessByLine(input, func(line string, num int) {
		choices := strings.Fields(line)
		opponent := choices[0]
		player := choices[1]
		score += strategy[opponent][player]
	})

	log.Printf("Total of game scores: %d", score)
	return score
}
