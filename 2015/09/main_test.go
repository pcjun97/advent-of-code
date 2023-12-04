package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`

func TestPart1(t *testing.T) {
	want := 605

	s := NewSolver(input)
	got := s.Part1()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 982

	s := NewSolver(input)
	got := s.Part2()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
