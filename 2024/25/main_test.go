package main

import "testing"

const (
	input = `#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`
)

func TestPart1(t *testing.T) {
	s := NewSolver(input)
	got := s.Part1()
	want := 3

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
