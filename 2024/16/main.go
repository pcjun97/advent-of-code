package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

const (
	Empty int = iota
	Wall
	Start
	End
)

const (
	North int = iota
	East
	South
	West
)

func ParseGrid(s string) *aoc.Grid {
	g := aoc.NewGrid()

	for y, line := range strings.Split(s, "\n") {
		for x, r := range line {
			c := aoc.NewCoordinate(x, y)

			v := Empty
			switch r {
			case '#':
				v = Wall
			case 'S':
				v = Start
			case 'E':
				v = End
			}

			g.Add(aoc.NewNode(c, v))
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

func (s *Solver) GetStart() *aoc.Node {
	for _, node := range s.Nodes() {
		if node.Value() == Start {
			return node
		}
	}

	return nil
}

func (s *Solver) GetEnd() *aoc.Node {
	for _, node := range s.Nodes() {
		if node.Value() == End {
			return node
		}
	}

	return nil
}

func (s *Solver) BestScore() (int, [][]aoc.Coordinate) {
	minscore := math.MaxInt

	type key struct {
		aoc.Coordinate
		bearing int
	}

	bestscores := make(map[key]int)
	bestfroms := make(map[key][]key)

	type visit struct {
		key
		from  key
		score int
	}

	tovisit := []visit{}

	addvisit := func(k key, score int) {
		tovisit = append(tovisit, visit{key{k.Coordinate, (k.bearing + 1) % 4}, k, score + 1000})
		tovisit = append(tovisit, visit{key{k.Coordinate, (k.bearing + 3) % 4}, k, score + 1000})

		cc := k.Coordinate
		switch k.bearing {
		case North:
			cc.Y--
		case South:
			cc.Y++
		case East:
			cc.X++
		case West:
			cc.X--
		}

		tovisit = append(tovisit, visit{key{cc, k.bearing}, k, score + 1})
	}

	k := key{s.GetStart().Coordinate, East}
	addvisit(k, 0)
	bestscores[k] = 0
	bestfroms[k] = []key{}

	for {
		i := 0
		for j, v := range tovisit {
			if v.score < tovisit[i].score {
				i = j
			}
		}

		cur := tovisit[i]
		tovisit = append(tovisit[:i], tovisit[i+1:]...)

		if cur.score > minscore {
			break
		}

		c := cur.Coordinate
		node := s.Get(c)
		if node == nil || node.Value() == Wall {
			continue
		}

		if _, ok := bestscores[cur.key]; !ok {
			bestscores[cur.key] = cur.score
		}

		if cur.score > bestscores[cur.key] {
			continue
		}

		if _, ok := bestfroms[cur.key]; !ok {
			bestfroms[cur.key] = []key{}
		}
		bestfroms[cur.key] = append(bestfroms[cur.key], cur.from)

		if node.Value() == End {
			minscore = cur.score
			continue
		}

		if len(bestfroms[cur.key]) > 1 {
			continue
		}

		addvisit(cur.key, cur.score)
	}

	end := s.GetEnd()

	var pathsto func(k key) [][]aoc.Coordinate
	pathsto = func(k key) [][]aoc.Coordinate {
		if len(bestfroms[k]) == 0 {
			return [][]aoc.Coordinate{{k.Coordinate}}
		}

		paths := [][]aoc.Coordinate{}

		for _, from := range bestfroms[k] {
			for _, path := range pathsto(from) {
				p := []aoc.Coordinate{}
				p = append(p, path...)
				p = append(p, k.Coordinate)
				paths = append(paths, p)
			}
		}

		return paths
	}

	paths := [][]aoc.Coordinate{}

	for i := 0; i < 4; i++ {
		score, ok := bestscores[key{end.Coordinate, i}]
		if !ok || score > minscore {
			continue
		}
		paths = append(paths, pathsto(key{end.Coordinate, i})...)
	}

	return minscore, paths
}

func (s *Solver) Part1() int {
	score, _ := s.BestScore()
	return score
}

func (s *Solver) Part2() int {
	visited := make(map[aoc.Coordinate]struct{})

	_, bps := s.BestScore()
	for _, path := range bps {
		for _, c := range path {
			visited[c] = struct{}{}
		}
	}

	return len(visited)
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
