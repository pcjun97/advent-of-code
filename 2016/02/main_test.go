package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `ULL
RRDDD
LURDL
UUUUD`

func TestPart1(t *testing.T) {
	s := NewSolver(input)

	got := s.Part1()
	want := "1985"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	s := NewSolver(input)

	got := s.Part2()
	want := "5DB3"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
