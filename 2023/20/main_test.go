package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input string
	want  int
}

const (
	input1 = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

	input2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`
)

func TestPart1(t *testing.T) {
	cases := []Case{
		{input1, 32000000},
		{input2, 11687500},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("input %d", i+1), func(t *testing.T) {
			s := NewSolver(c.input)
			got := s.Part1()
			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
}
