package main

import "testing"

type Case struct {
	input string
	want  bool
}

func TestRule1(t *testing.T) {
	cases := []Case{
		{"hijklmmn", true},
		{"abbceffg", false},
		{"abbcegjk", false},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			got := Rule1(c.input)
			if got != c.want {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func TestRule2(t *testing.T) {
	cases := []Case{
		{"hijklmmn", false},
		{"abbceffg", true},
		{"abbcegjk", true},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			got := Rule2(c.input)
			if got != c.want {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func TestRule3(t *testing.T) {
	cases := []Case{
		{"hijklmmn", false},
		{"abbceffg", true},
		{"abbcegjk", false},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			got := Rule3(c.input)
			if got != c.want {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	inputs := []string{"abcdefgh", "ghijklmn"}
	wants := []string{"abcdffaa", "ghjaabcc"}
	for i, input := range inputs {
		t.Run(input, func(t *testing.T) {
			s := NewSolver(input)
			got := s.Part1()
			if got != wants[i] {
				t.Errorf("got %s, want %s", got, wants[i])
			}
		})
	}
}
