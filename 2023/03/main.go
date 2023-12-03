package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Coordinate struct {
	X, Y int
}

type Part struct {
	coordinates []Coordinate
	number      int
}

type Symbol struct {
	c Coordinate
	s string
}

type Node struct {
	part   *Part
	symbol *Symbol
}

func parseSchematic(s string) (map[Coordinate]*Node, []*Symbol, []*Part) {
	schematic := make(map[Coordinate]*Node)
	symbols := []*Symbol{}
	parts := []*Part{}

	for y, line := range strings.Split(s, "\n") {
		rPart := regexp.MustCompile(`\d+`)
		mPart := rPart.FindAllStringIndex(line, -1)
		for _, m := range mPart {
			n := Node{}
			coordinates := []Coordinate{}
			for x := m[0]; x < m[1]; x++ {
				c := Coordinate{x, y}
				coordinates = append(coordinates, c)
				schematic[c] = &n
			}

			v, _ := strconv.Atoi(line[m[0]:m[1]])
			p := Part{coordinates, v}
			n.part = &p
			parts = append(parts, &p)
		}

		rSymbol := regexp.MustCompile(`[^\d.]`)
		mSymbol := rSymbol.FindAllStringIndex(line, -1)
		for _, m := range mSymbol {
			c := Coordinate{m[0], y}
			symbol := Symbol{c, line[m[0]:m[1]]}
			n := Node{
				symbol: &symbol,
			}
			schematic[c] = &n
			symbols = append(symbols, &symbol)
		}
	}

	return schematic, symbols, parts
}

type Solver struct {
	schematic map[Coordinate]*Node
	symbols   []*Symbol
	parts     []*Part
}

func NewSolver(input string) *Solver {
	schematic, symbols, parts := parseSchematic(input)
	s := Solver{schematic, symbols, parts}
	return &s
}

func (s *Solver) Part1() int {
	set := make(map[*Part]struct{})
	for _, symbol := range s.symbols {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				c := Coordinate{symbol.c.X + i, symbol.c.Y + j}
				node, ok := s.schematic[c]
				if !ok || node.part == nil {
					continue
				}
				set[node.part] = struct{}{}
			}
		}
	}

	sum := 0
	for part := range set {
		sum += part.number
	}

	return sum
}

func (s *Solver) Part2() int {
	sum := 0

	for _, symbol := range s.symbols {
		if symbol.s != "*" {
			continue
		}

		adj := make(map[*Part]struct{})
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				c := Coordinate{symbol.c.X + i, symbol.c.Y + j}
				node, ok := s.schematic[c]
				if !ok || node.part == nil {
					continue
				}
				adj[node.part] = struct{}{}
			}
		}

		if len(adj) != 2 {
			continue
		}

		ratio := 1
		for part := range adj {
			ratio *= part.number
		}
		sum += ratio
	}

	return sum
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
