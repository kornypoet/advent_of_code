package calendar

import (
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/kornypoet/advent_of_code/util"
)

func scenic(k int, j int, grid [][]int) int {
	// <- score
	height := grid[k][j]
	scoreLeft := 0
	for i := j - 1; i >= 0; i-- {
		tree := grid[k][i]
		if height > tree {
			scoreLeft++
		} else if height <= tree {
			scoreLeft++
			break
		}
	}

	// -> score
	scoreRight := 0
	for i := j + 1; i < len(grid[k]); i++ {
		tree := grid[k][i]
		if height > tree {
			scoreRight++
		} else if height <= tree {
			scoreRight++
			break
		}
	}

	// ^ score
	scoreUp := 0
	for i := k - 1; i >= 0; i-- {
		tree := grid[i][j]
		if height > tree {
			scoreUp++
		} else if height <= tree {
			scoreUp++
			break
		}
	}

	// down score
	scoreDown := 0
	for i := k + 1; i < len(grid); i++ {
		tree := grid[i][j]
		if height > tree {
			scoreDown++
		} else if height <= tree {
			scoreDown++
			break
		}
	}
	log.Printf("%d:%d %d*%d*%d*%d", k, j, scoreLeft, scoreRight, scoreDown, scoreUp)
	return scoreLeft * scoreRight * scoreDown * scoreUp
}

func Day8(input io.Reader, part int) int {
	var total int
	grid := make([][]int, 0)
	util.ProcessByLine(input, func(line string, num int) {
		row := make([]int, 0)
		for _, ch := range line {
			height, _ := strconv.Atoi(string(ch))
			row = append(row, height)
		}
		grid = append(grid, row)
	})
	log.Print(grid)

	if part == 1 {
		coordinates := make([][]int, 0)
		visible := -1
		log.Print("rows")
		for k, row := range grid {
			// left rows
			log.Printf("processing row %d", k)
			visible = -1
			for j, tree := range row {
				log.Printf("processing column %d", j)
				if tree > visible {
					coordinates = append(coordinates, []int{k, j})
					visible = tree
				}
				// if visible == 9 {
				//	break
				// }
			}
			// right rows
			log.Printf("processing row %d", k)
			visible = -1
			for j := len(row) - 1; j >= 0; j-- {
				log.Printf("processing column %d", j)
				tree := row[j]
				if tree > visible {
					coordinates = append(coordinates, []int{k, j})
					visible = tree
				}
				// if visible == 9 {
				//	break
				// }
			}
		}

		log.Print("columns")
		for j := 0; j < len(grid[0]); j++ {
			log.Printf("processing column %d", j)
			// top columns
			visible = -1
			for k := 0; k < len(grid); k++ {
				log.Printf("processing row %d", k)
				tree := grid[k][j]
				if tree > visible {
					coordinates = append(coordinates, []int{k, j})
					visible = tree
				}
				// if visible == 9 {
				//	break
				// }
			}
			// bottom columns
			visible = -1
			for k := len(grid) - 1; k >= 0; k-- {
				log.Printf("processing row %d", k)
				tree := grid[k][j]
				if tree > visible {
					coordinates = append(coordinates, []int{k, j})
					visible = tree
				}
				// if visible == 9 {
				//	break
				// }
			}
		}

		log.Print(coordinates)
		uniq := make(map[string]bool)
		for _, coord := range coordinates {
			strCoord := fmt.Sprintf("%d:%d", coord[0], coord[1])
			uniq[strCoord] = true
		}
		log.Print(uniq)
		total = len(uniq)
		log.Printf("total visible trees: %d", total)
	} else {

		mostScenic := 0
		// scenic score
		for k, row := range grid {
			// left rows
			log.Printf("processing row %d", k)
			for j := 0; j < len(row); j++ {
				log.Printf("processing column %d", j)
				s := scenic(k, j, grid)
				if s > mostScenic {
					mostScenic = s
				}
				log.Printf("scenic score %d", s)
			}
		}

		log.Print("columns")
		for j := 0; j < len(grid[0]); j++ {
			log.Printf("processing column %d", j)
			// top columns
			for k := 0; k < len(grid); k++ {
				log.Printf("processing row %d", k)
				s := scenic(k, j, grid)
				if s > mostScenic {
					mostScenic = s
				}
				log.Printf("scenic score %d", s)
			}
		}

		total = mostScenic
		log.Printf("Most scenic score: %d", mostScenic)
	}

	return total
}
