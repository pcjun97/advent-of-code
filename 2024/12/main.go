package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

func ParseGardenPlots(s string) *aoc.Grid {
	g := aoc.NewGrid()

	for y, line := range strings.Split(s, "\n") {
		for x, r := range line {
			c := aoc.NewCoordinate(x, y)
			g.Add(aoc.NewNode(c, int(r)))
		}
	}

	return g
}

type Solver struct {
	Plots *aoc.Grid
}

func NewSolver(input string) *Solver {
	s := Solver{ParseGardenPlots(input)}
	return &s
}

func (s *Solver) Regions() [][]aoc.Coordinate {
	regions := [][]aoc.Coordinate{}

	visited := make(map[aoc.Coordinate]struct{})

	var visit func(c aoc.Coordinate) []aoc.Coordinate
	visit = func(c aoc.Coordinate) []aoc.Coordinate {
		if _, ok := visited[c]; ok {
			return nil
		}
		visited[c] = struct{}{}

		coordinates := []aoc.Coordinate{c}
		for _, cc := range c.Neighbors4Way() {
			if !s.IsSameRegion(c, cc) {
				continue
			}
			coordinates = append(coordinates, visit(cc)...)
		}

		return coordinates
	}

	for _, node := range s.Plots.Nodes() {
		if _, ok := visited[node.Coordinate]; ok {
			continue
		}

		regions = append(regions, visit(node.Coordinate))
	}

	return regions
}

func (s *Solver) Part1() int {
	sum := 0

	for _, r := range s.Regions() {
		perimeter := 0

		for _, c := range r {
			for _, nc := range c.Neighbors4Way() {
				if s.Plots.Get(nc) != nil && s.Plots.Get(c).Value() == s.Plots.Get(nc).Value() {
					continue
				}
				perimeter++
			}
		}

		sum += len(r) * perimeter
	}

	return sum
}

const (
	left int = iota
	right
	top
	bottom
)

func (s *Solver) Part2() int {
	sum := 0

	for _, r := range s.Regions() {
		sides := make(map[int]map[int][]int)
		for i := 0; i < 4; i++ {
			sides[i] = make(map[int][]int)
		}

		for _, c := range r {
			if !s.IsSameRegion(c, aoc.NewCoordinate(c.X-1, c.Y)) {
				if _, ok := sides[left][c.X]; !ok {
					sides[left][c.X] = []int{}
				}
				sides[left][c.X] = append(sides[left][c.X], c.Y)
			}

			if !s.IsSameRegion(c, aoc.NewCoordinate(c.X+1, c.Y)) {
				if _, ok := sides[right][c.X]; !ok {
					sides[right][c.X] = []int{}
				}
				sides[right][c.X] = append(sides[right][c.X], c.Y)
			}

			if !s.IsSameRegion(c, aoc.NewCoordinate(c.X, c.Y-1)) {
				if _, ok := sides[top][c.Y]; !ok {
					sides[top][c.Y] = []int{}
				}
				sides[top][c.Y] = append(sides[top][c.Y], c.X)
			}

			if !s.IsSameRegion(c, aoc.NewCoordinate(c.X, c.Y+1)) {
				if _, ok := sides[bottom][c.Y]; !ok {
					sides[bottom][c.Y] = []int{}
				}
				sides[bottom][c.Y] = append(sides[bottom][c.Y], c.X)
			}
		}

		count := 0

		for i := range sides {
			for j := range sides[i] {
				sort.Ints(sides[i][j])

				for k, v := range sides[i][j] {
					if k == 0 || v - sides[i][j][k-1] > 1 {
						count++
					}
				}
			}
		}

		sum += len(r) * count
	}

	return sum
}

func (s *Solver) IsSameRegion(c1, c2 aoc.Coordinate) bool {
	n1 := s.Plots.Get(c1)
	n2 := s.Plots.Get(c2)

	if n1 == nil || n2 == nil {
		return false
	}

	return n1.Value() == n2.Value()
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
