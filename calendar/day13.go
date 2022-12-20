package calendar

import (
	"fmt"
	"io"
	"log"
	"reflect"
	"regexp"
	"sort"

	"github.com/kornypoet/advent_of_code/util"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type packetOrder []any

func (a packetOrder) Len() int      { return len(a) }
func (a packetOrder) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a packetOrder) Less(i, j int) bool {
	less, cont := evaluatePackets(a[i], a[j])
	if cont {
		log.Fatal("this should never be true")
	}
	return less
}

func evaluatePackets(left any, right any) (bool, bool) {
	leftType := reflect.TypeOf(left).Kind()
	rightType := reflect.TypeOf(right).Kind()

	if leftType == rightType {
		if leftType == reflect.Int {
			log.Print("comparing ints")
			leftVal, _ := left.(int)
			rightVal, _ := right.(int)
			if leftVal < rightVal {
				return true, false
			} else if leftVal == rightVal {
				return true, true
			} else {
				return false, false
			}
		} else {
			log.Print("comparing slices")

			leftVal, _ := left.([]any)
			rightVal, _ := right.([]any)
			log.Print(leftVal, rightVal)
			for i, val := range leftVal {
				if len(rightVal) < i + 1 {
					log.Print("left side had more packets")
					return false, false
				}
				inner, cont := evaluatePackets(val, rightVal[i])
				if !inner {
					log.Print("inside was false, whole expressions is false")
					return false, false
				} else {
					if !cont {
						log.Print("inside was true, no continue")
						return true, false
					}
				}
			}
			if len(leftVal) < len(rightVal) {
				// iterated through packets, all continue
				log.Print("left side ran out of comparators")
				return true, false
			} else {
				log.Print("all true, continue")
				return true, true
			}
		}
	} else {
		if leftType == reflect.Int {
			log.Print("left side was int")
			return evaluatePackets([]any{left}, right)
		} else {
			log.Print("right side was int")
			return evaluatePackets(left, []any{right})
		}
	}
}

func Day13(input io.Reader, part int) int {
	total := 0
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)

	index := 0
	packetPair := make([]any, 0)
	packetList := make([]any, 0)
	util.ProcessByLine(input, func(line string, num int) {
		if line != "" {
			left := regexp.MustCompile(`\[`)
			right := regexp.MustCompile(`\]`)
			line = right.ReplaceAllString(line, "}")
			line = left.ReplaceAllString(line, "[]any{")
			if packet, err := i.Eval("packet := " + line); err != nil {
				log.Fatal(err)
			} else {
				packetPair = append(packetPair, packet.Interface())
				packetList = append(packetList, packet.Interface())
			}
		}

		if len(packetPair) == 2 {
			// every third line is blank, compare packets
			index++
			result, cont := evaluatePackets(packetPair[0], packetPair[1])
			log.Printf("index %d, result %v, %v", index, result, cont)
			if result {
				total += index
			}
			packetPair = make([]any, 0)
		}
	})
	if part == 2 {
		packetList = append(packetList, []any{[]any{2}})
		packetList = append(packetList, []any{[]any{6}})
		sort.Sort(packetOrder(packetList))
		two, six := 0, 0
		for i, packet := range packetList {
			log.Print(packet)
			if fmt.Sprintf("%v", packet) == "[[2]]" {
				log.Print("found divider packet [[2]]")
				two = i + 1
			}
			if fmt.Sprintf("%v", packet) == "[[6]]" {
				log.Print("found divider packet [[6]]")
				six = i + 1
			}
		}
		total = two * six
	}
	log.Print(total)
	return total
}
