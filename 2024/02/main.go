package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Report struct {
	Levels []int
}

func ParseReport(s string) Report {
	levels := []int{}
	for _, f := range strings.Fields(s) {
		v, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}

		levels = append(levels, v)
	}
	return Report{levels}
}

func (r Report) IsSafe() bool {
	d := r.Levels[1] - r.Levels[0]

	if d == 0 || d < -3 || d > 3 {
		return false
	}

	increase := d > 0

	for i := 2; i < len(r.Levels); i++ {
		d = r.Levels[i] - r.Levels[i-1]

		if d == 0 || d < -3 || d > 3 {
			return false
		}

		if (increase && d < 0) || (!increase && d > 0) {
			return false
		}
	}

	return true
}

func (r Report) IsSafeWithProblemDampener() bool {
	if r.IsSafe() {
		return true
	}

	for i := 0; i < len(r.Levels); i++ {
		levels := make([]int, len(r.Levels)-1)
		levels = append(levels, r.Levels[:i]...)
		levels = append(levels, r.Levels[i+1:]...)
		rr := Report{levels}

		if rr.IsSafe() {
			return true
		}
	}

	return false
}

type Solver struct {
	Reports []Report
}

func NewSolver(input string) *Solver {
	reports := []Report{}

	for _, line := range strings.Split(input, "\n") {
		reports = append(reports, ParseReport(line))
	}

	s := Solver{reports}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0

	for _, r := range s.Reports {
		if r.IsSafe() {
			sum++
		}
	}

	return sum
}

func (s *Solver) Part2() int {
	sum := 0

	for _, r := range s.Reports {
		if r.IsSafeWithProblemDampener() {
			sum++
		}
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
