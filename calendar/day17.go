package calendar

import (
	"io"
	"log"

	"github.com/kornypoet/advent_of_code/util"
)

type Piece string

const (
	Filled Piece = "#"
	Empty        = " "
)

type Position struct {
	row      int
	col      int
	material Piece
}

type RockPiece struct {
	shape []*Position
}

func moveLeft(rock *RockPiece) {
	for _, piece := range rock.shape {
		piece.col--
	}
}

func moveRight(rock *RockPiece) {
	for _, pos := range rock.shape {
		pos.col++
	}
}

func moveDown(rock *RockPiece) {
	for _, pos := range rock.shape {
		pos.row--
	}
}

func (r *RockPiece) setHeight(height int) {
	for _, pos := range r.shape {
		pos.row += height
	}
}

func (c *Chamber) canMoveDown(rock *RockPiece) bool {
	return c.collide(rock)
}

func NewRock1() *RockPiece {
	return &RockPiece{
		shape: []*Position{
			&Position{row: 0, col: 3, material: Filled},
			&Position{row: 0, col: 4, material: Filled},
			&Position{row: 0, col: 5, material: Filled},
			&Position{row: 0, col: 6, material: Filled},
		},
	}
}

func NewRock2() *RockPiece {
	return &RockPiece{
		shape: []*Position{
			&Position{row: 0, col: 4, material: Filled},
			&Position{row: 1, col: 3, material: Filled},
			&Position{row: 1, col: 4, material: Filled},
			&Position{row: 1, col: 5, material: Filled},
			&Position{row: 2, col: 4, material: Filled},
		},
	}
}

func NewRock3() *RockPiece {
	return &RockPiece{
		shape: []*Position{
			&Position{row: 0, col: 3, material: Filled},
			&Position{row: 0, col: 4, material: Filled},
			&Position{row: 0, col: 5, material: Filled},
			&Position{row: 1, col: 5, material: Filled},
			&Position{row: 2, col: 5, material: Filled},
		},
	}
}

func NewRock4() *RockPiece {
	return &RockPiece{
		shape: []*Position{
			&Position{row: 0, col: 3, material: Filled},
			&Position{row: 1, col: 3, material: Filled},
			&Position{row: 2, col: 3, material: Filled},
			&Position{row: 3, col: 3, material: Filled},
		},
	}
}

func NewRock5() *RockPiece {
	return &RockPiece{
		shape: []*Position{
			&Position{row: 0, col: 3, material: Filled},
			&Position{row: 0, col: 4, material: Filled},
			&Position{row: 1, col: 3, material: Filled},
			&Position{row: 1, col: 4, material: Filled},
		},
	}
}

//  #  #### # rock1
//  #       #
//  #   #   #
//  #  ###  #
//  #   #   # rock2
//  #       #
//  #    #  #
//  #    #  #
//  #  ###  # rock3
//  #       #
//  #  #    #
//  #  #    #
//  #  #    #
//  #  #    # rock4
//  #       #
//  #  ##   #
//  #  ##   # rock5
// 1#       #
// 0#########
//  012345678

type Chamber struct {
	rows       map[int][]*Position
	maxHeight  int
	directions []rune
}

func (c *Chamber) collide(rock *RockPiece) bool {
	for _, pos := range rock.shape {
		if pos.col <= 0 || pos.col >= 8 {
			// left wall or right wall
			return true
		}
		if row, ok := c.rows[pos.row]; ok {
			// rock is on a row with other pieces
			for _, piece := range row {
				if piece.col == pos.col && piece.material == Filled {
					return true
				}
			}
		}
	}
	return false
}

func (c *Chamber) newRock(count int) *RockPiece {
	rockFactory := map[int]func()(*RockPiece){
		0: NewRock5,
		1: NewRock1,
		2: NewRock2,
		3: NewRock3,
		4: NewRock4,
	}

	rockMaker := rockFactory[count % len(rockFactory)]
	rock := rockMaker()
	rock.setHeight(c.maxHeight + 3)
	return rock
}

func (c *Chamber) addRock(rock *RockPiece) {
	for _, pos := range rock.shape {
		if row, ok := c.rows[pos.row]; ok {
			row = append(row, pos.col)
		} else {
			c.rows[pos.row] = []int{pos.col}
		}
		if pos.row > c.maxHeight {
			c.maxHeight = pos.row
		}
	}
}

func (c *Chamber) gasMoveRock(count int, rock *RockPiece) {
	gasFactory := map[rune]func(*RockPiece){
		'<': moveLeft,
		'>': moveRight,
	}
	dir := c.directions[(count - 1) % len(c.directions)]
	gasFactory[dir](rock)
}

func (c *Chamber) addRow(height int) {
	if r, ok := c.rows[height]; !ok {
		c.maxHeight = height
		c.rows[height] = []*Position{
			&Position{row: height, col: 1, material: Empty},
			&Position{row: height, col: 2, material: Empty},
			&Position{row: height, col: 3, material: Empty},
			&Position{row: height, col: 4, material: Empty},
			&Position{row: height, col: 5, material: Empty},
			&Position{row: height, col: 6, material: Empty},
			&Position{row: height, col: 7, material: Empty},
		}
	}
}

func (c *Chamber) render() {
	out := "\n"
	for height := c.maxHeight; height >= 0; height-- {
		for col := 0; col <= 8; col++ {
			if col == 1 || col == 8 {
				out += "|"
				continue
			}
			row, _ := c.rocks[height]
			out += string(row[col])
		}
		out += "\n"
	}
	log.Print(out)
}

func Day17(input io.Reader, part int) int {
	var total int
	if part == 1 {
		total = 1
	} else {
		total = 2
	}

	floor := map[int][]int{
		0: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	chamber := &Chamber{rocks: floor, maxHeight: 0}

	var directions []rune
	util.ProcessByLine(input, func(line string, num int) {
		for _, char := range line {
			directions = append(directions, char)
		}
	})
	chamber.directions = directions

	for rockCount := 1; rockCount <= 2; rockCount++ {
		rock := chamber.newRock(rockCount)
		for chamber.canMoveDown(rock) {
			chamber.gasMoveRock(rockCount, rock)
			moveDown(rock)
		}
		// rock cannot move down any further
		chamber.addRock(rock)
		chamber.render()
	}

	log.Print(total)
	return total
}
