package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type IPv7 struct {
	Sections []IPv7Section
}

func ParseIPv7(s string) IPv7 {
	sections := []IPv7Section{}

	for {
		hyper := false
		value := s

		i := strings.IndexAny(s, "[]")
		if i > 0 {
			hyper = s[i] == ']'
			value = s[:i]
			s = s[i+1:]
		}

		if len(value) <= 0 {
			continue
		}

		section := IPv7Section{
			Value:    value,
			Hypernet: hyper,
		}
		sections = append(sections, section)

		if i < 0 {
			break
		}
	}

	ip := IPv7{sections}
	return ip
}

func (ip IPv7) TLS() bool {
	tls := false
	for _, s := range ip.Sections {
		if s.Hypernet && s.ABBA() {
			return false
		}

		if s.ABBA() {
			tls = true
		}
	}
	return tls
}

type IPv7Section struct {
	Value    string
	Hypernet bool
}

func (s IPv7Section) ABBA() bool {
	for i := 0; i < len(s.Value)-3; i++ {
		abba := s.Value[i : i+4]
		if abba[0] != abba[1] && abba[0] == abba[3] && abba[1] == abba[2] {
			return true
		}
	}
	return false
}

type Solver struct {
	ips []IPv7
}

func NewSolver(input string) *Solver {
	ips := []IPv7{}
	for _, line := range strings.Split(input, "\n") {
		ips = append(ips, ParseIPv7(line))
	}

	s := Solver{ips}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0

	for _, ip := range s.ips {
		if ip.TLS() {
			sum++
		}
	}

	return sum
}

func (s *Solver) Part2() int {
	return 0
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
