package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Section struct {
	start int
	end   int
}

func NewSection(s string) *Section {
	i := strings.IndexRune(s, '-')

	start, err := strconv.Atoi(s[:i])
	if err != nil {
		panic("error parsing section assignment: " + s[:i])
	}

	end, err := strconv.Atoi(s[i+1:])
	if err != nil {
		panic("error parsing section assignment: " + s[i+1:])
	}

	section := Section{
		start: start,
		end:   end,
	}

	return &section
}

func (s *Section) Contains(other *Section) bool {
	return (s.start <= other.start) && (s.end >= other.end)
}

func (s *Section) Overlaps(other *Section) bool {
	if s.start <= other.start && s.end >= other.start {
		return true
	}

	if s.start <= other.end && s.end >= other.end {
		return true
	}

	if other.Contains(s) {
		return true
	}

	return false
}

type Pair struct {
	First  *Section
	Second *Section
}

func NewPair(s string) *Pair {
	i := strings.IndexRune(s, ',')

	s1 := NewSection(s[:i])
	s2 := NewSection(s[i+1:])

	p := Pair{
		First:  s1,
		Second: s2,
	}

	return &p
}

func main() {
	input := aoc.ReadInput()

	pairs := []*Pair{}
	for _, line := range strings.Split(input, "\n") {
		pair := NewPair(line)
		pairs = append(pairs, pair)
	}

	count := 0
	for _, p := range pairs {
		if p.First.Contains(p.Second) || p.Second.Contains(p.First) {
			count++
		}
	}

	fmt.Println(count)

	count = 0
	for _, p := range pairs {
		if p.First.Overlaps(p.Second) {
			count++
		}
	}

	fmt.Println(count)
}
