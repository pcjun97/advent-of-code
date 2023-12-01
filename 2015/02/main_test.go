package main

import "testing"

const (
	input = `2x3x4
1x1x10`
)

func TestPart1(t *testing.T) {
	s, err := NewSolver(input)
	if err != nil {
		t.Error(err)
	}

	want := 101
	got := s.Part1()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	s, err := NewSolver(input)
	if err != nil {
		t.Error(err)
	}

	want := 48
	got := s.Part2()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
