package main

import (
	"fmt"

	"github.com/pcjun97/advent-of-code/aoc"
)

type House struct {
	X int
	Y int
}

type Santa struct {
	current House
	Visited map[House]struct{}
}

func NewSanta() *Santa {
	s := Santa{
		current: House{0, 0},
		Visited: make(map[House]struct{}),
	}
	s.Visited[s.current] = struct{}{}
	return &s
}

func (s *Santa) Next(direction rune) {
	switch direction {
	case '>':
		s.current.X += 1
	case '<':
		s.current.X -= 1
	case '^':
		s.current.Y += 1
	case 'v':
		s.current.Y -= 1
	}
	s.Visited[s.current] = struct{}{}
}

type Solver struct {
	directions string
}

func NewSolver(input string) *Solver {
	s := Solver{input}
	return &s
}

func (s *Solver) Part1() int {
	santa := NewSanta()
	for _, r := range s.directions {
		santa.Next(r)
	}

	return len(santa.Visited)
}

func (s *Solver) Part2() int {
	var santas [2]*Santa
	for i := range santas {
		santas[i] = NewSanta()
	}

	for i, r := range s.directions {
		santas[i%len(santas)].Next(r)
	}

	visited := make(map[House]struct{})
	for _, santa := range santas {
		for k := range santa.Visited {
			visited[k] = struct{}{}
		}
	}

	return len(visited)
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
