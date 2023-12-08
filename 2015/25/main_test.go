package main

import "testing"

const input = `To continue, please consult the code grid in the manual.  Enter the code at row 2, column 2.`

func TestPart1(t *testing.T) {
	want := 21629792
	s := NewSolver(input)
	got := s.Part1(20151125, 1, 1)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
