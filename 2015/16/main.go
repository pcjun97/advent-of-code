package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

const (
	Unknown Attribute = iota
	Children
	Cats
	Samoyeds
	Pomeranians
	Akitas
	Vizslas
	Goldfish
	Trees
	Cars
	Perfumes
)

type Attribute int

func AttributeFromString(s string) Attribute {
	switch s {
	case "children":
		return Children
	case "cats":
		return Cats
	case "samoyeds":
		return Samoyeds
	case "pomeranians":
		return Pomeranians
	case "akitas":
		return Akitas
	case "vizslas":
		return Vizslas
	case "goldfish":
		return Goldfish
	case "trees":
		return Trees
	case "cars":
		return Cars
	case "perfumes":
		return Perfumes
	default:
		return Unknown
	}
}

type Sue struct {
	id         int
	attributes map[Attribute]int
}

func parseSue(s string) Sue {
	attributes := make(map[Attribute]int)

	r := regexp.MustCompile(`Sue (\d*): (.*)$`)
	m := r.FindStringSubmatch(s)
	id, _ := strconv.Atoi(m[1])
	for _, attr := range strings.Split(m[2], ", ") {
		r := regexp.MustCompile(`^(.*): (.*)$`)
		m := r.FindStringSubmatch(attr)
		a := AttributeFromString(m[1])
		c, _ := strconv.Atoi(m[2])
		attributes[a] = c
	}

	sue := Sue{id, attributes}
	return sue
}

type Solver struct {
	aunts  []Sue
	target Sue
}

func NewSolver(input string) *Solver {
	aunts := []Sue{}
	for _, line := range strings.Split(input, "\n") {
		aunts = append(aunts, parseSue(line))
	}

	target := parseSue("Sue 0: children: 3, cats: 7, samoyeds: 2, pomeranians: 3, akitas: 0, vizslas: 0, goldfish: 5, trees: 3, cars: 2, perfumes: 1")

	s := Solver{aunts, target}
	return &s
}

func (s *Solver) Part1() int {
	var x Sue
	for _, aunt := range s.aunts {
		found := true
		for k, v := range s.target.attributes {
			vv, ok := aunt.attributes[k]
			if ok && vv != v {
				found = false
				break
			}
		}
		if found {
			x = aunt
			break
		}
	}

	return x.id
}

func (s *Solver) Part2() int {
	var x Sue
	for _, aunt := range s.aunts {
		found := true
		for k, v := range s.target.attributes {
			vv, ok := aunt.attributes[k]
			if !ok {
				continue
			}

			switch k {
			case Cats, Trees:
				if vv <= v {
					found = false
				}
			case Pomeranians, Goldfish:
				if vv >= v {
					found = false
				}
			default:
				if v != vv {
					found = false
				}
			}

			if !found {
				break
			}
		}

		if found {
			x = aunt
			break
		}
	}

	return x.id
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
