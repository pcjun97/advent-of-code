package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input string
	want  int
}

const input1 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

const input2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

const input3 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestPart1(t *testing.T) {
	cases := []Case{
		{input1, 2},
		{input2, 6},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("input %d", i+1), func(t *testing.T) {
			s := NewSolver(c.input)
			got := s.Part1()
			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	want := 6
	s := NewSolver(input3)
	got := s.Part2()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
