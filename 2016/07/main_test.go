package main

import (
	"strings"
	"testing"
)

type Case struct {
	input string
	want  bool
}

var inputs = []string{
	"abba[mnop]qrst",
	"abcd[bddb]xyyx",
	"aaaa[qwer]tyui",
	"ioxxoj[asdfgh]zxcvbn",
}

func TestSplit(t *testing.T) {
	wants := []IPv7{
		{[]IPv7Section{
			{"abba", false},
			{"mnop", true},
			{"qrst", false},
		}},
		{[]IPv7Section{
			{"abcd", false},
			{"bddb", true},
			{"xyyx", false},
		}},
		{[]IPv7Section{
			{"aaaa", false},
			{"qwer", true},
			{"tyui", false},
		}},
		{[]IPv7Section{
			{"ioxxoj", false},
			{"asdfgh", true},
			{"zxcvbn", false},
		}},
	}

	for i, c := range inputs {
		t.Run(c, func(t *testing.T) {
			ip := ParseIPv7(c)

			for j, s := range ip.Sections {
				if s.Hypernet != wants[i].Sections[j].Hypernet {
					t.Errorf("got %v, want %v", s.Hypernet, wants[i].Sections[j].Hypernet)
				}

				if s.Value != wants[i].Sections[j].Value {
					t.Errorf("got %s, want %s", s.Value, wants[i].Sections[j].Value)
				}
			}
		})
	}
}

func TestTLS(t *testing.T) {
	wants := []bool{
		true,
		false,
		false,
		true,
	}

	for i, c := range inputs {
		t.Run(c, func(t *testing.T) {
			ip := ParseIPv7(c)
			got := ip.TLS()

			if got != wants[i] {
				t.Errorf("got %v, want %v", got, wants[i])
			}
		})
	}
}

func TestPart1(t *testing.T) {
	input := strings.Join(inputs, "\n")
	s := NewSolver(input)
	got := s.Part1()
	want := 2

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
}
