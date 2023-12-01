package main

import "testing"

type Case struct {
	input string
	want  int
}

var cases1 = []Case{
	{">", 2},
	{"^>v<", 4},
	{"^v^v^v^v^v", 2},
}

var cases2 = []Case{
	{"^v", 3},
	{"^>v<", 3},
	{"^v^v^v^v^v", 11},
}

func TestPart1(t *testing.T) {
	for _, c := range cases1 {
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
	for _, c := range cases2 {
		t.Run(c.input, func(t *testing.T) {
			s := NewSolver(c.input)
			got := s.Part2()
			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}
