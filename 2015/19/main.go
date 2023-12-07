package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Replacement struct {
	from, to string
}

func parseReplacement(s string) Replacement {
	r := regexp.MustCompile(`^(.*) => (.*)$`)
	m := r.FindStringSubmatch(s)
	from := m[1]
	to := m[2]
	replacement := Replacement{from, to}
	return replacement
}

func (r Replacement) AllReplaceOnce(s string) []string {
	result := []string{}

	i := 0
	for {
		j := strings.Index(s[i:], r.from)
		if j < 0 {
			break
		}

		replaced := s[:i+j] + r.to + s[i+j+len(r.from):]
		result = append(result, replaced)

		i += j + 1
	}

	return result
}

func (r Replacement) AllReplaceOnceReverse(s string) []string {
	rr := Replacement{r.to, r.from}
	return rr.AllReplaceOnce(s)
}

type Solver struct {
	molecule     string
	replacements []Replacement
}

func NewSolver(input string) *Solver {
	lines := strings.Split(input, "\n")
	molecule := lines[len(lines)-1]
	lines = lines[:len(lines)-2]
	replacements := []Replacement{}
	for _, line := range lines {
		replacements = append(replacements, parseReplacement(line))
	}
	s := Solver{molecule, replacements}
	return &s
}

func (s *Solver) Part1() int {
	replaced := make(map[string]struct{})
	for _, replacement := range s.replacements {
		for _, r := range replacement.AllReplaceOnce(s.molecule) {
			replaced[r] = struct{}{}
		}
	}
	return len(replaced)
}

func (s *Solver) Part2() int {
	return s.CYK()
}

func (s *Solver) CYK() int {
	revReplacementMap := make(map[string]string)
	for _, r := range s.replacements {
		revReplacementMap[r.to] = r.from
	}

	var recurseAddMetadata func(string, int, *CYKTableMetadata)
	recurseAddMetadata = func(value string, count int, m *CYKTableMetadata) {
		possible := value == "e"
		for k := range revReplacementMap {
			if possible {
				break
			}
			if strings.Contains(k, value) {
				possible = true
			}
		}
		if !possible {
			return
		}

		if count > 0 {
			m.Add(value, count)
		}

		if from, ok := revReplacementMap[value]; ok {
			recurseAddMetadata(from, count+1, m)
		}
	}

	table := make(map[string]*CYKTableMetadata)

	for i := 0; i < len(s.molecule); i++ {
		v := s.molecule[i : i+1]

		m, ok := table[v]
		if ok {
			continue
		}

		m = NewCYKTableMetadata()
		recurseAddMetadata(v, 0, m)
		table[v] = m
	}

	for l := 1; l < len(s.molecule); l++ {
		for x := 0; x < len(s.molecule)-l; x++ {
			y := x + l

			v := s.molecule[x : y+1]
			m, ok := table[v]
			if ok {
				continue
			}

			m = NewCYKTableMetadata()
			recurseAddMetadata(v, 0, m)

			for z := x; z < y; z++ {
				s1 := s.molecule[x : z+1]
				s2 := s.molecule[z+1 : y+1]
				m1 := table[s1]
				m2 := table[s2]
				for _, v1 := range m1.Values() {
					v := v1 + s2
					recurseAddMetadata(v, m1.Get(v1), m)
					for _, v2 := range m2.Values() {
						v := v1 + v2
						recurseAddMetadata(v, m1.Get(v1)+m2.Get(v2), m)
					}
				}
				for _, v2 := range m2.Values() {
					v := s1 + v2
					recurseAddMetadata(v, m2.Get(v2), m)
				}
			}
			table[v] = m
		}
	}

	return table[s.molecule].Get("e")
}

type CYKTableMetadata struct {
	m map[string]int
}

func NewCYKTableMetadata() *CYKTableMetadata {
	m := CYKTableMetadata{
		m: make(map[string]int),
	}
	return &m
}

func (m *CYKTableMetadata) Get(v string) int {
	return m.m[v]
}

func (m *CYKTableMetadata) Values() []string {
	result := []string{}
	for k := range m.m {
		result = append(result, k)
	}
	return result
}

func (m *CYKTableMetadata) Add(value string, count int) {
	if c, ok := m.m[value]; ok && c <= count {
		return
	}
	m.m[value] = count
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
