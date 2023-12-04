package main

import (
	"fmt"
	"regexp"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	initial string
}

func NewSolver(input string) *Solver {
	s := Solver{input}
	return &s
}

func (s *Solver) Part1() string {
	p := s.initial
	for {
		p = RotatePassword(p)
		if Rule1(p) && Rule2(p) && Rule3(p) {
			break
		}
	}
	return p
}

func (s *Solver) Part2() string {
	p := s.Part1()
	for {
		p = RotatePassword(p)
		if Rule1(p) && Rule2(p) && Rule3(p) {
			break
		}
	}
	return p
}

func RotatePassword(s string) string {
	b := []byte(s)
	for i := len(b) - 1; i >= 0; i-- {
		b[i] = b[i] + 1
		if b[i] > 'z' {
			b[i] = 'a'
		} else {
			break
		}
	}
	return string(b)
}

func Rule1(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i+1]-s[i] == 1 && s[i+2]-s[i] == 2 {
			return true
		}
	}
	return false
}

func Rule2(s string) bool {
	r := regexp.MustCompile(`i|o|l`)
	return !r.MatchString(s)
}

func Rule3(s string) bool {
	findPairIndex := func(ss string) int {
		for i := 0; i < len(ss)-1; i++ {
			if ss[i] == ss[i+1] {
				return i
			}
		}
		return -1
	}
	i := findPairIndex(s)
	if i >= 0 {
		return findPairIndex(s[i+2:]) >= 0
	}
	return false
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
