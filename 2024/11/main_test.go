package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `125 17`

func TestPart1(t *testing.T) {
	s := NewSolver(input)
	got := s.Part1()
	want := 55312

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
