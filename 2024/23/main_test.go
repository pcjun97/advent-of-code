package main

import "testing"

const (
	input1 = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`
)

func TestPart1(t *testing.T) {
	s := NewSolver(input1)
	got := s.Part1()
	want := 7

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	s := NewSolver(input1)
	got := s.Part2()
	want := "co,de,ka,ta"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
