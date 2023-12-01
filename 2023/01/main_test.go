package main

import "testing"

type Case struct {
	input string
	want  int
}

const (
	input1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	input2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
)

func TestPart1(t *testing.T) {
	s := NewSolver(input1)
	got := s.Part1()
	want := 142

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	s := NewSolver(input2)
	got := s.Part2()
	want := 281

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
