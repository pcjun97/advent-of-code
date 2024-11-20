package main

import (
	"testing"
)

type Case struct {
	input string
	want  int
}

func TestPart1(t *testing.T) {
	cases := []Case{
		{"R2, L3", 5},
		{"R2, R2, R2", 2},
		{"R5, L5, R5, R3", 12},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			s := NewSolver(c.input)
			got := s.Part1()

			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []Case{
		{"R8, R4, R4, R8", 4},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			s := NewSolver(c.input)
			got := s.Part2()

			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}
