package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Guest struct {
	name      string
	happiness map[string]int
}

func parseGuests(s string) map[string]Guest {
	guests := make(map[string]Guest)

	r := regexp.MustCompile(`(.*) would (.*) (.*) happiness units by sitting next to (.*).`)
	for _, line := range strings.Split(s, "\n") {
		m := r.FindStringSubmatch(line)
		name := m[1]
		change := m[2]
		happiness := m[3]
		neighbor := m[4]

		g, ok := guests[name]
		if !ok {
			g = Guest{
				name:      name,
				happiness: make(map[string]int),
			}
			guests[name] = g
		}

		h, _ := strconv.Atoi(happiness)
		switch change {
		case "gain":
			g.happiness[neighbor] = h
		case "lose":
			g.happiness[neighbor] = -h
		}
	}

	return guests
}

type Solver struct {
	guests map[string]Guest
}

func NewSolver(input string) *Solver {
	guests := parseGuests(input)
	s := Solver{guests}
	return &s
}

func (s *Solver) Part1() int {
	guestList := []Guest{}
	for _, g := range s.guests {
		guestList = append(guestList, g)
	}

	p := Permutation(guestList[1:], nil)
	for i := range p {
		p[i] = append(p[i], guestList[0])
	}

	max := math.MinInt
	for _, l := range p {
		c := s.TotalHappinessChange(l)
		if c > max {
			max = c
		}
	}

	return max
}

func (s *Solver) Part2() int {
	self := Guest{"", make(map[string]int)}
	guestList := []Guest{}
	for _, g := range s.guests {
		guestList = append(guestList, g)
	}

	p := Permutation(guestList, nil)
	for i := range p {
		p[i] = append(p[i], self)
	}

	max := math.MinInt
	for _, l := range p {
		c := s.TotalHappinessChange(l)
		if c > max {
			max = c
		}
	}

	return max
}

func (s *Solver) TotalHappinessChange(list []Guest) int {
	sum := 0

	for i, g := range list {
		l := i - 1
		if l < 0 {
			l = len(list) - 1
		}
		left := s.guests[list[l].name]
		sum += s.guests[g.name].happiness[left.name]

		r := i + 1
		if r >= len(list) {
			r = 0
		}
		right := s.guests[list[r].name]
		sum += s.guests[g.name].happiness[right.name]
	}

	return sum
}

func Permutation(list []Guest, p []Guest) [][]Guest {
	if len(list) == 0 {
		return [][]Guest{p}
	}

	result := [][]Guest{}

	for i, g := range list {
		q := make([]Guest, len(p))
		copy(q, p)
		q = append(q, g)

		l := make([]Guest, i)
		copy(l, list[:i])
		l = append(l, list[i+1:]...)

		result = append(result, Permutation(l, q)...)
	}

	return result
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
