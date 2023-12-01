package main

import (
	"fmt"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
}

func NewSolver(input string) *Solver {
	s := Solver{}
	return &s
}

func (s *Solver) Part1() int {
	return 0
}

func (s *Solver) Part2() int {
	return 0
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
