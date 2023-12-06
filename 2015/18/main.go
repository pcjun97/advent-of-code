package main

import (
	"fmt"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func parseGrid(s string) *aoc.Grid {
	g := aoc.NewGrid()
	for y, line := range strings.Split(s, "\n") {
		for x, v := range []byte(line) {
			c := aoc.NewCoordinate(x, y)
			state := 0
			if v == '#' {
				state = 1
			}
			n := aoc.NewNode(c, state)
			g.Add(n)
		}
	}
	return g
}

func nextGrid(g *aoc.Grid) *aoc.Grid {
	next := aoc.NewGrid()
	for _, n := range g.Nodes() {
		neighbors := g.Neighbors8Way(n)
		count := 0
		for _, neighbor := range neighbors {
			count += neighbor.Value()
		}

		v := n.Value()
		switch {
		case n.Value() == 1 && (count < 2 || count > 3):
			v = 0
		case n.Value() == 0 && count == 3:
			v = 1
		}

		nextNode := aoc.NewNode(n.Coordinate, v)
		next.Add(nextNode)
	}
	return next
}

type Solver struct {
	grid *aoc.Grid
}

func NewSolver(input string) *Solver {
	grid := parseGrid(input)
	s := Solver{grid}
	return &s
}

func (s *Solver) Part1(steps int) int {
	grid := s.grid
	for i := 0; i < steps; i++ {
		grid = nextGrid(grid)
	}

	count := 0
	for _, n := range grid.Nodes() {
		count += n.Value()
	}
	return count
}

func (s *Solver) Part2(steps int) int {
	grid := s.grid
	maxX, maxY := 0, 0
	for _, n := range grid.Nodes() {
		if n.Coordinate.X > maxX {
			maxX = n.Coordinate.X
		}
		if n.Coordinate.Y > maxY {
			maxY = n.Coordinate.Y
		}
	}

	TurnOnCorners(grid, maxX, maxY)
	for i := 0; i < steps; i++ {
		grid = nextGrid(grid)
		TurnOnCorners(grid, maxX, maxY)
	}

	count := 0
	for _, n := range grid.Nodes() {
		count += n.Value()
	}
	return count
}

func TurnOnCorners(g *aoc.Grid, maxX, maxY int) {
	g.Get(aoc.NewCoordinate(0, 0)).Set(1)
	g.Get(aoc.NewCoordinate(maxX, 0)).Set(1)
	g.Get(aoc.NewCoordinate(0, maxY)).Set(1)
	g.Get(aoc.NewCoordinate(maxX, maxY)).Set(1)
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1(100))
	fmt.Println(s.Part2(100))
}
