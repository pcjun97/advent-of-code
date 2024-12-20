package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

func TestPart1(t *testing.T) {
	s := NewSolver(input)
	got := s.Part1(2)
	want := 44

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	s := NewSolver(input)
	got := s.Part2(50)
	want := 285

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
