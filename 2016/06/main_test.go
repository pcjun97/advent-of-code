package main

import "testing"

const input = `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

func TestPart1(t *testing.T) {
	want := "easter"

	s := NewSolver(input)
	got := s.Part1()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := "advent"

	s := NewSolver(input)
	got := s.Part2()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
