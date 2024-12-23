package main

import "testing"

const (
	input = `029A
980A
179A
456A
379A`
)

func TestPart1(t *testing.T) {
	s := NewSolver(input)
	got := s.Part1()
	want := 126384

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
