package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type CoordinatePair [2]aoc.Coordinate

func NewCoordinatePair(c1, c2 aoc.Coordinate) CoordinatePair {
	if c2.Y < c1.Y || (c2.Y == c1.Y && c2.X < c1.X) {
		c1, c2 = c2, c1
	}
	return [2]aoc.Coordinate{c1, c2}
}

func parseImage(s string) *aoc.Grid {
	image := aoc.NewGrid()
	for y, line := range strings.Split(s, "\n") {
		for x, r := range []byte(line) {
			if r != '#' {
				continue
			}
			c := aoc.NewCoordinate(x, y)
			image.Add(aoc.NewNode(c, 0))
		}
	}
	return image
}

type Solver struct {
	image *aoc.Grid
}

func NewSolver(input string) *Solver {
	image := parseImage(input)
	s := Solver{image}
	return &s
}

func (s *Solver) Part1() int {
	expanded := expandImage(s.image, 2)
	galaxies := expanded.Nodes()

	sum := 0
	for i, n := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += n.Coordinate.ManhattanDistance(galaxies[j].Coordinate)
		}
	}

	return sum
}

func (s *Solver) Part2(d int) int {
	expanded := expandImage(s.image, d)
	galaxies := expanded.Nodes()

	sum := 0
	for i, n := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += n.Coordinate.ManhattanDistance(galaxies[j].Coordinate)
		}
	}

	return sum
}

func expandImage(image *aoc.Grid, d int) *aoc.Grid {
	xMap := make(map[int]struct{})
	yMap := make(map[int]struct{})
	for _, n := range image.Nodes() {
		xMap[n.X] = struct{}{}
		yMap[n.Y] = struct{}{}
	}

	xList := []int{}
	yList := []int{}
	for x := range xMap {
		xList = append(xList, x)
	}
	for y := range yMap {
		yList = append(yList, y)
	}

	sort.Ints(xList)
	sort.Ints(yList)

	refX := make(map[int]int)
	refY := make(map[int]int)
	for i, x := range xList {
		if i == 0 {
			refX[x] = x * d
			continue
		}
		dx := x - xList[i-1] - 1
		refX[x] = refX[xList[i-1]] + dx*d + 1
	}
	for i, y := range yList {
		if i == 0 {
			refY[y] = y * d
			continue
		}
		dy := y - yList[i-1] - 1
		refY[y] = refY[yList[i-1]] + dy*d + 1
	}

	expanded := aoc.NewGrid()
	for _, n := range image.Nodes() {
		x := refX[n.X]
		y := refY[n.Y]
		c := aoc.NewCoordinate(x, y)
		node := aoc.NewNode(c, 0)
		expanded.Add(node)
	}

	return expanded
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2(1000000))
}
