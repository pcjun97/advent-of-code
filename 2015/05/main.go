package main

import (
	"fmt"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func HasAtLeastThreeVowels(s string) bool {
	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	count := 0
	for _, r := range s {
		if _, ok := vowels[r]; ok {
			count += 1
			if count >= 3 {
				return true
			}
		}
	}

	return false
}

func HasALetterTwiceInARow(s string) bool {
	for c := 'a'; c <= 'z'; c++ {
		if strings.Contains(s, fmt.Sprintf("%c%c", c, c)) {
			return true
		}
	}
	return false
}

func HasNoIllegalString(s string) bool {
	banned := []string{
		"ab",
		"cd",
		"pq",
		"xy",
	}
	for _, b := range banned {
		if strings.Contains(s, b) {
			return false
		}
	}
	return true
}

func HasTwoNonOverlappingPairs(s string) bool {
	for i := 0; i <= len(s)-4; i++ {
		pair := s[i : i+2]
		if strings.Contains(s[i+2:], pair) {
			return true
		}
	}
	return false
}

func HasRepeatingLetterWithLetterInBetween(s string) bool {
	for i := 0; i <= len(s)-3; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

type Solver struct {
	lines []string
}

func NewSolver(input string) *Solver {
	s := Solver{
		lines: strings.Split(input, "\n"),
	}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0
	for _, line := range s.lines {
		if HasAtLeastThreeVowels(line) && HasALetterTwiceInARow(line) && HasNoIllegalString(line) {
			sum += 1
		}
	}
	return sum
}

func (s *Solver) Part2() int {
	sum := 0
	for _, line := range s.lines {
		if HasTwoNonOverlappingPairs(line) && HasRepeatingLetterWithLetterInBetween(line) {
			sum += 1
		}
	}
	return sum
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
