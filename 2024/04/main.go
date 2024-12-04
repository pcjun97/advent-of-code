package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	Grid *aoc.Grid
}

func ParseGrid(s string) *aoc.Grid {
	g := aoc.NewGrid()

	for y, line := range strings.Split(s, "\n") {
		for x, r := range line {
			c := aoc.NewCoordinate(x, y)
			n := aoc.NewNode(c, int(r))
			g.Add(n)
		}
	}

	return g
}

func NewSolver(input string) *Solver {
	s := Solver{ParseGrid(input)}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0
	var p [4]int

	minX := s.Grid.MinX()
	maxX := s.Grid.MaxX()
	minY := s.Grid.MinY()
	maxY := s.Grid.MaxY()

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if y+3 <= maxY {
				for i := 0; i < 4; i++ {
					p[i] = s.Grid.Get(aoc.NewCoordinate(x, y+i)).Value()
				}

				if IsXMAS(p) {
					sum++
				}
			}

			if x+3 <= maxX {
				for i := 0; i < 4; i++ {
					p[i] = s.Grid.Get(aoc.NewCoordinate(x+i, y)).Value()
				}

				if IsXMAS(p) {
					sum++
				}
			}

			if x+3 <= maxX && y+3 <= maxY {
				for i := 0; i < 4; i++ {
					p[i] = s.Grid.Get(aoc.NewCoordinate(x+i, y+i)).Value()
				}

				if IsXMAS(p) {
					sum++
				}
			}

			if x-3 >= minX && y+3 <= maxY {
				for i := 0; i < 4; i++ {
					p[i] = s.Grid.Get(aoc.NewCoordinate(x-i, y+i)).Value()
				}

				if IsXMAS(p) {
					sum++
				}
			}
		}
	}

	return sum
}

func (s *Solver) Part2() int {
	sum := 0
	var p [2]int

	minX := s.Grid.MinX()
	maxX := s.Grid.MaxX()
	minY := s.Grid.MinY()
	maxY := s.Grid.MaxY()

	for x := minX + 1; x < maxX; x++ {
		for y := minY + 1; y < maxY; y++ {
			c := aoc.NewCoordinate(x, y)

			if s.Grid.Get(c).Value() != 'A' {
				continue
			}

			p[0] = s.Grid.Get(aoc.NewCoordinate(x-1, y-1)).Value()
			p[1] = s.Grid.Get(aoc.NewCoordinate(x+1, y+1)).Value()

			if !(p[0] == 'M' && p[1] == 'S') && !(p[0] == 'S' && p[1] == 'M') {
				continue
			}

			p[0] = s.Grid.Get(aoc.NewCoordinate(x+1, y-1)).Value()
			p[1] = s.Grid.Get(aoc.NewCoordinate(x-1, y+1)).Value()

			if !(p[0] == 'M' && p[1] == 'S') && !(p[0] == 'S' && p[1] == 'M') {
				continue
			}

			sum++
		}
	}

	return sum
}

func IsXMAS(p [4]int) bool {
	if p[0] == 'X' && p[1] == 'M' && p[2] == 'A' && p[3] == 'S' {
		return true
	}

	if p[3] == 'X' && p[2] == 'M' && p[1] == 'A' && p[0] == 'S' {
		return true
	}

	return false
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
