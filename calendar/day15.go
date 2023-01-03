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

func manhattanDistance(a []int, b []int) int {
	return int(math.Abs(float64(a[0] - b[0])) + math.Abs(float64(a[1] - b[1])))
}

func Day15(input io.Reader, part int) int {


	horizon := 10 // test
	// horizon := 2000000 // real
	inputRegex := regexp.MustCompile(`Sensor at x=([-\d]+), y=([-\d]+): closest beacon is at x=([-\d]+), y=([-\d]+)`)
	noBeacons := make(map[string]bool)
	beacons := make([]string, 0)
	util.ProcessByLine(input, func(line string, num int) {
		matches := inputRegex.FindStringSubmatch(line)
		sensorX, _ := strconv.Atoi(matches[1])
		sensorY, _ := strconv.Atoi(matches[2])
		beaconX, _ := strconv.Atoi(matches[3])
		beaconY, _ := strconv.Atoi(matches[4])
		sensor := []int{sensorX, sensorY}
		beacon := []int{beaconX, beaconY}
		mDst := manhattanDistance(sensor, beacon)
		log.Printf("sensor %v beacon %v %d", sensor, beacon, mDst)
		coords := fmt.Sprintf("%d:%d", beaconX, beaconY)
		beacons = append(beacons, coords)
		horizonDst := manhattanDistance(sensor, []int{sensorX, horizon})
		if horizonDst <= mDst {
			log.Printf("sensor %v is within %d of the horizon", sensor, mDst)
			delta := int(math.Abs(float64(horizonDst - mDst)))
			for x := sensorX - delta; x<= sensorX + delta; x++ {
				unavailable := fmt.Sprintf("%d:%d", x, horizon)
				noBeacons[unavailable] = true
			}
		}
	})

	for _, b := range beacons {
		if _, ok := noBeacons[b]; ok {
			// beacon overlaps with no beacon
			delete(noBeacons, b)
		}
	}

	log.Print(len(noBeacons))
	return len(noBeacons)
}
