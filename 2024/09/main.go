package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

func ParseGrid(s string) *aoc.Grid {
	g := aoc.NewGrid()

	for y, line := range strings.Split(s, "\n") {
		for x, r := range line {
			c := aoc.NewCoordinate(x, y)
			g.Add(aoc.NewNode(c, int(r-'0')))
		}
	}

	return g
}

type Solver struct {
	Grid *aoc.Grid
}

func NewSolver(input string) *Solver {
	s := Solver{ParseGrid(input)}
	return &s
}

func (s *Solver) TrailHeadCoordinates() []aoc.Coordinate {
	thcs := []aoc.Coordinate{}

	for _, node := range s.Grid.Nodes() {
		if node.Value() != 0 {
			continue
		}

		thcs = append(thcs, node.Coordinate)
	}

	return thcs
}

func (s *Solver) ReachableCoordinates(c aoc.Coordinate) []aoc.Coordinate {
	node := s.Grid.Get(c)
	if node == nil {
		return nil
	}

	rcs := []aoc.Coordinate{}
	for _, n := range s.Grid.Neighbors4Way(node) {
		if n.Value()-node.Value() == 1 {
			rcs = append(rcs, n.Coordinate)
		}
	}

	return rcs
}

func (s *Solver) Part1() int {
	cache := make(map[aoc.Coordinate]map[aoc.Coordinate]struct{})
	thcs := s.TrailHeadCoordinates()

	var trackHeads func(c aoc.Coordinate)
	trackHeads = func(c aoc.Coordinate) {
		if s.Grid.Get(c).Value() == 9 {
			cache[c][c] = struct{}{}
			return
		}

		next := s.ReachableCoordinates(c)

		for _, n := range next {
			if _, ok := cache[n]; !ok {
				cache[n] = make(map[aoc.Coordinate]struct{})
				trackHeads(n)
			}
			for nc := range cache[n] {
				cache[c][nc] = struct{}{}
			}
		}
	}

	sum := 0

	for _, head := range thcs {
		if _, ok := cache[head]; !ok {
			cache[head] = make(map[aoc.Coordinate]struct{})
		}

		trackHeads(head)
		sum += len(cache[head])
	}

	return sum
}

func (s *Solver) Part2() int {
	cache := make(map[aoc.Coordinate]int)
	thcs := s.TrailHeadCoordinates()

	var trackHeads func(c aoc.Coordinate)
	trackHeads = func(c aoc.Coordinate) {
		if s.Grid.Get(c).Value() == 9 {
			cache[c] = 1
			return
		}

		next := s.ReachableCoordinates(c)

		cache[c] = 0
		for _, n := range next {
			if _, ok := cache[n]; !ok {
				trackHeads(n)
			}
			cache[c] += cache[n]
		}
	}

	sum := 0

	for _, head := range thcs {
		trackHeads(head)
		sum += cache[head]
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
