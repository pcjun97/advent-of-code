package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Package struct {
	ID, Weight int
}

type Group []Package

func NewGroup() Group {
	g := Group{}
	return g
}

func (g Group) Copy() Group {
	gg := make(Group, len(g))
	copy(gg, g)
	return gg
}

func (g Group) Difference(gg Group) Group {
	ref := make(map[Package]struct{})
	for _, p := range gg {
		ref[p] = struct{}{}
	}
	result := NewGroup()
	for _, p := range g {
		if _, ok := ref[p]; !ok {
			result = append(result, p)
		}
	}
	return result
}

func (g Group) HasSubsetOfMinLengthWithWeight(minLength, weight int) bool {
	for i := minLength; i < len(g); i++ {
		if len(g.CombinationsOfNWithWeight(i, weight)) > 0 {
			return true
		}
	}
	return false
}

func (g Group) CombinationsOfNWithWeight(n int, weight int) []Group {
	result := []Group{}

	var combinations func([]Package, Group)
	combinations = func(list []Package, current Group) {
		if len(current) >= n {
			return
		}

		for i, p := range list {
			w := current.Weight() + p.Weight
			if w > weight {
				continue
			}
			next := current.Copy()
			next = append(next, p)
			if w == weight {
				result = append(result, next)
				continue
			}
			combinations(list[i+1:], next)
		}
	}
	combinations(g, NewGroup())
	return result
}

func (g Group) Weight() int {
	sum := 0
	for _, p := range g {
		sum += p.Weight
	}
	return sum
}

func (g Group) QE() int {
	qe := 1
	for _, p := range g {
		qe *= p.Weight
	}
	return qe
}

func (g Group) Less(gg Group) bool {
	if len(g) < len(gg) {
		return true
	}
	if len(g) == len(gg) && g.QE() < gg.QE() {
		return true
	}
	return false
}

type Solver struct {
	group Group
}

func NewSolver(input string) *Solver {
	group := NewGroup()
	for i, line := range strings.Split(input, "\n") {
		w, _ := strconv.Atoi(line)
		p := Package{i, w}
		group = append(group, p)
	}
	s := Solver{group}
	return &s
}

func (s *Solver) Part1() int {
	var min Group = nil
	target := s.group.Weight() / 3
	for i := 1; i < len(s.group) && min == nil; i++ {
		for _, g := range s.group.CombinationsOfNWithWeight(i, target) {
			remaining := s.group.Difference(g)
			if !remaining.HasSubsetOfMinLengthWithWeight(i, target) {
				continue
			}
			if min == nil || g.Less(min) {
				min = g
			}
		}
	}

	return min.QE()
}

func (s *Solver) Part2() int {
	var min Group = nil
	target := s.group.Weight() / 4
	for i := 1; i < len(s.group) && min == nil; i++ {
		for _, g := range s.group.CombinationsOfNWithWeight(i, target) {
			remaining := s.group.Difference(g)
			if !remaining.HasSubsetOfMinLengthWithWeight(i, target) {
				continue
			}
			if min == nil || g.Less(min) {
				min = g
			}
		}
	}

	return min.QE()
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
	fmt.Println(time.Since(start))
}
