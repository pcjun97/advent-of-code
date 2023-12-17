package main

import "testing"

type Case struct {
	input string
	want  int
}

const input = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func TestPart1(t *testing.T) {
	want := 102
	s := NewSolver(input)
	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 94
	s := NewSolver(input)
	got := s.Part2()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	input2 := `111111111111
999999999991
999999999991
999999999991
999999999991`
	want = 71
	s = NewSolver(input2)
	got = s.Part2()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
