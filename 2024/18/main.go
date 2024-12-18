package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

const (
	Empty int = iota
	Corrupted
)

type Solver struct {
	Size          int
	IncomingBytes []aoc.Coordinate
}

func NewSolver(input string, size int) *Solver {
	ib := []aoc.Coordinate{}
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Split(line, ",")
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		ib = append(ib, aoc.NewCoordinate(x, y))
	}

	s := Solver{size, ib}
	return &s
}

func (s *Solver) Part1(steps int) int {
	g := aoc.NewGrid()
	for y := 0; y <= s.Size; y++ {
		for x := 0; x <= s.Size; x++ {
			g.Add(aoc.NewNode(aoc.NewCoordinate(x, y), Empty))
		}
	}

	for i := 0; i < steps; i++ {
		ib := s.IncomingBytes[i]
		g.Get(ib).Set(Corrupted)
	}

	return BestScore(g)
}

func (s *Solver) Part2() string {
	min := 0
	max := len(s.IncomingBytes) - 1
	for min != max {
		mid := (min + max) / 2

		g := aoc.NewGrid()
		for y := 0; y <= s.Size; y++ {
			for x := 0; x <= s.Size; x++ {
				g.Add(aoc.NewNode(aoc.NewCoordinate(x, y), Empty))
			}
		}

		for i := 0; i <= mid; i++ {
			g.Get(s.IncomingBytes[i]).Set(Corrupted)
		}

		if BestScore(g) > 0 {
			min = mid + 1
		} else {
			max = mid
		}
	}

	c := s.IncomingBytes[min]
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func BestScore(g *aoc.Grid) int {
	minX := g.MinX()
	maxX := g.MaxX()
	minY := g.MinY()
	maxY := g.MaxY()

	start := aoc.NewCoordinate(minX, minY)
	end := aoc.NewCoordinate(maxX, maxY)

	bestscores := make(map[aoc.Coordinate]int)
	tovisit := []aoc.Coordinate{start}
	bestscores[start] = 0

	for {
		if len(tovisit) <= 0 {
			return -1
		}

		i := 0
		for j, v := range tovisit {
			if bestscores[v] < bestscores[tovisit[i]] {
				i = j
			}
		}

		visiting := tovisit[i]
		tovisit = append(tovisit[:i], tovisit[i+1:]...)

		for _, c := range visiting.Neighbors4Way() {
			node := g.Get(c)
			if node == nil || node.Value() == Corrupted {
				continue
			}

			if _, ok := bestscores[c]; ok {
				continue
			}
			bestscores[c] = bestscores[visiting] + 1

			if c == end {
				return bestscores[end]
			}

			tovisit = append(tovisit, c)
		}
	}
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input, 70)

	start := time.Now()
	fmt.Println(s.Part1(1024), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
