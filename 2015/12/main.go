package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	line string
}

func NewSolver(input string) *Solver {
	s := Solver{input}
	return &s
}

func (s *Solver) Part1() int {
	return SumJSONNumbers(s.line)
}

func (s *Solver) Part2() int {
	r := regexp.MustCompile(`:"red"`)

	line := s.line
	for {
		m := r.FindStringIndex(line)
		if m == nil {
			break
		}

		var start, end int

		c := 0
		for i := m[0]; i >= 0; i-- {
			if line[i] == '}' {
				c += 1
			}
			if line[i] == '{' {
				c -= 1
				if c < 0 {
					start = i
					break
				}
			}
		}

		c = 0
		for i := m[1]; i < len(line); i++ {
			if line[i] == '{' {
				c += 1
			}
			if line[i] == '}' {
				c -= 1
				if c < 0 {
					end = i
					break
				}
			}
		}
		line = line[:start+1] + line[end:]
	}

	return SumJSONNumbers(line)
}

func SumJSONNumbers(line string) int {
	sum := 0
	r := regexp.MustCompile(`\W(-?\d+)\W`)

	for {
		m := r.FindStringSubmatchIndex(line)
		if m == nil {
			break
		}
		n, _ := strconv.Atoi(line[m[2]:m[3]])
		sum += n
		line = line[m[3]-1:]
	}
	return sum
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
