package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `Time:      7  15   30
Distance:  9  40  200`

func TestPart1(t *testing.T) {
	want := 288
	s := NewSolver(input)
	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 71503
	s := NewSolver(input)
	got := s.Part2()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
