package main

import (
	"fmt"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

var (
	pairsPart1 = map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"0": 0,
	}

	pairsPart2 = map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"0":     0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}
)

type Solver struct {
	lines []string
}

func NewSolver(input string) *Solver {
	s := Solver{
		lines: strings.Split(input, "\n"),
	}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0

	for _, line := range s.lines {
		sum += extractCalibrationValue(line, pairsPart1)
	}

	return sum
}

func (s *Solver) Part2() int {
	sum := 0

	for _, line := range s.lines {
		sum += extractCalibrationValue(line, pairsPart2)
	}

	return sum
}

func extractCalibrationValue(line string, pairs map[string]int) int {
	max := 0
	min := -1

	for k := range pairs {
		if len(k) > max {
			max = len(k)
		}

		if min < 0 || len(k) < min {
			min = len(k)
		}
	}

	first := -1
	last := -1

	for i := range line {
		for j := i + min; j <= i+max && j <= len(line); j++ {
			digit, ok := pairs[line[i:j]]
			if !ok {
				continue
			}

			if first < 0 {
				first = digit
			}
			last = digit
		}
	}

	return (first * 10) + last
}

func main() {
	input := aoc.ReadInput()

	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
