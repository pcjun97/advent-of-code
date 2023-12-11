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
	input1 = `.....
.S-7.
.|.|.
.L-J.
.....`

	input2 = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

	input3 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

	input4 = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

	input5 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`
)

func TestPart1(t *testing.T) {
	cases := []Case{
		{input1, 4},
		{input2, 8},
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
	cases := []Case{
		{input3, 4},
		{input4, 8},
		{input5, 10},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("input %d", i+3), func(t *testing.T) {
			s := NewSolver(c.input)
			got := s.Part2()
			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}
