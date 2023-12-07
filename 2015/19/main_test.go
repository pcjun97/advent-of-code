package main

import "testing"

type Case struct {
	input string
	want  int
}

const input1 = `H => HO
H => OH
O => HH`

const input2 = `e => H
e => O
H => HO
H => OH
O => HH`

func TestPart1(t *testing.T) {
	cases := []Case{
		{"HOH", 4},
		{"HOHOHO", 7},
	}
	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			s := NewSolver(input1 + "\n\n" + c.input)
			got := s.Part1()
			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []Case{
		{"HOH", 3},
		{"HOHOHO", 6},
	}
	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			s := NewSolver(input2 + "\n\n" + c.input)
			got := s.Part2()
			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}
