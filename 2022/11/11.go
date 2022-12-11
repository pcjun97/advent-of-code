package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()

	m1 := NewMonkeys(input, true)
	for i := 0; i < 20; i++ {
		m1.StartRound()
	}
	fmt.Println(m1.MonkeyBusinessLevel())

	m2 := NewMonkeys(input, false)
	for i := 0; i < 10000; i++ {
		m2.StartRound()
	}
	fmt.Println(m2.MonkeyBusinessLevel())
}

type Monkey struct {
	Items     []int
	Operation []string
	TestVal   int
	Targets   [2]*Monkey
	Inspected int
}

func NewMonkey(s string) *Monkey {
	m := Monkey{
		Inspected: 0,
	}

	return &m
}

func (m *Monkey) InspectItem(i int) {
	item := m.Items[i]
	var a, b int
	var err error

	if m.Operation[0] == "old" {
		a = item
	} else {
		a, err = strconv.Atoi(m.Operation[0])
		if err != nil {
			panic("error parsing operation: " + m.Operation[0])
		}
	}

	if m.Operation[2] == "old" {
		b = item
	} else {
		b, err = strconv.Atoi(m.Operation[2])
		if err != nil {
			panic("error parsing operation: " + m.Operation[2])
		}
	}

	switch m.Operation[1] {
	case "+":
		item = a + b
	case "*":
		item = a * b
	}

	m.Items[i] = item
	m.Inspected++
}

func (m *Monkey) ThrowItem(i int) {
	item := m.Items[i]
	if item%m.TestVal == 0 {
		m.Targets[0].Items = append(m.Targets[0].Items, item)
	} else {
		m.Targets[1].Items = append(m.Targets[1].Items, item)
	}
	m.Items = append(m.Items[:i], m.Items[i+1:]...)
}

type Monkeys struct {
	monkeys       []*Monkey
	cm            int
	decreaseWorry bool
}

func NewMonkeys(s string, decreaseWorry bool) *Monkeys {
	sections := strings.Split(s, "\n\n")

	monkeys := make([]*Monkey, len(sections))
	for i, section := range sections {
		monkeys[i] = NewMonkey(section)
	}

	divisors := make(map[int]bool)
	for i, section := range sections {
		monkey := monkeys[i]

		lines := make([]string, 6)
		for i, line := range strings.Split(section, "\n")[1:] {
			colon := strings.Index(line, ":")
			lines[i] = strings.TrimSpace(line[colon+1:])
		}

		monkey.Items = []int{}
		for _, it := range strings.Split(lines[0], ", ") {
			item, err := strconv.Atoi(it)
			if err != nil {
				panic(err)
			}
			monkey.Items = append(monkey.Items, item)
		}

		monkey.Operation = strings.Split(lines[1], " ")[2:]

		testVal, err := strconv.Atoi(strings.Split(lines[2], " ")[2])
		if err != nil {
			panic(err)
		}
		monkey.TestVal = testVal
		divisors[testVal] = true

		t1, err := strconv.Atoi(strings.Split(lines[3], " ")[3])
		if err != nil {
			panic(err)
		}
		monkey.Targets[0] = monkeys[t1]

		t2, err := strconv.Atoi(strings.Split(lines[4], " ")[3])
		if err != nil {
			panic(err)
		}
		monkey.Targets[1] = monkeys[t2]
	}

	cm := 1
	for d := range divisors {
		cm *= d
	}

	m := Monkeys{
		monkeys:       monkeys,
		cm:            cm,
		decreaseWorry: decreaseWorry,
	}

	return &m
}

func (m *Monkeys) StartRound() {
	for _, monkey := range m.monkeys {
		for len(monkey.Items) > 0 {
			monkey.InspectItem(0)
			monkey.Items[0] = monkey.Items[0] % m.cm
			if m.decreaseWorry {
				monkey.Items[0] = monkey.Items[0] / 3
			}
			monkey.ThrowItem(0)
		}
	}
}

func (m *Monkeys) MonkeyBusinessLevel() int {
	tmp := make([]*Monkey, len(m.monkeys))
	copy(tmp, m.monkeys)
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].Inspected > tmp[j].Inspected
	})
	return tmp[0].Inspected * tmp[1].Inspected
}
