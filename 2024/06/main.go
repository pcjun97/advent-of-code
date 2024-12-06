package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

const (
	Empty int = iota
	Guard
	Obstacle
)

const (
	Up int = iota
	Right
	Down
	Left
)

func ParseGrid(s string) *aoc.Grid {
	g := aoc.NewGrid()

	for y, line := range strings.Split(s, "\n") {
		for x, r := range line {
			v := Empty
			switch r {
			case '^':
				v = Guard
			case '#':
				v = Obstacle
			}

			c := aoc.NewCoordinate(x, y)
			g.Add(aoc.NewNode(c, v))
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

func (s *Solver) Part1() int {
	return len(s.VisitedCoordinates())
}

func (s *Solver) VisitedCoordinates() []aoc.Coordinate {
	visited := make(map[aoc.Coordinate]struct{})

	c := s.GuardCoordinate()
	d := Up

	for {
		visited[c] = struct{}{}
		next := nextCoordinate(c, d)

		for s.Grid.Get(next) != nil && s.Grid.Get(next).Value() == Obstacle {
			d = (d + 1) % 4
			next = nextCoordinate(c, d)
		}

		if s.Grid.Get(next) == nil {
			break
		}

		c = next
	}

	slice := []aoc.Coordinate{}
	for k := range visited {
		slice = append(slice, k)
	}

	return slice
}

func (s *Solver) GuardCoordinate() aoc.Coordinate {
	for _, n := range s.Grid.Nodes() {
		if n.Value() == Guard {
			return n.Coordinate
		}
	}

	return aoc.Coordinate{}
}

func nextCoordinate(c aoc.Coordinate, d int) aoc.Coordinate {
	switch d {
	case Up:
		return aoc.NewCoordinate(c.X, c.Y-1)
	case Down:
		return aoc.NewCoordinate(c.X, c.Y+1)
	case Left:
		return aoc.NewCoordinate(c.X-1, c.Y)
	case Right:
		return aoc.NewCoordinate(c.X+1, c.Y)
	}

	return c
}

func (s *Solver) Part2() int {
	count := 0
	gc := s.GuardCoordinate()
	visited := s.VisitedCoordinates()

	for _, c := range visited {
		if c == gc {
			continue
		}

		s.Grid.Get(c).Set(Obstacle)
		if s.HasLoop() {
			count++
		}

		s.Grid.Get(c).Set(Empty)
	}

	return count
}

func (s *Solver) HasLoop() bool {
	type key struct {
		c aoc.Coordinate
		d int
	}

	visited := make(map[key]struct{})

	c := s.GuardCoordinate()
	d := Up

	for {
		if _, ok := visited[key{c, d}]; ok {
			return true
		}

		visited[key{c, d}] = struct{}{}
		next := nextCoordinate(c, d)

		for s.Grid.Get(next) != nil && s.Grid.Get(next).Value() == Obstacle {
			d = (d + 1) % 4
			next = nextCoordinate(c, d)
		}

		if s.Grid.Get(next) == nil {
			return false
		}

		c = next
	}
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
