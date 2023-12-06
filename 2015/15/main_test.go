package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`

func TestPart1(t *testing.T) {
	want := 62842880
	s := NewSolver(input)
	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 57600000
	s := NewSolver(input)
	got := s.Part2()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
