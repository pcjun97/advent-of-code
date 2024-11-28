package main

import "testing"

const input = "abc"

func TestPart1(t *testing.T) {
	want := "18f47a30"
	s := NewSolver(input)
	got := s.Part1()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := "05ace8e3"
	s := NewSolver(input)
	got := s.Part2()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
