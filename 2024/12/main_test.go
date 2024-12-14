package main

import "testing"

type Case struct {
	input string
	want  int
}

const (
	input1 = `AAAA
BBCD
BBCC
EEEC`

	input2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

	input3 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

	input4 = `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

	input5 = `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`
)

func TestPart1(t *testing.T) {
	cases := []Case{
		{input1, 140},
		{input2, 772},
		{input3, 1930},
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
		{input1, 80},
		{input2, 436},
		{input3, 1206},
		{input4, 236},
		{input5, 368},
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
