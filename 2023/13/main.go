package main

import (
	"fmt"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Pattern int

const (
	Ash Pattern = iota
	Rock
)

func parsePattern(r byte) Pattern {
	if r == '#' {
		return Rock
	}
	return Ash
}

func parseGrid(s string) *aoc.Grid {
	grid := aoc.NewGrid()

	for y, line := range strings.Split(s, "\n") {
		for x, r := range []byte(line) {
			c := aoc.NewCoordinate(x, y)
			grid.Add(aoc.NewNode(c, int(parsePattern(r))))
		}
	}

	return grid
}

type Solver struct {
	grids []*aoc.Grid
}

func NewSolver(input string) *Solver {
	grids := []*aoc.Grid{}
	for _, lines := range strings.Split(input, "\n\n") {
		grids = append(grids, parseGrid(lines))
	}
	s := Solver{grids}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0
	for _, grid := range s.grids {
		sum += reflectValue(grid, 0)
	}
	return sum
}

func (s *Solver) Part2() int {
	sum := 0
	for _, grid := range s.grids {
		sum += reflectValue(grid, 1)
	}
	return sum
}

func reflectValue(grid *aoc.Grid, diff int) int {
	maxX, maxY := grid.MaxX(), grid.MaxY()

	columns := [][]Pattern{}
	for x := 0; x <= maxX; x++ {
		col := []Pattern{}
		for y := 0; y <= maxY; y++ {
			c := aoc.NewCoordinate(x, y)
			col = append(col, Pattern(grid.Get(c).Value()))
		}
		columns = append(columns, col)
	}
	if i := reflectIndex(columns, diff); i >= 0 {
		return i + 1
	}

	rows := [][]Pattern{}
	for y := 0; y <= maxY; y++ {
		row := []Pattern{}
		for x := 0; x <= maxX; x++ {
			c := aoc.NewCoordinate(x, y)
			row = append(row, Pattern(grid.Get(c).Value()))
		}
		rows = append(rows, row)
	}

	if i := reflectIndex(rows, diff); i >= 0 {
		return (i + 1) * 100
	}

	return 0
}

func reflectIndex(patterns [][]Pattern, diff int) int {
	for i := range patterns[:len(patterns)-1] {
		d := 0
		for j := 0; i-j >= 0 && i+1+j < len(patterns); j++ {
			d += linesDiff(patterns[i-j], patterns[i+j+1])
		}
		if d == diff {
			return i
		}
	}

	return -1
}

func linesDiff(line1, line2 []Pattern) int {
	diff := 0
	for i := range line1 {
		if line1[i] != line2[i] {
			diff += 1
		}
	}
	return diff
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
