package main

import "testing"

const input = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestPart1(t *testing.T) {
	s := NewSolver(input)
	got := s.Part1()

	want := 36
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	s := NewSolver(input)
	got := s.Part2()

	want := 81
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
