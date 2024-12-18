package main

import "testing"

const input = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

func TestPart1(t *testing.T) {
	s := NewSolver(input)
	got := s.Part1()
	want := "4,6,3,5,6,3,5,2,1,0"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
