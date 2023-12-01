package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Dimension [3]int

func parseDimension(s string) (Dimension, error) {
	var zero Dimension
	var d Dimension
	var err error

	for i, f := range strings.Split(s, "x") {
		d[i], err = strconv.Atoi(f)
		if err != nil {
			return zero, err
		}
	}

	return d, nil
}

func (d Dimension) SmallestSurfaceArea() int {
	switch {
	case d[0] >= d[1] && d[0] >= d[2]:
		return d[1] * d[2]
	case d[1] >= d[0] && d[1] >= d[2]:
		return d[0] * d[2]
	default:
		return d[0] * d[1]
	}
}

func (d Dimension) TotalSurfaceArea() int {
	return (2 * d[0] * d[1]) + (2 * d[1] * d[2]) + (2 * d[0] * d[2])
}

func (d Dimension) Volume() int {
	return d[0] * d[1] * d[2]
}

func (d Dimension) SmallestSurfacePerimeter() int {
	switch {
	case d[0] >= d[1] && d[0] >= d[2]:
		return 2 * (d[1] + d[2])
	case d[1] >= d[0] && d[1] >= d[2]:
		return 2 * (d[0] + d[2])
	default:
		return 2 * (d[0] + d[1])
	}
}

type Solver struct {
	dimensions []Dimension
}

func NewSolver(input string) (*Solver, error) {
	dimensions := []Dimension{}

	for _, line := range strings.Split(input, "\n") {
		d, err := parseDimension(line)
		if err != nil {
			return nil, err
		}

		dimensions = append(dimensions, d)
	}

	s := Solver{dimensions}
	return &s, nil
}

func (s *Solver) Part1() int {
	sum := 0

	for _, d := range s.dimensions {
		sum += d.SmallestSurfaceArea() + d.TotalSurfaceArea()
	}

	return sum
}

func (s *Solver) Part2() int {
	sum := 0

	for _, d := range s.dimensions {
		sum += d.SmallestSurfacePerimeter() + d.Volume()
	}

	return sum
}

func main() {
	input := aoc.ReadInput()

	s, err := NewSolver(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
