package main

import "testing"

type Case struct {
	input string
	want  int
}

const input1 = `turn on 0,0 through 999,999
toggle 0,0 through 999,0
turn off 499,499 through 500,500`

const input2 = `turn on 0,0 through 0,0
toggle 0,0 through 999,999`

func TestPart1(t *testing.T) {
	s, err := NewSolver(input1)
	if err != nil {
		t.Fatal(err)
	}

	want := 1000000 - 1000 - 4
	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	s, err := NewSolver(input2)
	if err != nil {
		t.Fatal(err)
	}

	want := 1 + 2000000
	got := s.Part2()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
