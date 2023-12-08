package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	row, column int
}

func NewSolver(input string) *Solver {
	r := regexp.MustCompile(`\d+`)
	m := r.FindAllString(input, -1)
	row, _ := strconv.Atoi(m[0])
	column, _ := strconv.Atoi(m[1])
	s := Solver{row, column}
	return &s
}

func (s *Solver) Part1(initial, row, column int) int {
	result := initial
	target := PositionToIndex(s.row, s.column)
	for i := PositionToIndex(row, column); i < target; i++ {
		result = NextCode(result)
	}
	return result
}

func PositionToIndex(row, column int) int {
	x := column - 1
	y := row + x

	s := 0
	for i := 1; i < y; i++ {
		s += i
	}
	return s + x
}

func NextCode(from int) int {
	result := from * 252533
	result = result % 33554393
	return result
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1(27995004, 6, 6))
}
