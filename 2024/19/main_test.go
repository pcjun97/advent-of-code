package main

import "testing"

const input = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

func TestPart1(t *testing.T) {
	s := NewSolver(input)
	got := s.Part1()
	want := 6

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	s := NewSolver(input)
	got := s.Part2()
	want := 16

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
