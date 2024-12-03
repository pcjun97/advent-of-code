package main

import (
	"testing"
)

const input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPart1(t *testing.T) {
	want := 2
	s := NewSolver(input)
	got := s.Part1()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 4
	s := NewSolver(input)
	got := s.Part2()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
