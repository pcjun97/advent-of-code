package main

import "testing"

const input = `1`

func TestLookAndSay(t *testing.T) {
	want := "312211"

	got := input
	for i := 0; i < 5; i++ {
		got = LookAndSay(got)
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
