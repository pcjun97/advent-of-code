package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

func ParsePattern(s string) *aoc.Grid {
	g := aoc.NewGrid()

	for y, line := range strings.Split(s, "\n") {
		for x, r := range line {
			c := aoc.NewCoordinate(x, y)
			g.Add(aoc.NewNode(c, int(r)))
		}
	}

	return g
}

func PatternHasOverlaps(g1, g2 *aoc.Grid) bool {
	for _, node := range g1.Nodes() {
		if node.Value() == '.' {
			continue
		}

		if node.Value() == g2.Get(node.Coordinate).Value() {
			return true
		}
	}

	return false
}

type Solver struct {
	Patterns []*aoc.Grid
}

func NewSolver(input string) *Solver {
	patterns := []*aoc.Grid{}

	for _, block := range strings.Split(input, "\n\n") {
		patterns = append(patterns, ParsePattern(block))
	}

	s := Solver{patterns}
	return &s
}

func (s *Solver) Part1() int {
	count := 0

	for i, p1 := range s.Patterns {
		for _, p2 := range s.Patterns[i+1:] {
			if !PatternHasOverlaps(p1, p2) {
				count++
			}
		}
	}

	return count
}

func (s *Solver) Part2() int {
	return 0
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
