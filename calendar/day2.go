package calendar

import (
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

var gameOutcomes = map[string]map[string]int{
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

func Day2() {
	score := 0

	util.ProcessByLine("input/day2.txt", func(line string, num int) {
		choices := strings.Fields(line)
		opponent := choices[0]
		player := choices[1]
		score += gameOutcomes[opponent][player]
	})

	log.Printf("Total of game scores: %d", score)
}
