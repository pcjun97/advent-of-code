package main

import "testing"

type Case struct {
	input string
	want  int
}

const (
	input1 = `ugknbfddgicrmopn
aaa
jchzalrnumimnmhp
haegwjzuvuyypxyu
dvszwmarrgswjxmb`

	input2 = `qjhvhtzxzqqjkmpb
xxyxx
uurcxstgmygtbstg
ieodomkazucvgmuy`
)

func TestPart1(t *testing.T) {
	s := NewSolver(input1)
	want := 2
	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	s := NewSolver(input2)
	want := 2
	got := s.Part2()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
