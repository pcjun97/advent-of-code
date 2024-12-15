package main

import "testing"

const input = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func TestPart1(t *testing.T) {
	s := NewSolver(input, 7, 11)
	got := s.Part1()
	want := 12

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
