package main

import "testing"

type Case struct {
	input string
	want  int
}

func TestPart1(t *testing.T) {
	input := `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`

	want := 1514

	s := NewSolver(input)

	got := s.Part1()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDecrypt(t *testing.T) {
	r := Room{
		Name:     "qzmt-zixmtkozy-ivhz",
		SectorID: 343,
	}

	want := "very encrypted name"
	got := r.Decrypt()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
