package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input string
	want  int
}

func TestMix(t *testing.T) {
	got := Mix(42, 15)
	want := 37

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPrune(t *testing.T) {
	got := Prune(100000000)
	want := 16113920

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestNextSecretNumber(t *testing.T) {
	n := 123

	wants := []int{
		15887950,
		16495136,
		527345,
		704524,
		1553684,
		12683156,
		11100544,
		12249484,
		7753432,
		5908254,
	}

	for _, want := range wants {
		n = NextSecretNumber(n)

		if n != want {
			t.Errorf("got %d, want %d", n, want)
		}
	}
}

func Test2000thSecretNumber(t *testing.T) {
	cases := [][2]int{
		{1, 8685429},
		{10, 4700978},
		{100, 15273692},
		{2024, 8667524},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d", c[0]), func(t *testing.T) {
			got := NthSecretNumber(c[0], 2000)
			want := c[1]

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	input := `1
10
100
2024`

	s := NewSolver(input)
	got := s.Part1()
	want := 37327623

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `1
2
3
2024`

	s := NewSolver(input)
	got := s.Part2()
	want := 23

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
