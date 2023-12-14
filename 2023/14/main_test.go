package main

import "testing"

const input = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func TestPart1(t *testing.T) {
	want := 136
	s := NewSolver(input)
	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 64
	s := NewSolver(input)
	got := s.Part2()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
