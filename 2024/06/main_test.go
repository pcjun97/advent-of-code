package main

import "testing"

const input = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPart1(t *testing.T) {
	want := 41

	s := NewSolver(input)
	got := s.Part1()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 6

	s := NewSolver(input)
	got := s.Part2()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
