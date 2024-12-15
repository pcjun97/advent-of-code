package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	Warehouse *Warehouse
	Movements []Direction
}

func ParseMovements(s string) []Direction {
	movements := []Direction{}

	for _, r := range s {
		var d Direction
		switch r {
		case '<':
			d = Left
		case '>':
			d = Right
		case '^':
			d = Up
		case 'v':
			d = Down
		default:
			continue
		}

		movements = append(movements, d)
	}

	return movements
}

func NewSolver(input string) *Solver {
	blocks := strings.Split(input, "\n\n")
	s := Solver{ParseWarehouse(blocks[0]), ParseMovements(blocks[1])}
	return &s
}

func (s *Solver) Part1() int {
	w := s.Warehouse.Clone()

	for _, m := range s.Movements {
		w.Move(w.Robot, m)
	}

	sum := 0

	for _, box := range w.Boxes {
		sum += (box.Y * 100) + box.X
	}

	return sum
}

func (s *Solver) Part2() int {
	w := s.Warehouse.Clone()
	w.ExpandWidth(2)

	for _, m := range s.Movements {
		w.Move(w.Robot, m)
	}

	sum := 0

	for _, box := range w.Boxes {
		sum += (box.Y * 100) + box.X
	}

	return sum
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
