package calendar

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/kornypoet/advent_of_code/util"
)

type Material string

const (
	Rock Material = "#"
	Sand          = "o"
	Air           = "."
	Wall          = "\n"
)

var cave *Cave

type Cave struct {
	smallestX int
	largestX  int
	smallestY int
	largestY  int
	sandCount int
	locations map[string]Material
}

func (c *Cave) addLocation(x int, y int, m Material) {
	coords := fmt.Sprintf("%d:%d", x, y)
	c.locations[coords] = m
	if m == Sand {
		c.sandCount++
	}
	c.expand(x, y)
}

func (c *Cave) expand(x int, y int) {
	if x >= c.largestX {
		c.largestX = x
	}
	if y >= c.largestY {
		c.largestY = y
	}
	if x < c.smallestX {
		c.smallestX = x
	}
	// skip smallest Y as it will always start at 0
}

func (c *Cave) canContinue(x int, y int) bool {
	below := fmt.Sprintf("%d:%d", x, y + 1)
	if _, ok := c.locations[below]; ok {
		// there is an object in the cave directly below this one
		return false
	} else {
		return true
	}
}

func (c *Cave) canShiftLeft(x int, y int) bool {
	left := fmt.Sprintf("%d:%d", x - 1, y + 1)
	if _, ok := c.locations[left]; ok {
		// there is an object in the cave diagonally left below this one
		return false
	} else {
		return true
	}
}

func (c *Cave) canShiftRight(x int, y int) bool {
	right := fmt.Sprintf("%d:%d", x + 1, y + 1)
	if _, ok := c.locations[right]; ok {
		// there is an object in the cave diagonally right below this one
		return false
	} else {
		return true
	}
}

func (c *Cave) pourSand(x int, y int) bool {
	if y >= c.largestY {
		log.Printf("hit largest Y %d", c.largestY)
		return true
	} else {
		if c.canContinue(x, y) {
			return c.pourSand(x, y + 1)
		} else {
			if c.canShiftLeft(x, y) {
				return c.pourSand(x - 1, y + 1)
			} else if c.canShiftRight(x, y) {
				return c.pourSand(x + 1, y + 1)
			} else {
				c.addLocation(x, y, Sand)
				if x == 500 && y == 0 {
					// we've hit the "ceiling"
					return true
				} else {
					return false
				}
			}
		}
	}
}

func (c *Cave) String() string {
	var caveStr Material
	caveStr += Wall
	for y := c.smallestY; y <= c.largestY; y++ {
		for x := c.smallestX; x <= c.largestX; x++ {
			coords := fmt.Sprintf("%d:%d", x, y)
			if material, ok := c.locations[coords]; ok {
				caveStr += material
			} else {
				caveStr += Air
			}
		}
		caveStr += Wall
	}
	return string(caveStr)
}

func drawLine(point1 []int, point2 []int) {
	log.Printf("drawing a line from %v to %v", point1, point2)
	xPoint1, yPoint1, xPoint2, yPoint2 := point1[0], point1[1], point2[0], point2[1]
	if xPoint1 < xPoint2 {
		for x := xPoint1; x <= xPoint2; x++ {
			log.Printf("drawing rock at (%d, %d)", x, yPoint1)
			cave.addLocation(x, yPoint1, Rock)
		}
	} else if xPoint2 < xPoint1 {
		for x := xPoint1; x >= xPoint2; x-- {
			log.Printf("drawing rock at (%d, %d)", x, yPoint1)
			cave.addLocation(x, yPoint1, Rock)
		}
	} else if yPoint1 < yPoint2 {
		for y := yPoint1; y <= yPoint2; y++ {
			log.Printf("drawing rock at (%d, %d)", xPoint1, y)
			cave.addLocation(xPoint1, y, Rock)
		}
	} else {
		for y := yPoint1; y >= yPoint2; y-- {
			log.Printf("drawing rock at (%d, %d)", xPoint1, y)
			cave.addLocation(xPoint1, y, Rock)
		}
	}
}

func Day14(input io.Reader, part int) int {
	// declare cave with arbitrary smallest and largest
	cave = &Cave{smallestX: 1000, smallestY: 0, largestX: 0, largestY: 0, locations: make(map[string]Material)}

	util.ProcessByLine(input, func(line string, num int) {
		rockLines := strings.Split(line, " -> ")
		rockCoords := make([][]int, 0)
		for _, rockLine := range rockLines {
			rocks := strings.Split(rockLine, ",")
			xCoord, _ := strconv.Atoi(rocks[0])
			yCoord, _ := strconv.Atoi(rocks[1])
			coords := []int{xCoord, yCoord}
			rockCoords = append(rockCoords, coords)
		}
		log.Print(rockCoords)
		var previous []int
		for i, rockCoord := range rockCoords {
			if i == 0 {
				previous = rockCoord
				continue
			} else {
				drawLine(previous, rockCoord)
				previous = rockCoord
			}
		}
	})
	if part == 2 {
		// add floor
		y := cave.largestY + 2
		xMin := 500 - y - 1
		xMax := 500 + y + 1
		for x := xMin; x <= xMax; x++ {
			cave.addLocation(x, y, Rock)
		}
	}
	log.Printf("%d:%d -> %d:%d", cave.smallestX, cave.smallestY, cave.largestX, cave.largestY)

	abyss := false
	for !abyss {
		log.Print("dropping sand")
		abyss = cave.pourSand(500, 0)
		// log.Print(cave)
	}
	log.Print(cave.sandCount)
	return cave.sandCount
}
