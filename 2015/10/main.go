package main

import (
	"fmt"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	initial string
}

func NewSolver(input string) *Solver {
	s := Solver{input}
	return &s
}

func (s *Solver) Part1() int {
	result := s.initial
	for i := 0; i < 40; i++ {
		result = LookAndSay(result)
	}
	return len(result)
}

func (s *Solver) Part2() int {
	result := s.initial
	for i := 0; i < 50; i++ {
		result = LookAndSay(result)
	}
	return len(result)
}

func LookAndSay(s string) string {
	var result strings.Builder

	count := 0
	for j := range s {
		count += 1
		if j == len(s)-1 || s[j+1] != s[j] {
			fmt.Fprintf(&result, "%d%c", count, s[j])
			count = 0
			continue
		}
	}

	return result.String()
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
