package main

import "testing"

func TestPart1(t *testing.T) {
	cases := []struct {
		instructions string
		want         int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, c := range cases {
		t.Run(c.instructions, func(t *testing.T) {
			s := NewSolver(c.instructions)
			got, err := s.FinalLevel()

			if err != nil {
				t.Errorf("expected no error: %v", err)
			}

			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		instructions string
		want         int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, c := range cases {
		t.Run(c.instructions, func(t *testing.T) {
			s := NewSolver(c.instructions)
			got, err := s.Part2()

			if err != nil {
				t.Errorf("expected no error: %v", err)
			}

			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}
