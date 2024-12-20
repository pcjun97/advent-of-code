package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

const (
	Wall  = '#'
	Track = '.'
	Start = 'S'
	End   = 'E'
)

func ParseGrid(s string) *aoc.Grid {
	g := aoc.NewGrid()

	for y, line := range strings.Split(s, "\n") {
		for x, r := range line {
			c := aoc.NewCoordinate(x, y)
			g.Add(aoc.NewNode(c, int(r)))
		}
	}

	return g
}

type Solver struct {
	*aoc.Grid
}

func NewSolver(input string) *Solver {
	s := Solver{ParseGrid(input)}
	return &s
}

func (s *Solver) GetStart() aoc.Coordinate {
	for _, n := range s.Nodes() {
		if n.Value() == Start {
			return n.Coordinate
		}
	}
	return aoc.NewCoordinate(-1, -1)
}

func (s *Solver) GetEnd() aoc.Coordinate {
	for _, n := range s.Nodes() {
		if n.Value() == End {
			return n.Coordinate
		}
	}
	return aoc.NewCoordinate(-1, -1)
}

type Cheat struct {
	Start, End aoc.Coordinate
}

func (s *Solver) BestCheats(cheatperiod, minsaving int) []Cheat {
	besttime := make(map[aoc.Coordinate]int)

	start := s.GetStart()
	end := s.GetEnd()

	besttime[end] = 0

	t := 0
	cur := end
	for {
		besttime[cur] = t

		if cur == start {
			break
		}

		for _, c := range cur.Neighbors4Way() {
			n := s.Get(c)
			if n == nil || n.Value() == Wall {
				continue
			}

			if _, ok := besttime[c]; ok {
				continue
			}

			cur = c
			break
		}

		t++
	}

	cheats := []Cheat{}

	for from, t := range besttime {
		for x := -cheatperiod; x <= cheatperiod; x++ {
			for y := -cheatperiod; y <= cheatperiod; y++ {
				to := aoc.NewCoordinate(from.X+x, from.Y+y)

				if from.ManhattanDistance(to) > cheatperiod {
					continue
				}

				if _, ok := besttime[to]; !ok {
					continue
				}

				if t-(besttime[to]+from.ManhattanDistance(to)) >= minsaving {
					cheats = append(cheats, Cheat{from, to})
				}
			}
		}
	}

	return cheats
}

func (s *Solver) Part1(minsaving int) int {
	return len(s.BestCheats(2, minsaving))
}

func (s *Solver) Part2(minsaving int) int {
	return len(s.BestCheats(20, minsaving))
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(100), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(100), time.Since(start).String())
}
