package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `20
15
10
5
5`

func TestPart1(t *testing.T) {
	want := 4
	s := NewSolver(input)
	got := s.Part1(25)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 3
	s := NewSolver(input)
	got := s.Part2(25)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
