package calendar

import (
	"fmt"
	"io"
	"log"

	//"github.com/kornypoet/advent_of_code/util"
)

type Monkey struct {
	items []*Item
	operation func(int)int
	test func(int)int
	inspected int
}

type Item struct {
	worry int
}

func (i *Item) String() string {
	return fmt.Sprintf("%d", i.worry)
}

func (m *Monkey) handle() (int, *Item) {
	item := m.items[0]
	log.Printf("Handling item %s", item)
	m.items = m.items[1:]
	log.Printf("Monkey is now holding %s", m.items)
	inspect := m.operation(item.worry)
	log.Printf("Monkey inspection caused worry %d", inspect)
	m.inspected++
	bored := inspect % (11 * 5 * 7 * 2 * 17 * 13 * 3 * 19)
	// log.Printf("Monkey got bored %d", bored)
	nextMonkey := m.test(bored)
	// nextMonkey := m.test(inspect)
	item.worry = bored
	// item.worry = inspect
	log.Printf("Throwing %s to monkey %d", item, nextMonkey)
	return nextMonkey, item
}

func (m *Monkey) catch(item *Item) {
	m.items = append(m.items, item)
}

func (m *Monkey) String() string {
	return fmt.Sprintf("%s", m.items)
}


func TestMonkeys() map[int]*Monkey {
	return map[int]*Monkey{
		0: &Monkey{
			items: []*Item{ &Item{79}, &Item{98} },
			operation: func(old int)(int){ return (old * 19) },
			test: func(worry int)(int){
				if worry % 23 == 0 {
					return 2
				} else {
					return 3
				}
			},
		},
		1: &Monkey{
			items: []*Item{&Item{54}, &Item{65}, &Item{75}, &Item{74}},
			operation: func(old int)(int){ return (old + 6)},
			test: func(worry int)(int){
				if worry % 19 == 0 {
					return 2
				} else {
					return 0
				}
			},
		},
		2: &Monkey{
			items: []*Item{&Item{79}, &Item{60}, &Item{97}},
			operation: func(old int)(int){ return (old * old) },
			test: func(worry int)(int){
				if worry % 13 == 0 {
					return 1
				} else {
					return 3
				}
			},
		},
		3: &Monkey{
			items: []*Item{&Item{74}},
			operation: func(old int)(int){ return (old + 3) },
			test: func(worry int)(int){
				if worry % 17 == 0 {
					return 0
				} else {
					return 1
				}
			},
		},
	}
}

func RealMonkeys() map[int]*Monkey {
	return map[int]*Monkey{
		0: &Monkey{
			items: []*Item{&Item{54}, &Item{82}, &Item{90}, &Item{88}, &Item{86}, &Item{54}},
			operation: func(old int)(int){ return (old * 7) },
			test: func(worry int)(int){
				if worry % 11 == 0 {
					return 2
				} else {
					return 6
				}
			},
		},
		1: &Monkey{
			items:[]*Item{ &Item{91}, &Item{65}},
			operation: func(old int)(int){ return (old * 13) },
			test: func(worry int)(int){
				if worry % 5 == 0 {
					return 7
				} else {
					return 4
				}
			},
		},
		2: &Monkey{
			items:[]*Item{ &Item{62}, &Item{54}, &Item{57}, &Item{92}, &Item{83}, &Item{63}, &Item{63}},
			operation: func(old int)(int){ return (old + 1) },
			test: func(worry int)(int){
				if worry % 7 == 0 {
					return 1
				} else {
					return 7
				}
			},
		},
		3: &Monkey{
			items: []*Item{ &Item{67}, &Item{72}, &Item{68}},
			operation: func(old int)(int){ return (old * old) },
			test: func(worry int)(int){
				if worry % 2 == 0 {
					return 0
				} else {
					return 6
				}
			},
		},
		4: &Monkey{
			items: []*Item{ &Item{68}, &Item{89}, &Item{90}, &Item{86}, &Item{84}, &Item{57}, &Item{72}, &Item{84}},
			operation: func(old int)(int){ return (old + 7) },
			test: func(worry int)(int){
				if worry % 17 == 0 {
					return 3
				} else {
					return 5
				}
			},
		},
		5: &Monkey{
			items: []*Item{ &Item{79}, &Item{83}, &Item{64}, &Item{58} },
			operation: func(old int)(int){ return (old + 6) },
			test: func(worry int)(int){
				if worry % 13 == 0 {
					return 3
				} else {
					return 0
				}
			},
		},
		6: &Monkey{
			items: []*Item{ &Item{96}, &Item{72}, &Item{89}, &Item{70}, &Item{88} },
			operation: func(old int)(int){ return (old + 4) },
			test: func(worry int)(int){
				if worry % 3 == 0 {
					return 1
				} else {
					return 2
				}
			},
		},
		7: &Monkey{
			items: []*Item{ &Item{79} },
			operation: func(old int)(int){ return (old + 8) },
			test: func(worry int)(int){
				if worry % 19 == 0 {
					return 4
				} else {
					return 5
				}
			},
		},
	}
}

func Day11(input io.Reader, part int) int {
	var total int
	if part == 1 {
		total = 1
	} else {
		total = 2
	}

	// util.ProcessByLine(input, func(line string, num int) {
	//	log.Printf("%s, %d", line, num)
	// })

	// monkeys := TestMonkeys()
	monkeys := RealMonkeys()
	for round := 1; round <= 10000; round++ {
		for i := 0; i < len(monkeys); i++ {
			log.Printf("Monkey %d", i)
			monkey := monkeys[i]
			for range monkey.items {
				next, item := monkey.handle()
				monkeys[next].catch(item)
			}
		}
	}

	// log.Printf("%s", monkeys)
	for _, monkey := range monkeys {
		log.Printf("Monkey inspected %d", monkey.inspected)
	}
	return total
}
