package main

import "testing"

type Case struct {
	input string
	want  int
}

var cases = []Case{
	{"abcdef", 609043},
	{"pqrstuv", 1048970},
}

func TestPart1(t *testing.T) {
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
