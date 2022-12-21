package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()
	m := NewMonkeys(input)
	fmt.Println(m.Value("root"))
	fmt.Println(m.Balance("root", "humn"))
}

type Monkey struct {
	name string
	job  string
}

func (m Monkey) Dependencies() []string {
	fields := strings.Split(m.job, " ")
	if len(fields) == 1 {
		return []string{}
	}
	return []string{fields[0], fields[2]}
}

type Monkeys struct {
	monkeys map[string]Monkey
	cache   map[string]int
}

func NewMonkeys(s string) Monkeys {
	monkeys := make(map[string]Monkey)
	for _, line := range strings.Split(s, "\n") {
		fields := strings.Split(line, ": ")
		monkey := Monkey{fields[0], fields[1]}
		monkeys[fields[0]] = monkey
	}
	m := Monkeys{
		monkeys: monkeys,
		cache:   make(map[string]int),
	}
	return m
}

func (m Monkeys) Value(name string) int {
	if val, ok := m.cache[name]; ok {
		return val
	}

	monkey, ok := m.monkeys[name]
	if !ok {
		panic("monkey not found: " + name)
	}

	fields := strings.Split(monkey.job, " ")
	if len(fields) == 1 {
		val, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}
		m.cache[name] = val
		return val
	}

	m1 := m.Value(fields[0])
	m2 := m.Value(fields[2])
	switch fields[1] {
	case "+":
		return m1 + m2
	case "-":
		return m1 - m2
	case "*":
		return m1 * m2
	case "/":
		return m1 / m2
	default:
		panic("Unknown operator: " + fields[1])
	}
}

func (m Monkeys) HasDependency(root string, dependency string) bool {
	monkey, ok := m.monkeys[root]
	if !ok {
		return false
	}

	if _, ok := m.monkeys[dependency]; !ok {
		return false
	}

	for _, d := range monkey.Dependencies() {
		if d == dependency {
			return true
		}
		if m.HasDependency(d, dependency) {
			return true
		}
	}

	return false
}

func (m Monkeys) Balance(source string, target string) int {
	monkey, ok := m.monkeys[source]
	if !ok {
		panic("monkey not found: " + source)
	}

	if _, ok := m.monkeys[target]; !ok {
		panic("monkey not found: " + target)
	}

	if !m.HasDependency(source, target) {
		panic("source does not depend on target")
	}

	var val, targetVal int
	var dependant string

	fields := strings.Split(monkey.job, " ")
	if m.HasDependency(fields[0], target) {
		dependant = fields[0]
		targetVal = m.Value(fields[2])
	} else {
		dependant = fields[2]
		targetVal = m.Value(fields[0])
	}

	for dependant != target {
		monkey = m.monkeys[dependant]
		fields = strings.Split(monkey.job, " ")

		if m.HasDependency(fields[0], target) || fields[0] == target {
			dependant = fields[0]
			val = m.Value(fields[2])
		} else {
			dependant = fields[2]
			val = m.Value(fields[0])
		}

		switch fields[1] {
		case "+":
			targetVal -= val
		case "*":
			targetVal /= val
		case "-":
			if dependant == fields[0] {
				targetVal += val
			} else {
				targetVal = val - targetVal
			}
		case "/":
			if dependant == fields[0] {
				targetVal *= val
			} else {
				targetVal = targetVal / val
			}
		default:
			panic("Unknown operator: " + fields[1])
		}
	}

	return targetVal
}
