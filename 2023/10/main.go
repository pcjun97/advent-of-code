package main

import (
	"fmt"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Grid map[aoc.Coordinate]byte

func parseGrid(s string) Grid {
	g := make(Grid)

	for y, line := range strings.Split(s, "\n") {
		for x, r := range []byte(line) {
			g[aoc.NewCoordinate(x, y)] = r
		}
	}

	return g
}

func (g Grid) Destination(c aoc.Coordinate, from aoc.Coordinate) aoc.Coordinate {
	dx := from.X - c.X
	dy := from.Y - c.Y

	switch {
	case dx == 1 && dy == 0:
		if g[c] == '-' {
			return aoc.NewCoordinate(c.X-1, c.Y)
		}
		if g[c] == 'L' {
			return aoc.NewCoordinate(c.X, c.Y-1)
		}
		if g[c] == 'F' {
			return aoc.NewCoordinate(c.X, c.Y+1)
		}
	case dx == -1 && dy == 0:
		if g[c] == '-' {
			return aoc.NewCoordinate(c.X+1, c.Y)
		}
		if g[c] == 'J' {
			return aoc.NewCoordinate(c.X, c.Y-1)
		}
		if g[c] == '7' {
			return aoc.NewCoordinate(c.X, c.Y+1)
		}
	case dy == 1 && dx == 0:
		if g[c] == '|' {
			return aoc.NewCoordinate(c.X, c.Y-1)
		}
		if g[c] == '7' {
			return aoc.NewCoordinate(c.X-1, c.Y)
		}
		if g[c] == 'F' {
			return aoc.NewCoordinate(c.X+1, c.Y)
		}
	case dy == -1 && dx == 0:
		if g[c] == '|' {
			return aoc.NewCoordinate(c.X, c.Y+1)
		}
		if g[c] == 'J' {
			return aoc.NewCoordinate(c.X-1, c.Y)
		}
		if g[c] == 'L' {
			return aoc.NewCoordinate(c.X+1, c.Y)
		}
	}

	return aoc.NewCoordinate(-1, -1)
}

func (g Grid) FindStartingLoop() []aoc.Coordinate {
	start := g.StartingPoint()
	candidates := []aoc.Coordinate{
		aoc.NewCoordinate(start.X-1, start.Y),
		aoc.NewCoordinate(start.X+1, start.Y),
		aoc.NewCoordinate(start.X, start.Y-1),
	}
	for _, c := range candidates {
		loop := []aoc.Coordinate{start}
		next := c
		for next.X >= 0 && next.Y >= 0 {
			if next == start {
				return loop
			}
			loop = append(loop, next)
			next = g.Destination(loop[len(loop)-1], loop[len(loop)-2])
		}
	}

	return nil
}

func (g Grid) StartingPoint() aoc.Coordinate {
	for c, v := range g {
		if v == 'S' {
			return c
		}
	}
	return aoc.NewCoordinate(-1, -1)
}

type Solver struct {
	grid Grid
}

func NewSolver(input string) *Solver {
	grid := parseGrid(input)
	s := Solver{grid}
	return &s
}

func (s *Solver) Part1() int {
	return len(s.grid.FindStartingLoop()) / 2
}

func (s *Solver) Part2() int {
	loop := s.grid.FindStartingLoop()
	loopMap := make(map[aoc.Coordinate]byte)
	for _, c := range loop {
		loopMap[c] = s.grid[c]
	}

	start := loop[0]
	lastLoop := loop[len(loop)-1]
	up := aoc.NewCoordinate(start.X, start.Y-1)
	down := aoc.NewCoordinate(start.X, start.Y+1)
	left := aoc.NewCoordinate(start.X-1, start.Y)
	right := aoc.NewCoordinate(start.X+1, start.Y)

	switch {
	case (loop[1] == up && lastLoop == down) || (loop[1] == down && lastLoop == up):
		loopMap[loop[0]] = '|'
	case (loop[1] == left && lastLoop == right) || (loop[1] == right && lastLoop == left):
		loopMap[loop[0]] = '-'
	case (loop[1] == up && lastLoop == left) || (loop[1] == left && lastLoop == up):
		loopMap[loop[0]] = 'L'
	case (loop[1] == up && lastLoop == right) || (loop[1] == right && lastLoop == up):
		loopMap[loop[0]] = 'J'
	case (loop[1] == down && lastLoop == left) || (loop[1] == left && lastLoop == down):
		loopMap[loop[0]] = '&'
	case (loop[1] == down && lastLoop == right) || (loop[1] == right && lastLoop == down):
		loopMap[loop[0]] = 'F'
	}

	maxX, maxY := 0, 0
	for c := range s.grid {
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}

	count := 0
	var last byte
	for y := 0; y <= maxY; y++ {
		inside := false
		for x := 0; x <= maxX; x++ {
			c := aoc.NewCoordinate(x, y)
			if r, ok := loopMap[c]; ok {
				switch {
				case r == '7':
					if last == 'F' {
						inside = !inside
					}
				case r == 'J':
					if last == 'L' {
						inside = !inside
					}
				case r != '-':
					inside = !inside
				}
				if r != '-' {
					last = r
				}
				continue
			}
			if inside {
				count += 1
			}
		}
	}

	return count
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
