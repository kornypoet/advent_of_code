package calendar

import (
	"log"
	"strings"

	"github.com/kornypoet/advent_of_code/util"
)

type Move int

const (
	Rock     Move = 1
	Paper         = 2
	Scissors      = 3
)

type GameResult int

const (
	Loss GameResult = 0
	Draw            = 3
	Win             = 6
)

var choiceValues = map[string]Move{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var gameOutcomes = map[string]map[string]int{
	"A": { // rock
		"X": 3 + 0, // lose
		"Y": 1 + 3, // draw
		"Z": 2 + 6, // win
	},
	"B": { // paper
		"X": 1 + 0, // lose
		"Y": 2 + 3, // draw
		"Z": 3 + 6, // win
	},
	"C": { // scissors
		"X": 2 + 0, // lose
		"Y": 3 + 3, // draw
		"Z": 1 + 6, // win
	},
}

func Day2() {
	score := 0

	util.ProcessByLine("input/day2.txt", func(line string) {
		choices := strings.Fields(line)
		opponent := choices[0]
		player := choices[1]
		// score += choiceValues[player]
		score += gameOutcomes[opponent][player]
	})

	log.Printf("Total of game scores: %d", score)
}
