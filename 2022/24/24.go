package main

import (
	"fmt"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()
	v := NewValley(input)
	t := v.BFS(v.Start, v.End, 0)
	t = v.BFS(v.End, v.Start, t)
	fmt.Println(v.BFS(v.Start, v.End, t))
}

type Valley struct {
	blizzards     []map[Coordinate][]Direction
	Start, End    Coordinate
	Height, Width int
	cache         map[State]bool
}

func NewValley(s string) *Valley {
	lines := strings.Split(s, "\n")

	var start, end int
	for x, c := range lines[0][1 : len(lines[0])-1] {
		if c == '.' {
			start = x
		}
	}
	for y, c := range lines[len(lines)-1][1 : len(lines[0])-1] {
		if c == '.' {
			end = y
		}
	}

	height := len(lines) - 2
	width := len(lines[0]) - 2

	blizzards := make(map[Coordinate][]Direction)
	for y, line := range lines[1 : len(lines)-1] {
		var d Direction
		for x, c := range line[1 : len(line)-1] {
			switch c {
			case '>':
				d = RIGHT
			case 'v':
				d = DOWN
			case '<':
				d = LEFT
			case '^':
				d = UP
			default:
				continue
			}
			blizzards[Coordinate{x, y}] = []Direction{d}
		}
	}

	valley := Valley{
		blizzards: []map[Coordinate][]Direction{blizzards},
		Start:     Coordinate{start, -1},
		End:       Coordinate{end, height},
		Height:    height,
		Width:     width,
	}

	return &valley
}

func (v *Valley) BFS(start, end Coordinate, t int) int {
	s0 := State{
		Coordinate: start,
		minutes:    t,
	}
	states := []State{s0}
	v.cache = make(map[State]bool)
	for len(states) > 0 {
		s := states[0]
		states = states[1:]
		if _, ok := v.cache[s]; ok {
			continue
		}
		v.cache[s] = true

		for s.minutes+1 >= len(v.blizzards) {
			v.blizzards = append(v.blizzards, v.Next(v.blizzards[len(v.blizzards)-1]))
		}

		blizzards := v.blizzards[s.minutes+1]
		for _, c := range v.Explore(s.Coordinate, blizzards) {
			if c == end {
				return s.minutes + 1
			}
			ss := State{
				Coordinate: c,
				minutes:    s.minutes + 1,
			}
			states = append(states, ss)
		}
	}
	return 0
}

func (v *Valley) Next(m map[Coordinate][]Direction) map[Coordinate][]Direction {
	next := make(map[Coordinate][]Direction)
	for c, dirs := range m {
		for _, d := range dirs {
			x := (c.X + directions[d][0] + v.Width) % v.Width
			y := (c.Y + directions[d][1] + v.Height) % v.Height
			cc := Coordinate{x, y}

			if _, ok := next[cc]; ok {
				next[cc] = append(next[cc], d)
			} else {
				next[cc] = []Direction{d}
			}
		}
	}
	return next
}

func (v *Valley) Explore(c Coordinate, blizzards map[Coordinate][]Direction) []Coordinate {
	next := []Coordinate{}
	for d := 0; d < 5; d++ {
		cc := c
		if d < 4 {
			cc.X += directions[d][0]
			cc.Y += directions[d][1]
		}
		if !v.Contains(cc) {
			continue
		}
		if b, ok := blizzards[cc]; !ok || len(b) == 0 {
			next = append(next, cc)
		}
	}
	return next
}

func (v *Valley) Contains(c Coordinate) bool {
	if c == v.Start || c == v.End {
		return true
	}

	if c.X >= 0 && c.X < v.Width && c.Y >= 0 && c.Y < v.Height {
		return true
	}

	return false
}

type State struct {
	Coordinate
	minutes int
}

type Coordinate struct {
	X, Y int
}

type Direction int

const (
	RIGHT Direction = iota
	DOWN
	LEFT
	UP
)

var directions [4][2]int = [4][2]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}
