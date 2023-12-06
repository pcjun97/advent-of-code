package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

func TestPart1(t *testing.T) {
	want := 1120
	s := NewSolver(input)
	got := s.Part1(1000)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 689
	s := NewSolver(input)
	got := s.Part2(1000)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
