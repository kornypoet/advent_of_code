package calendar

import (
	"log"
	"regexp"
	"strconv"

	"github.com/kornypoet/advent_of_code/util"
)

/*
	[C] [B] [H]
[W]     [D] [J] [Q] [B]
[P] [F] [Z] [F] [B] [L]
[G] [Z] [N] [P] [J] [S] [V]
[Z] [C] [H] [Z] [G] [T] [Z]     [C]
[V] [B] [M] [M] [C] [Q] [C] [G] [H]
[S] [V] [L] [D] [F] [F] [G] [L] [F]
[B] [J] [V] [L] [V] [G] [L] [N] [J]
 1   2   3   4   5   6   7   8   9
*/

var stacks = map[string][]string{
	"stack1": []string{"B", "S", "V", "Z", "G", "P", "W"},
	"stack2": []string{"J", "V", "B", "C", "Z", "F"},
	"stack3": []string{"V", "L", "M", "H", "N", "Z", "D", "C"},
	"stack4": []string{"L", "D", "M", "Z", "P", "F", "J", "B"},
	"stack5": []string{"V", "F", "C", "G", "J", "B", "Q", "H"},
	"stack6": []string{"G", "F", "Q", "T", "S", "L", "B"},
	"stack7": []string{"L", "G", "C", "Z", "V"},
	"stack8": []string{"N", "L", "G"},
	"stack9": []string{"J", "F", "H", "C"},
}

func pop(stack *[]string) (val string) {
	val, *stack = (*stack)[len(*stack)-1], (*stack)[:len(*stack)-1]
	return
}

func cut(stack *[]string, amount int) (val []string) {
	val = (*stack)[len(*stack)-amount:]
	*stack = (*stack)[:len(*stack)-amount]
	return
}

func Day5() {
	util.ProcessByLine("input/day5.txt", func(line string, num int) {
		r := regexp.MustCompile(`move (?P<Count>\d+) from (?P<Source>\d+) to (?P<Dest>\d+)`)
		directions := r.FindStringSubmatch(line)
		if len(directions) != 0 {
			count, _ := strconv.Atoi(directions[1])
			source := directions[2]
			dest := directions[3]
			src := stacks["stack"+source]
			dst := stacks["stack"+dest]
			log.Printf("moving from %v to %v %d times", src, dst, count)
			// for i := 1; i <= count; i++ {
			// 	val := pop(&src)
			// 	dst = append(dst, val)
			// }
			val := cut(&src, count)
			dst = append(dst, val...)
			stacks["stack"+source] = src
			stacks["stack"+dest] = dst
			log.Print(stacks)
		}
	})

	code := ""
	for i := '1'; i <= '9'; i++ {
		stack := stacks["stack"+string(i)]
		code += stack[len(stack)-1]
	}

	log.Printf("the final code: %s", code)
}
