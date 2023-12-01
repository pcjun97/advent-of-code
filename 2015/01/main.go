package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Solver struct {
	instructions string
}

func NewSolver(instructions string) *Solver {
	s := Solver{
		instructions: instructions,
	}
	return &s
}

func (s *Solver) FinalLevel() (int, error) {
	finalLevel := 0
	for _, r := range s.instructions {
		switch r {
		case '(':
			finalLevel++

		case ')':
			finalLevel--

		default:
			return 0, errors.New("unknown instruction")
		}
	}

	return finalLevel, nil
}

func (s *Solver) Part2() (int, error) {
	position := 0
	level := 0

	for _, r := range s.instructions {
		position++

		switch r {
		case '(':
			level++

		case ')':
			level--
			if level == -1 {
				return position, nil
			}

		default:
			return 0, errors.New("unknown instruction")
		}
	}

	return 0, errors.New("never enter basement")
}

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.TrimSuffix(string(input), "\n")
	s := NewSolver(instructions)

	finalLevel, err := s.FinalLevel()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", finalLevel)

	pos, err := s.Part2()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", pos)
}
