package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()
	g := NewGrove(input)
	stable := false
	for i := 0; !stable; i++ {
		stable = g.Run()
		if i == 9 {
			fmt.Println(g.EmptyCount())
		}
		if stable {
			fmt.Println(i + 1)
		}
	}
}

type Elf struct {
	Coordinate
	grove *Grove
}

type Grove struct {
	elves      map[Coordinate]*Elf
	checkIndex int
}

func NewGrove(s string) *Grove {
	g := Grove{
		elves:      make(map[Coordinate]*Elf),
		checkIndex: 0,
	}
	for y, line := range strings.Split(s, "\n") {
		for x, c := range line {
			switch c {
			case '#':
				e := Elf{Coordinate{x, y}, &g}
				g.elves[Coordinate{x, y}] = &e
			default:
				continue
			}
		}
	}
	return &g
}

func (g *Grove) EmptyCount() int {
	s := 0
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for c := range g.elves {
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y > maxY {
			maxY = c.Y
		}
		if c.X < minX {
			minX = c.X
		}
		if c.Y < minY {
			minY = c.Y
		}
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, ok := g.elves[Coordinate{x, y}]; !ok {
				s += 1
			}
		}
	}
	return s
}

func (g *Grove) Run() bool {
	stable := true
	p := make(map[Coordinate][]*Elf)
	for _, e := range g.elves {
		c, move := e.Propose()
		if !move {
			continue
		}
		stable = false
		if _, ok := p[c]; !ok {
			p[c] = []*Elf{e}
		} else {
			p[c] = append(p[c], e)
		}
	}
	for c, elves := range p {
		if len(elves) == 1 {
			delete(g.elves, elves[0].Coordinate)
			g.elves[c] = elves[0]
			elves[0].Coordinate = c
		}
	}
	g.checkIndex = (g.checkIndex + 1) % len(checks)
	return stable
}

type Coordinate struct {
	X, Y int
}

type Direction int

const (
	NORTH Direction = iota
	SOUTH
	WEST
	EAST
	NORTHEAST
	NORTHWEST
	SOUTHEAST
	SOUTHWEST
)

var directions [8][2]int = [8][2]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
	{1, -1},
	{-1, -1},
	{1, 1},
	{-1, 1},
}

var checks [4][3]Direction = [4][3]Direction{
	{NORTH, NORTHEAST, NORTHWEST},
	{SOUTH, SOUTHEAST, SOUTHWEST},
	{WEST, NORTHWEST, SOUTHWEST},
	{EAST, NORTHEAST, SOUTHEAST},
}

func (e *Elf) Propose() (Coordinate, bool) {
	alone := true
	for i := 0; i < len(directions); i++ {
		c := Coordinate{
			e.X + directions[i][0],
			e.Y + directions[i][1],
		}
		if _, ok := e.grove.elves[c]; ok {
			alone = false
			break
		}
	}
	if alone {
		return Coordinate{}, false
	}
	for i := 0; i < 4; i++ {
		index := (e.grove.checkIndex + i) % 4
		check := checks[index]
		empty := true
		for _, d := range check {
			c := Coordinate{
				e.X + directions[d][0],
				e.Y + directions[d][1],
			}
			if _, ok := e.grove.elves[c]; ok {
				empty = false
				break
			}
		}
		if empty {
			c := Coordinate{
				e.X + directions[index][0],
				e.Y + directions[index][1],
			}
			return c, true
		}
	}
	return Coordinate{}, false
}
