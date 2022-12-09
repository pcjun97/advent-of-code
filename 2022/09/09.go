package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()

	r1 := NewRope(2)
	r2 := NewRope(10)

	for _, line := range strings.Split(input, "\n") {
		fields := strings.Split(line, " ")
		d := fields[0]
		s, err := strconv.Atoi(fields[1])
		if err != nil {
			panic("error parsing steps: " + fields[1])
		}

		r1.MoveHead(d, s)
		r2.MoveHead(d, s)
	}

	fmt.Println(len(r1.Tail().Visited))
	fmt.Println(len(r2.Tail().Visited))
}

type Rope []*Knot

func NewRope(length int) Rope {
	r := make(Rope, length)
	for i := 0; i < length; i++ {
		r[i] = NewKnot(0, 0)
	}
	return r
}

func (r Rope) MoveHead(direction string, steps int) {
	for steps > 0 {
		r[0].Move(direction)
		for i := 1; i < len(r); i++ {
			r[i].Follow(r[i-1])
		}
		steps--
	}
}

func (r Rope) Tail() *Knot {
	return r[len(r)-1]
}

type Knot struct {
	X, Y    int
	Visited map[string]bool
}

func NewKnot(x int, y int) *Knot {
	k := Knot{
		X:       x,
		Y:       y,
		Visited: make(map[string]bool),
	}

	k.Visited[k.CoordinateString()] = true
	return &k
}

func (k *Knot) Move(direction string) {
	switch direction {
	case "R":
		k.X += 1
	case "L":
		k.X -= 1
	case "U":
		k.Y += 1
	case "D":
		k.Y -= 1
	}

	k.Visited[k.CoordinateString()] = true
}

func (k *Knot) Distance(other *Knot) (int, int) {
	return other.X - k.X, other.Y - k.Y
}

func (k *Knot) Follow(other *Knot) {
	dX, dY := k.Distance(other)
	if abs(dX) < 2 && abs(dY) < 2 {
		return
	}

	if dX != 0 {
		k.X += dX / abs(dX)
	}

	if dY != 0 {
		k.Y += dY / abs(dY)
	}

	k.Visited[k.CoordinateString()] = true
}

func (k *Knot) CoordinateString() string {
	return fmt.Sprintf("%d,%d", k.X, k.Y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
