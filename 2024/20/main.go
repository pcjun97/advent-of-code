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

	end := s.GetEnd()

	besttime[end] = 0
	tovisit := []aoc.Coordinate{end}

	for len(tovisit) > 0 {
		i := 0
		for j, v := range tovisit {
			if besttime[v] < besttime[tovisit[i]] {
				i = j
			}
		}
		visiting := tovisit[i]
		tovisit = append(tovisit[:i], tovisit[i+1:]...)

		for _, c := range visiting.Neighbors4Way() {
			n := s.Get(c)
			if n == nil || n.Value() == Wall {
				continue
			}

			if _, ok := besttime[c]; ok {
				continue
			}
			besttime[c] = besttime[visiting] + 1

			tovisit = append(tovisit, c)
		}
	}

	tovisit = []aoc.Coordinate{s.GetStart()}
	paths := make(map[aoc.Coordinate]struct{})

	for len(tovisit) > 0 {
		visiting := tovisit[0]
		tovisit = tovisit[1:]

		if _, ok := paths[visiting]; ok {
			continue
		}
		paths[visiting] = struct{}{}

		for _, c := range visiting.Neighbors4Way() {
			if _, ok := besttime[c]; !ok {
				continue
			}
			if besttime[visiting]-besttime[c] != 1 {
				continue
			}
			tovisit = append(tovisit, c)
		}
	}

	cheats := []Cheat{}

	for from := range paths {
		for x := -cheatperiod; x <= cheatperiod; x++ {
			for y := -cheatperiod; y <= cheatperiod; y++ {
				to := aoc.NewCoordinate(from.X+x, from.Y+y)

				if from.ManhattanDistance(to) > cheatperiod {
					continue
				}

				if _, ok := besttime[to]; !ok {
					continue
				}

				if besttime[from]-(besttime[to]+from.ManhattanDistance(to)) >= minsaving {
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
