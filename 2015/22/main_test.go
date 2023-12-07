package main

import "testing"

type Case struct {
	input string
	want  int
}

const input1 = `Hit Points: 13
Damage: 8`

const input2 = `Hit Points: 14
Damage: 8`

func TestPart1(t *testing.T) {
	cases := []Case{
		{input1, 173 + 53},
		{input2, 229 + 113 + 73 + 173 + 53},
	}
	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			s := NewSolver(c.input)
			got := s.Part1(10, 250)
			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}
