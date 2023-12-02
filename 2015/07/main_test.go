package main

import "testing"

type Case struct {
	input string
	want  int
}

const input1 = `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`

var cases = []Case{
	{"d", 72},
	{"e", 507},
	{"f", 492},
	{"g", 114},
	{"h", 65412},
	{"i", 65079},
	{"x", 123},
	{"y", 456},
}

func TestValueOf(t *testing.T) {
	s := NewSolver(input1)

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			got := s.valueOf(c.input)
			if got != uint16(c.want) {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}
