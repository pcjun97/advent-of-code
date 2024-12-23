package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

var NumericKeypad = map[rune]aoc.Coordinate{
	'A': aoc.NewCoordinate(2, 3),
	'0': aoc.NewCoordinate(1, 3),
	'1': aoc.NewCoordinate(0, 2),
	'2': aoc.NewCoordinate(1, 2),
	'3': aoc.NewCoordinate(2, 2),
	'4': aoc.NewCoordinate(0, 1),
	'5': aoc.NewCoordinate(1, 1),
	'6': aoc.NewCoordinate(2, 1),
	'7': aoc.NewCoordinate(0, 0),
	'8': aoc.NewCoordinate(1, 0),
	'9': aoc.NewCoordinate(2, 0),
}

var DirectionalKeypad = map[rune]aoc.Coordinate{
	'A': aoc.NewCoordinate(2, 0),
	'^': aoc.NewCoordinate(1, 0),
	'v': aoc.NewCoordinate(1, 1),
	'<': aoc.NewCoordinate(0, 1),
	'>': aoc.NewCoordinate(2, 1),
}

type Solver struct {
	Codes []string
}

func NewSolver(input string) *Solver {
	codes := strings.Split(input, "\n")
	s := Solver{codes}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0
	for _, c := range s.Codes {
		sum += extractNumeric(c) * ShortestInputs(c, 3)
	}
	return sum
}

func (s *Solver) Part2() int {
	sum := 0
	for _, c := range s.Codes {
		sum += extractNumeric(c) * ShortestInputs(c, 26)
	}
	return sum
}

func ShortestInputs(code string, nkeypads int) int {
	cache := make([]map[string]int, nkeypads+1)
	for i := range cache {
		cache[i] = make(map[string]int)
	}

	var f func(code string, nkeypads int) int
	f = func(code string, nkeypads int) int {
		if nkeypads == 0 {
			return len(code)
		}

		if _, ok := cache[nkeypads][code]; ok {
			return cache[nkeypads][code]
		}

		v := 0

		for _, d := range ToDirectional(code, DirectionalKeypad) {
			shortest := f(d[0], nkeypads-1)

			for _, dd := range d[1:] {
				final := f(dd, nkeypads-1)
				if final < shortest {
					shortest = final
				}
			}

			v += shortest
		}

		cache[nkeypads][code] = v
		return v
	}

	v := 0
	for _, d := range ToDirectional(code, NumericKeypad) {
		shortest := f(d[0], nkeypads-1)

		for _, dd := range d[1:] {
			final := f(dd, nkeypads-1)
			if final < shortest {
				shortest = final
			}
		}

		v += shortest
	}

	return v
}

func ToDirectional(code string, keypad map[rune]aoc.Coordinate) [][]string {
	results := [][]string{}
	cur := keypad['A']

	for _, r := range code {
		moves := ""
		c := keypad[r]
		dx := c.X - cur.X
		dy := c.Y - cur.Y

		for i := dy; i < 0; i++ {
			moves += "^"
		}

		for i := 0; i < dy; i++ {
			moves += "v"
		}

		for i := dx; i < 0; i++ {
			moves += "<"
		}

		for i := 0; i < dx; i++ {
			moves += ">"
		}

		valid := []string{}
		for _, p := range permute(moves) {
			if IsValidInput(p, cur, keypad) {
				valid = append(valid, fmt.Sprintf("%s%c", p, 'A'))
			}
		}

		results = append(results, valid)
		cur = c
	}

	if len(results) == 0 {
		results = append(results, []string{"A"})
	}

	return results
}

func IsValidInput(input string, from aoc.Coordinate, keypad map[rune]aoc.Coordinate) bool {
	m := make(map[aoc.Coordinate]struct{})
	for _, c := range keypad {
		m[c] = struct{}{}
	}

	cur := from
	for _, r := range input {
		switch r {
		case '<':
			cur.X--
		case '>':
			cur.X++
		case '^':
			cur.Y--
		case 'v':
			cur.Y++
		}
		if _, ok := m[cur]; !ok {
			return false
		}
	}

	return true
}

func extractNumeric(code string) int {
	r := regexp.MustCompile(`0*(\d+)`)
	m := r.FindStringSubmatch(code)
	v, _ := strconv.Atoi(m[1])
	return v
}

func permute(s string) []string {
	if len(s) <= 1 {
		return []string{s}
	}

	set := make(map[string]struct{})
	for i, r := range s {
		next := s[:i]
		next += s[i+1:]
		for _, p := range permute(next) {
			set[fmt.Sprintf("%c%s", r, p)] = struct{}{}
		}
	}

	list := []string{}
	for ss := range set {
		list = append(list, ss)
	}
	return list
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
