package main

import (
	"math"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()
	grid := NewGrid(input)

	p := grid.ShortestPathToAny(grid.End, []*Square{grid.Start})
	println(len(p))

	ls := grid.LowestSquares()
	p = grid.ShortestPathToAny(grid.End, ls)
	println(len(p))
}

type Grid struct {
	this  [][]*Square
	Start *Square
	End   *Square
}

func NewGrid(s string) *Grid {
	lines := strings.Split(s, "\n")

	var start, end *Square
	g := make([][]*Square, len(lines))
	for i := range g {
		g[i] = make([]*Square, len(lines[i]))
		for j, val := range lines[i] {
			switch val {
			case 'S':
				g[i][j] = NewSquare(i, j, 0)
				start = g[i][j]
			case 'E':
				g[i][j] = NewSquare(i, j, int('z'-'a'))
				end = g[i][j]
			default:
				g[i][j] = NewSquare(i, j, int(val-'a'))
			}
		}
	}

	for i := range g {
		for j := range g[i] {
			s := g[i][j]
			if i > 0 {
				s.SetNeighbour(0, g[i-1][j])
			}
			if j > 0 {
				s.SetNeighbour(1, g[i][j-1])
			}
			if i < len(g)-1 {
				s.SetNeighbour(2, g[i+1][j])
			}
			if j < len(g[i])-1 {
				s.SetNeighbour(3, g[i][j+1])
			}
		}
	}

	grid := Grid{
		this:  g,
		Start: start,
		End:   end,
	}

	return &grid
}

func (g *Grid) ShortestPathToAny(start *Square, target []*Square) []*Square {
	visit := map[*Square]bool{start: true}
	paths := map[*Square][]*Square{start: {}}

	for len(visit) > 0 {
		var s *Square
		min := math.MaxInt32

		for square := range visit {
			for _, t := range target {
				if square.Distance(t)+len(paths[square]) < min {
					min = square.Distance(t) + len(paths[square])
					s = square
				}
			}
		}

		for _, n := range s.Neighbours {
			if n == nil || s.Elevation-n.Elevation > 1 {
				continue
			}

			if _, ok := paths[n]; !ok || len(paths[s])+1 < len(paths[n]) {
				visit[n] = true
				paths[n] = append([]*Square{s}, paths[s]...)

				for _, t := range target {
					if n == t {
						return paths[n]
					}
				}
			}
		}

		delete(visit, s)
	}

	return nil
}

func (g *Grid) LowestSquares() []*Square {
	ls := []*Square{}
	for _, row := range g.this {
		for _, s := range row {
			if s.Elevation == 0 {
				ls = append(ls, s)
			}
		}
	}
	return ls
}

type Square struct {
	X, Y       int
	Elevation  int
	Neighbours []*Square
}

func NewSquare(x, y, e int) *Square {
	s := Square{
		X:          x,
		Y:          y,
		Elevation:  e,
		Neighbours: make([]*Square, 4),
	}
	return &s
}

func (s *Square) SetNeighbour(direction int, n *Square) {
	s.Neighbours[direction] = n
}

func (s *Square) Distance(other *Square) int {
	dX := s.X - other.X
	if dX < 0 {
		dX = -dX
	}

	dY := s.Y - other.Y
	if dY < 0 {
		dY = -dY
	}

	return dX + dY
}
