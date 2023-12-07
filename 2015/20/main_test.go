package main

import "testing"

func TestPart1(t *testing.T) {
	want := 6
	s := NewSolver("100")
	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
