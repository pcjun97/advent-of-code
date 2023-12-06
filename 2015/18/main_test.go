package main

import "testing"

const input = `.#.#.#
...##.
#....#
..#...
#.#..#
####..`

func TestPart1(t *testing.T) {
	want := 4
	s := NewSolver(input)
	got := s.Part1(4)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 17
	s := NewSolver(input)
	got := s.Part2(5)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
