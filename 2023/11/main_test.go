package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestPart1(t *testing.T) {
	want := 374
	s := NewSolver(input)
	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 1030
	s := NewSolver(input)
	got := s.Part2(10)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	want = 8410
	got = s.Part2(100)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
