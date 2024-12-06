package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	Rules   map[int]map[int]struct{}
	Updates [][]int
}

func NewSolver(input string) *Solver {
	blocks := strings.Split(input, "\n\n")
	if len(blocks) > 2 {
		log.Fatal("more than 2 blocks")
	}

	r := regexp.MustCompile(`^(\d+)\|(\d+)$`)
	rules := make(map[int]map[int]struct{})
	for _, line := range strings.Split(blocks[0], "\n") {
		m := r.FindStringSubmatch(line)
		v1, _ := strconv.Atoi(m[1])
		v2, _ := strconv.Atoi(m[2])

		if _, ok := rules[v1]; !ok {
			rules[v1] = make(map[int]struct{})
		}
		if _, ok := rules[v2]; !ok {
			rules[v2] = make(map[int]struct{})
		}
		rules[v1][v2] = struct{}{}
	}

	updates := [][]int{}
	for _, line := range strings.Split(blocks[1], "\n") {
		fields := strings.Split(line, ",")

		fint := []int{}
		for _, f := range fields {
			v, _ := strconv.Atoi(f)
			fint = append(fint, v)
		}

		updates = append(updates, fint)
	}

	s := Solver{rules, updates}
	return &s
}

func (s *Solver) IsValidUpdate(update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		curr := update[i]
		next := update[i+1]

		if _, ok := s.Rules[curr]; !ok {
			log.Fatalf("rule not found for: %d", curr)
		}

		if _, ok := s.Rules[curr][next]; ok {
			continue
		}

		if _, okk := s.Rules[next][curr]; !okk {
			log.Fatalf("relation not found for: %d %d", curr, next)
		}

		return false
	}

	return true
}

func middleOfIntSlice(ints []int) int {
	if len(ints)%2 == 0 {
		log.Fatal("even number of elements in slice")
	}

	return ints[len(ints)/2]
}

func (s *Solver) Part1() int {
	sum := 0

	for _, u := range s.Updates {
		if !s.IsValidUpdate(u) {
			continue
		}

		sum += middleOfIntSlice(u)
	}

	return sum
}

func (s *Solver) Part2() int {
	sum := 0

	for _, u := range s.Updates {
		uu := CopyIntSlice(u)
		if s.IsValidUpdate(uu) {
			continue
		}

		sort.Slice(uu, func(i, j int) bool {
			_, ok := s.Rules[uu[i]][uu[j]]
			return ok
		})

		sum += middleOfIntSlice(uu)
	}

	return sum
}

func CopyIntSlice(ints []int) []int {
	copy := make([]int, len(ints))
	for i, v := range ints {
		copy[i] = v
	}
	return copy
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
