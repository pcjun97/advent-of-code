package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

const (
	GardenPlotTile int = iota
	RockTile
)

func parseGrid(s string) *aoc.Grid {
	grid := aoc.NewGrid()

	for y, line := range strings.Split(s, "\n") {
		for x, r := range []byte(line) {
			c := aoc.NewCoordinate(x, y)
			tile := GardenPlotTile
			if r == '#' {
				tile = RockTile
			}
			grid.Add(aoc.NewNode(c, tile))
		}
	}

	return grid
}

func startingPosition(s string) aoc.Coordinate {
	for y, line := range strings.Split(s, "\n") {
		for x, r := range []byte(line) {
			if r == 'S' {
				return aoc.NewCoordinate(x, y)
			}
		}
	}

	return aoc.NewCoordinate(-1, -1)
}

type Solver struct {
	startingPosition aoc.Coordinate
	grid             *aoc.Grid
}

func NewSolver(input string) *Solver {
	s := Solver{startingPosition(input), parseGrid(input)}
	return &s
}

func (s *Solver) Part1(steps int) int {
	current := make(map[aoc.Coordinate]struct{})
	current[s.startingPosition] = struct{}{}

	for i := 0; i < steps; i++ {
		next := make(map[aoc.Coordinate]struct{})

		for c := range current {
			for _, n := range s.grid.Neighbors4Way(s.grid.Get(c)) {
				if n.Value() == GardenPlotTile {
					next[n.Coordinate] = struct{}{}
				}
			}
		}

		current = next
	}

	return len(current)
}

func (s *Solver) Part2(steps int) int {
	width := s.grid.MaxX() + 1
	modulo := steps % width

	cache := [2]int{0, 0}
	visited := make(map[aoc.Coordinate]struct{})

	current := make(map[aoc.Coordinate]struct{})
	current[s.startingPosition] = struct{}{}

	list := []int{}

	i := 0
	for len(list) < 4 {
		i += 1

		next := make(map[aoc.Coordinate]struct{})
		for c := range current {
			for _, cc := range c.Neighbors4Way() {
				n := cc

				n.X %= width
				if n.X < 0 {
					n.X += width
				}

				n.Y %= width
				if n.Y < 0 {
					n.Y += width
				}

				if _, ok := visited[cc]; ok {
					continue
				}
				visited[cc] = struct{}{}

				if s.grid.Get(n).Value() == GardenPlotTile {
					next[cc] = struct{}{}
				}
			}
		}
		current = next
		cache[i%2] += len(current)

		if i == steps {
			return cache[i%2]
		}

		if (i % width) == modulo {
			list = append(list, cache[i%2])
		}
		if len(list) == 4 && !IsQuadratic(list) {
			list = list[1:]
		}
	}

	for i < steps {
		list = append(list, aoc.Extrapolate(list[len(list)-3:]))
		i += width
	}

	return list[len(list)-1]
}

func IsQuadratic(list []int) bool {
	d1 := []int{}
	for i := 1; i < len(list); i++ {
		d1 = append(d1, list[i]-list[i-1])
	}

	d2 := []int{}
	for i := 1; i < len(d1); i++ {
		d2 = append(d2, d1[i]-d1[i-1])
	}

	for i := 1; i < len(d2); i++ {
		if d2[i] != d2[i-1] {
			return false
		}
	}
	return true
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(64), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(26501365), time.Since(start).String())
}
