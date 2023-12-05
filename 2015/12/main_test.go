package main

import "testing"

type Case struct {
	input string
	want  int
}

func TestPart1(t *testing.T) {
	cases := []Case{
		{"[1,2,3]", 6},
		{"{\"a\":2,\"b\":4}", 6},
		{"[[[3]]]", 3},
		{"{\"a\":{\"b\":4},\"c\":-1}", 3},
		{"{\"a\":[-1,1]}", 0},
		{"[-1,{\"a\":1}]", 0},
		{"[]", 0},
		{"{}", 0},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			s := NewSolver(c.input)
			got := s.Part1()
			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []Case{
		{"[1,2,3]", 6},
		{"[1,{\"c\":\"red\",\"b\":2},3]", 4},
		{"{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}", 0},
		{"[1,\"red\",5]", 6},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			s := NewSolver(c.input)
			got := s.Part2()
			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}
