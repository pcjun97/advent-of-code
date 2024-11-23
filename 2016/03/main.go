package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Triangle struct {
	sides [3]int
}

func NewTriangle(s1, s2, s3 int) Triangle {
	return Triangle{[3]int{s1, s2, s3}}
}

func parseTriangle(str string) (Triangle, error) {
	var zero Triangle

	r := regexp.MustCompile(`^\s*(\d+)\s+(\d+)\s+(\d+)\s*$`)
	m := r.FindStringSubmatch(str)

	s1, err := strconv.Atoi(m[1])
	if err != nil {
		return zero, err
	}

	s2, err := strconv.Atoi(m[2])
	if err != nil {
		return zero, err
	}

	s3, err := strconv.Atoi(m[3])
	if err != nil {
		return zero, err
	}

	return Triangle{[3]int{s1, s2, s3}}, nil
}

func (t Triangle) Valid() bool {
	sum := 0
	for _, s := range t.sides {
		sum += s
	}

	for _, s := range t.sides {
		if (sum - s) <= s {
			return false
		}
	}

	return true
}

type Solver struct {
	Sides [][]int
}

func NewSolver(input string) *Solver {
	sides := [][]int{}

	for _, line := range strings.Split(input, "\n") {
		ss := []int{}

		for _, f := range strings.Fields(line) {
			s, err := strconv.Atoi(f)
			if err != nil {
				log.Fatal(err)
			}

			ss = append(ss, s)
		}

		sides = append(sides, ss)
	}

	s := Solver{sides}
	return &s
}

func (s *Solver) Part1() int {
	// invalid count
	ic := 0

	for _, s := range s.Sides {
		t := NewTriangle(s[0], s[1], s[2])
		if t.Valid() {
			ic += 1
		}
	}

	return ic
}

func (s *Solver) Part2() int {
	ic := 0

	if len(s.Sides) <= 0 {
		return ic
	}

	if (len(s.Sides) % 3) != 0 {
		log.Fatal("sides not multiple of 3")
	}

	for i := 0; i < len(s.Sides[0]); i++ {
		for j := 0; j < len(s.Sides); j += 3 {
			t := NewTriangle(s.Sides[j][i], s.Sides[j+1][i], s.Sides[j+2][i])
			if t.Valid() {
				ic += 1
			}
		}
	}

	return ic
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
