package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	steps []string
}

func hash(key string) int {
	result := 0
	for _, r := range []byte(key) {
		result += int(r)
		result = (result * 17) % 256
	}
	return result
}

func NewSolver(input string) *Solver {
	steps := strings.Split(input, ",")
	s := Solver{steps}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0
	for _, step := range s.steps {
		sum += hash(step)
	}
	return sum
}

func (s *Solver) Part2() int {
	boxes := [256][]Lens{}
	r := regexp.MustCompile(`(.+)(=|-)(.*)`)
	for _, step := range s.steps {
		m := r.FindStringSubmatch(step)
		label := m[1]
		i := hash(label)

		switch m[2] {
		case "-":
			next := []Lens{}
			for _, l := range boxes[i] {
				if l.label != label {
					next = append(next, l)
				}
			}
			boxes[i] = next
		case "=":
			value, _ := strconv.Atoi(m[3])
			found := false
			for j, l := range boxes[i] {
				if l.label == label {
					found = true
					boxes[i][j].value = value
				}
			}

			if !found {
				boxes[i] = append(boxes[i], Lens{label, value})
			}
		}
	}

	sum := 0
	for i, box := range boxes {
		for j, l := range box {
			sum += (i + 1) * (j + 1) * l.value
		}
	}

	return sum
}

type Lens struct {
	label string
	value int
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
