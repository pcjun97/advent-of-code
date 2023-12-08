package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `1
2
3
4
5
7
8
9
10
11`

func TestPart1(t *testing.T) {
	want := 99
	s := NewSolver(input)
	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 44
	s := NewSolver(input)
	got := s.Part2()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
