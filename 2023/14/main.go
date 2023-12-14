package main

import (
	"fmt"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

const (
	RoundedRock = iota
	CubeRock
	EmptySpace
)

func moveRocks(row []int) []int {
	var current []int
	next := row
	for !rowsEqual(current, next) {
		current = next
		next = []int{}
		for i, rock := range current {
			if rock == RoundedRock && i-1 >= 0 && next[i-1] == EmptySpace {
				next[i-1] = RoundedRock
				next = append(next, EmptySpace)
				continue
			}
			next = append(next, rock)
		}
	}

	return current
}

func rowsEqual(row1, row2 []int) bool {
	if len(row1) != len(row2) {
		return false
	}

	for i := range row1 {
		if row1[i] != row2[i] {
			return false
		}
	}

	return true
}

func parseGrid(s string) *aoc.Grid {
	g := aoc.NewGrid()
	for y, line := range strings.Split(s, "\n") {
		for x, r := range []byte(line) {
			c := aoc.NewCoordinate(x, y)
			var rock int
			switch r {
			case 'O':
				rock = RoundedRock
			case '#':
				rock = CubeRock
			case '.':
				rock = EmptySpace
			}
			g.Add(aoc.NewNode(c, rock))
		}
	}
	return g
}

func gridToString(g *aoc.Grid) string {
	var b strings.Builder
	maxX, maxY := g.MaxX(), g.MaxY()
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			c := aoc.NewCoordinate(x, y)
			switch g.Get(c).Value() {
			case RoundedRock:
				b.WriteByte('O')
			case CubeRock:
				b.WriteByte('#')
			case EmptySpace:
				b.WriteByte('.')
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func moveRocksInGrid(grid *aoc.Grid, direction Direction) *aoc.Grid {
	maxX, maxY := grid.MaxX(), grid.MaxY()
	result := aoc.NewGrid()

	switch direction {
	case North:
		for x := 0; x <= maxX; x++ {
			row := []int{}
			for y := 0; y <= maxY; y++ {
				c := aoc.NewCoordinate(x, y)
				row = append(row, grid.Get(c).Value())
			}
			row = moveRocks(row)
			for y := 0; y <= maxY; y++ {
				c := aoc.NewCoordinate(x, y)
				result.Add(aoc.NewNode(c, row[y]))
			}
		}
	case South:
		for x := 0; x <= maxX; x++ {
			row := []int{}
			for y := maxY; y >= 0; y-- {
				c := aoc.NewCoordinate(x, y)
				row = append(row, grid.Get(c).Value())
			}
			row = moveRocks(row)
			for y := 0; y <= maxY; y++ {
				c := aoc.NewCoordinate(x, maxY-y)
				result.Add(aoc.NewNode(c, row[y]))
			}
		}
	case West:
		for y := 0; y <= maxY; y++ {
			row := []int{}
			for x := 0; x <= maxX; x++ {
				c := aoc.NewCoordinate(x, y)
				row = append(row, grid.Get(c).Value())
			}
			row = moveRocks(row)
			for x := 0; x <= maxX; x++ {
				c := aoc.NewCoordinate(x, y)
				result.Add(aoc.NewNode(c, row[x]))
			}
		}
	case East:
		for y := 0; y <= maxY; y++ {
			row := []int{}
			for x := maxX; x >= 0; x-- {
				c := aoc.NewCoordinate(x, y)
				row = append(row, grid.Get(c).Value())
			}
			row = moveRocks(row)
			for x := 0; x <= maxX; x++ {
				c := aoc.NewCoordinate(maxX-x, y)
				result.Add(aoc.NewNode(c, row[x]))
			}
		}
	}

	return result
}

func supportBeamLoad(grid *aoc.Grid, direction Direction) int {
	sum := 0
	maxX, maxY := grid.MaxX(), grid.MaxY()

	for _, n := range grid.Nodes() {
		if n.Value() != RoundedRock {
			continue
		}
		switch direction {
		case North:
			sum += maxY - n.Y + 1
		case South:
			sum += n.Y + 1
		case East:
			sum += n.X + 1
		case West:
			sum += maxX - n.X + 1
		}
	}

	return sum
}

type Solver struct {
	grid *aoc.Grid
}

func NewSolver(input string) *Solver {
	grid := parseGrid(input)
	s := Solver{grid}
	return &s
}

func (s *Solver) Part1() int {
	g := moveRocksInGrid(s.grid, North)
	return supportBeamLoad(g, North)
}

func (s *Solver) Part2() int {
	cache := make(map[string]int)
	g := s.grid
	for i := 1000000000; i > 0; i-- {
		gs := gridToString(g)
		if j, ok := cache[gs]; ok {
			cycle := j - i
			rem := i % cycle
			target := j - rem
			for str, k := range cache {
				if k == target {
					return supportBeamLoad(parseGrid(str), North)
				}
			}
		}
		cache[gs] = i

		g = moveRocksInGrid(g, North)
		g = moveRocksInGrid(g, West)
		g = moveRocksInGrid(g, South)
		g = moveRocksInGrid(g, East)
	}
	return 0
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
