package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input string
	want  int
}

const input = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func TestPart1(t *testing.T) {
	want := 16
	s := NewSolver(input)
	got := s.Part1(6)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	cases := [][2]int{
		{6, 16},
		{10, 50},
		{50, 1594},
		{100, 6536},
		{500, 167004},
		{1000, 668697},
		{5000, 16733044},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d", c[0]), func(t *testing.T) {
			want := c[1]
			s := NewSolver(input)
			got := s.Part2(c[0])
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
