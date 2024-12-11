package main

import "testing"

const input = `2333133121414131402`

func TestPart1(t *testing.T) {
	s := NewSolver(input)
	got := s.Part1()

	want := 1928
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	s := NewSolver(input)
	got := s.Part2()

	want := 2858
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
