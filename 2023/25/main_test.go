package main

import "testing"

const input = `jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr`

func TestPart1(t *testing.T) {
	want := 54
	s := NewSolver(input)
	got := s.Part1()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
