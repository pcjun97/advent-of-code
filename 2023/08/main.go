package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Node struct {
	label       string
	left, right string
}

func parseNode(s string) Node {
	r := regexp.MustCompile(`(.*) = \((.*), (.*)\)`)
	m := r.FindStringSubmatch(s)
	n := Node{m[1], m[2], m[3]}
	return n
}

type Solver struct {
	nodes       map[string]Node
	instruction string
}

func NewSolver(input string) *Solver {
	lines := strings.Split(input, "\n")
	nodes := make(map[string]Node)
	for _, line := range lines[2:] {
		n := parseNode(line)
		nodes[n.label] = n
	}
	s := Solver{nodes, lines[0]}
	return &s
}

func (s *Solver) Part1() int {
	return s.Steps("AAA", regexp.MustCompile(`ZZZ`))
}

func (s *Solver) Part2() int {
	to := regexp.MustCompile(`..Z`)
	steps := []int{}
	for from := range s.nodes {
		if from[2] != 'A' {
			continue
		}
		step := s.Steps(from, to)
		steps = append(steps, step)
	}
	return LCM(steps)
}

func (s *Solver) Steps(from string, to *regexp.Regexp) int {
	count := 0
	current := s.nodes[from]
	for i := 0; !to.MatchString(current.label); i = (i + 1) % len(s.instruction) {
		switch s.instruction[i] {
		case 'L':
			current = s.nodes[current.left]
		case 'R':
			current = s.nodes[current.right]
		}
		count++
	}
	return count
}

func LCM(numbers []int) int {
	result := numbers[0]
	for _, n := range numbers[1:] {
		result = (result * n) / GCD(result, n)
	}

	return result
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
