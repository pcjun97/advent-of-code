package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	Towels   []string
	Patterns []string
}

func NewSolver(input string) *Solver {
	blocks := strings.Split(input, "\n\n")
	towels := strings.Split(blocks[0], ", ")
	patterns := strings.Split(blocks[1], "\n")
	s := Solver{towels, patterns}
	return &s
}

func (s *Solver) CountPossibleArrangements(pattern string) int {
	cache := make(map[string]int)

	var count func(p string) int
	count = func(p string) int {
		if n, ok := cache[p]; ok {
			return n
		}

		c := 0
		for _, towel := range s.Towels {
			if towel == p {
				c++
				continue
			}
			if strings.HasPrefix(p, towel) {
				c += count(p[len(towel):])
			}
		}

		cache[p] = c
		return c
	}

	return count(pattern)
}

func (s *Solver) Part1() int {
	sum := 0

	for _, pattern := range s.Patterns {
		if s.CountPossibleArrangements(pattern) > 0 {
			sum++
		}
	}

	return sum
}

func (s *Solver) Part2() int {
	sum := 0
	for _, pattern := range s.Patterns {
		sum += s.CountPossibleArrangements(pattern)
	}
	return sum
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
