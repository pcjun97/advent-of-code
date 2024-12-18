package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

func TestPart1(t *testing.T) {
	s := NewSolver(input, 6)
	got := s.Part1(12)
	want := 22

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	s := NewSolver(input, 6)
	got := s.Part2()
	want := "6,1"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
