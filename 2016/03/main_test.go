package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `3 5 3
5 10 25
3 25 23`

func TestPart1(t *testing.T) {
	want := 2

	s := NewSolver(input)

	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 2

	s := NewSolver(input)

	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
