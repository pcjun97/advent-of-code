package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

const (
	PathTile       int = '.'
	ForestTile     int = '#'
	SlopeUpTile    int = '^'
	SlopeDownTile  int = 'v'
	SlopeLeftTile  int = '<'
	SlopeRightTile int = '>'
)

func parseGrid(s string) *aoc.Grid {
	grid := aoc.NewGrid()

	for y, line := range strings.Split(s, "\n") {
		for x, r := range []byte(line) {
			c := aoc.NewCoordinate(x, y)
			grid.Add(aoc.NewNode(c, int(r)))
		}
	}

	return grid
}

type Solver struct {
	grid *aoc.Grid
}

func NewSolver(input string) *Solver {
	s := Solver{parseGrid(input)}
	return &s
}

func (s *Solver) Part1() int {
	graph := s.DirectedGraph()

	start := s.StartCoordinate()
	end := s.EndCoordinate()

	ref := make(map[aoc.Coordinate]map[aoc.Coordinate]struct{})
	for c, m := range graph {
		for cc := range m {
			if _, ok := ref[cc]; !ok {
				ref[cc] = make(map[aoc.Coordinate]struct{})
			}
			ref[cc][c] = struct{}{}
		}
	}

	distances := make(map[aoc.Coordinate]int)
	var maxd func(c aoc.Coordinate) int
	maxd = func(c aoc.Coordinate) int {
		if c == start {
			return 0
		}

		if d, ok := distances[c]; ok {
			return d
		}

		max := math.MinInt
		for from := range ref[c] {
			d := maxd(from) + graph[from][c]
			if d > max {
				max = d
			}
		}
		distances[c] = max
		return max
	}

	return maxd(end)
}

func (s *Solver) Part2() int {
	dgraph := s.DirectedGraph()
	graph := make(map[aoc.Coordinate]map[aoc.Coordinate]int)
	for from, m := range dgraph {
		for to, d := range m {
			if _, ok := graph[from]; !ok {
				graph[from] = make(map[aoc.Coordinate]int)
			}
			graph[from][to] = d

			if _, ok := graph[to]; !ok {
				graph[to] = make(map[aoc.Coordinate]int)
			}
			graph[to][from] = d
		}
	}

	i := 0
	im := make(map[aoc.Coordinate]int)
	for c := range graph {
		im[c] = i
		i += 1
	}

	hasvisit := func(visited int, c aoc.Coordinate) bool {
		mask := 1 << im[c]
		return (visited & mask) != 0
	}

	addvisit := func(visited int, c aoc.Coordinate) int {
		mask := 1 << im[c]
		return (visited | mask)
	}

	max := 0
	end := s.EndCoordinate()

	var visit func(c aoc.Coordinate, visited int, travelled int)
	visit = func(c aoc.Coordinate, visited int, travelled int) {
		if c == end {
			if travelled > max {
				max = travelled
			}
			return
		}

		for n, d := range graph[c] {
			if hasvisit(visited, n) {
				continue
			}

			visit(n, addvisit(visited, n), travelled+d)
		}
	}
	visit(s.StartCoordinate(), 0, 0)

	return max
}

func (s *Solver) DirectedGraph() map[aoc.Coordinate]map[aoc.Coordinate]int {
	graph := make(map[aoc.Coordinate]map[aoc.Coordinate]int)

	start := s.StartCoordinate()

	var walk func(c, prev, from aoc.Coordinate, count int)
	walk = func(c, prev, from aoc.Coordinate, count int) {
		next := s.NextCoordinates(c, prev)

		if len(next) == 1 && !s.IsFork(c) {
			walk(next[0], c, from, count+1)
			return
		}

		if _, ok := graph[from]; !ok {
			graph[from] = make(map[aoc.Coordinate]int)
		}
		graph[from][c] = count

		if len(next) == 0 {
			return
		}

		for _, cc := range next {
			walk(cc, c, c, 1)
		}
	}
	walk(start, start, start, 0)

	return graph
}

func (s *Solver) IsFork(c aoc.Coordinate) bool {
	neighbors := c.Neighbors4Way()
	sum := 0
	for _, n := range neighbors {
		node := s.grid.Get(n)
		if node == nil {
			continue
		}

		if node.Value() != ForestTile && node.Value() != PathTile {
			sum += 1
		}
	}

	return sum >= 3
}

func (s *Solver) NextCoordinates(c, p aoc.Coordinate) []aoc.Coordinate {
	coordinates := []aoc.Coordinate{}

	switch s.grid.Get(c).Value() {
	case SlopeUpTile:
		coordinates = append(coordinates, aoc.NewCoordinate(c.X, c.Y-1))
	case SlopeDownTile:
		coordinates = append(coordinates, aoc.NewCoordinate(c.X, c.Y+1))
	case SlopeLeftTile:
		coordinates = append(coordinates, aoc.NewCoordinate(c.X-1, c.Y))
	case SlopeRightTile:
		coordinates = append(coordinates, aoc.NewCoordinate(c.X+1, c.Y))
	case PathTile:
		coordinates = append(coordinates, aoc.NewCoordinate(c.X, c.Y-1))
		coordinates = append(coordinates, aoc.NewCoordinate(c.X, c.Y+1))
		coordinates = append(coordinates, aoc.NewCoordinate(c.X-1, c.Y))
		coordinates = append(coordinates, aoc.NewCoordinate(c.X+1, c.Y))
	}

	maxX, maxY := s.grid.MaxX(), s.grid.MaxY()

	filter := []aoc.Coordinate{}
	for _, cc := range coordinates {
		if cc == p {
			continue
		}

		if cc.X < 0 || cc.X > maxX || cc.Y < 0 || cc.Y > maxY {
			continue
		}

		switch s.grid.Get(cc).Value() {
		case SlopeUpTile:
			if cc.Y-c.Y == 1 {
				continue
			}
		case SlopeDownTile:
			if c.Y-cc.Y == 1 {
				continue
			}
		case SlopeLeftTile:
			if cc.X-c.X == 1 {
				continue
			}
		case SlopeRightTile:
			if c.X-cc.X == 1 {
				continue
			}
		case ForestTile:
			continue
		}

		filter = append(filter, cc)
	}

	return filter
}

func (s *Solver) StartCoordinate() aoc.Coordinate {
	maxX := s.grid.MaxX()
	for x := 0; x <= maxX; x++ {
		c := aoc.NewCoordinate(x, 0)
		if s.grid.Get(c).Value() == PathTile {
			return c
		}
	}

	return aoc.NewCoordinate(-1, -1)
}

func (s *Solver) EndCoordinate() aoc.Coordinate {
	maxX, y := s.grid.MaxX(), s.grid.MaxY()
	for x := 0; x <= maxX; x++ {
		c := aoc.NewCoordinate(x, y)
		if s.grid.Get(c).Value() == PathTile {
			return c
		}
	}

	return aoc.NewCoordinate(-1, -1)
}

type GraphEdge struct {
	from, to aoc.Coordinate
	distance int
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
