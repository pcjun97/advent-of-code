package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	containers []int
}

func NewSolver(input string) *Solver {
	containers := []int{}
	for _, line := range strings.Split(input, "\n") {
		c, _ := strconv.Atoi(line)
		containers = append(containers, c)
	}
	s := Solver{containers}
	return &s
}

func (s *Solver) Part1(n int) int {
	return len(GenerateCombinations(n, s.containers, nil))
}

func (s *Solver) Part2(n int) int {
	l := []int{}
	min := math.MaxInt
	for _, c := range GenerateCombinations(n, s.containers, nil) {
		l = append(l, len(c))
		if len(c) < min {
			min = len(c)
		}
	}
	count := 0
	for _, x := range l {
		if x == min {
			count += 1
		}
	}
	return count
}

func GenerateCombinations(remaining int, containers []int, current []int) [][]int {
	if remaining == 0 {
		return [][]int{current}
	}

	com := [][]int{}
	for i, c := range containers {
		if c > remaining {
			continue
		}

		next := make([]int, len(current))
		copy(next, current)
		next = append(next, c)
		com = append(com, GenerateCombinations(remaining-c, containers[i+1:], next)...)
	}
	return com
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1(150))
	fmt.Println(s.Part2(150))
}
