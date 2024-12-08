package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1(t *testing.T) {
	want := 11

	s := NewSolver(input)
	got := s.Part1()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 31

	s := NewSolver(input)
	got := s.Part2()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
