package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func TestPart1(t *testing.T) {
	want := 405
	s := NewSolver(input)
	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 400
	s := NewSolver(input)
	got := s.Part2()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
