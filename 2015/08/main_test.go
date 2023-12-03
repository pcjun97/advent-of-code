package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `""
"abc"
"aaa\"aaa"
"\x27"`

func TestPart1(t *testing.T) {
	s := NewSolver(input)
	want := 12
	got := s.Part1()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 19

	s := NewSolver(input)
	got := s.Part2()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
