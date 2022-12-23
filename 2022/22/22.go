package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()
	lines := strings.Split(input, "\n")
	path := lines[len(lines)-1]
	m := lines[:len(lines)-2]

	p1 := NewFlatPuzzle(m)
	s1 := NewSolver(p1, path)
	fmt.Println(s1.Password())

	p2 := NewCubePuzzle(m)
	s2 := NewSolver(p2, path)
	fmt.Println(s2.Password())
}

type Solver struct {
	puzzle     Puzzle
	path       string
	coordinate Coordinate
	facing     Direction
}

func NewSolver(puzzle Puzzle, path string) *Solver {
	solver := Solver{
		puzzle:     puzzle,
		path:       path,
		coordinate: puzzle.TopLeft(),
		facing:     RIGHT,
	}
	return &solver
}

func (s *Solver) Password() int {
	s.coordinate = s.puzzle.TopLeft()

	path := s.path
	for len(path) > 0 {
		var i int
		if path[0] == 'R' || path[0] == 'L' {
			i = 1
		} else {
			i = strings.IndexAny(path, "RL")
		}
		if i < 0 {
			i = len(path)
		}
		s.Do(path[:i])
		path = path[i:]
	}

	return 1000*(s.coordinate.Y+1) + 4*(s.coordinate.X+1) + int(s.facing)
}

func (s *Solver) Do(ins string) {
	if ins == "R" {
		s.facing = (s.facing + 1) % 4
		return
	}

	if ins == "L" {
		s.facing = (s.facing + 3) % 4
		return
	}

	steps, err := strconv.Atoi(ins)
	if err != nil {
		panic(err)
	}

	for steps > 0 {
		c, d := s.puzzle.Next(s.coordinate, s.facing)
		if s.puzzle.TileAt(c) == WALL {
			break
		}
		s.coordinate = c
		s.facing = d
		steps--
	}
}

type Puzzle interface {
	TileAt(Coordinate) Tile
	Next(Coordinate, Direction) (Coordinate, Direction)
	TopLeft() Coordinate
}

type FlatPuzzle struct {
	tiles map[Coordinate]Tile
	col   map[int][2]int
	row   map[int][2]int
}

func NewFlatPuzzle(s []string) *FlatPuzzle {
	puzzle := FlatPuzzle{
		tiles: make(map[Coordinate]Tile),
		col:   make(map[int][2]int),
		row:   make(map[int][2]int),
	}
	for y, line := range s {
		for x, c := range line {
			switch c {
			case '.':
				puzzle.tiles[Coordinate{x, y}] = OPEN
			case '#':
				puzzle.tiles[Coordinate{x, y}] = WALL
			default:
				continue
			}

			if _, ok := puzzle.col[x]; !ok {
				puzzle.col[x] = [2]int{math.MaxInt, 0}
			}
			if y < puzzle.col[x][0] {
				puzzle.col[x] = [2]int{y, puzzle.col[x][1]}
			}
			if y > puzzle.col[x][1] {
				puzzle.col[x] = [2]int{puzzle.col[x][0], y}
			}

			if _, ok := puzzle.row[y]; !ok {
				puzzle.row[y] = [2]int{math.MaxInt, 0}
			}
			if x < puzzle.row[y][0] {
				puzzle.row[y] = [2]int{x, puzzle.row[y][1]}
			}
			if x > puzzle.row[y][1] {
				puzzle.row[y] = [2]int{puzzle.row[y][0], x}
			}
		}
	}
	return &puzzle
}

func (puzzle *FlatPuzzle) TileAt(c Coordinate) Tile {
	return puzzle.tiles[c]
}

func (puzzle *FlatPuzzle) TopLeft() Coordinate {
	return Coordinate{puzzle.row[0][0], 0}
}

func (puzzle *FlatPuzzle) Next(c Coordinate, d Direction) (Coordinate, Direction) {
	c.X += directions[d][0]
	c.Y += directions[d][1]

	if _, ok := puzzle.tiles[c]; ok {
		return c, d
	}

	switch d {
	case RIGHT:
		c.X = puzzle.row[c.Y][0]
	case DOWN:
		c.Y = puzzle.col[c.X][0]
	case LEFT:
		c.X = puzzle.row[c.Y][1]
	case UP:
		c.Y = puzzle.col[c.X][1]
	}
	return c, d
}

type CubePuzzle struct {
	tiles   map[Coordinate]Tile
	mapping map[Coordinate]*Square
	squares [6]*Square
	sidelen int
}

func NewCubePuzzle(s []string) *CubePuzzle {
	var squares [6]*Square
	tiles := make(map[Coordinate]Tile)
	mapping := make(map[Coordinate]*Square)

	w, h := 0, len(s)
	for _, line := range s {
		if len(line) > w {
			w = len(line)
		}
	}

	var sidelen int
	switch {
	case h/5 == w/2 && h%5 == 0 && w%2 == 0:
		sidelen = h / 5
	case h/2 == w/5 && h%2 == 0 && w%5 == 0:
		sidelen = h / 2
	case h/3 == w/4 && h%3 == 0 && w%4 == 0:
		sidelen = h / 3
	case h/4 == w/3 && h%4 == 0 && w%3 == 0:
		sidelen = h / 4
	default:
		panic("not a valid cube net")
	}

	i := 0
	for y, line := range s {
		for x, c := range line {
			coordinate := Coordinate{x, y}
			switch c {
			case '.':
				tiles[coordinate] = OPEN
			case '#':
				tiles[coordinate] = WALL
			default:
				continue
			}

			if x%sidelen == 0 && y%sidelen == 0 {
				for squares[i] != nil {
					i++
				}
				s := NewSquare(i, Coordinate{x, y}, sidelen)
				squares[i] = s
				mapping[coordinate] = squares[i]
			} else {
				mapping[coordinate] = mapping[Coordinate{x - (x % sidelen), y - (y % sidelen)}]
			}
		}
	}

	for _, s := range squares {
		if a, ok := mapping[Coordinate{s.edges[RIGHT].p1.X + 1, s.edges[RIGHT].p1.Y}]; ok {
			s.edges[RIGHT].adjacent = a.edges[LEFT]
			a.edges[LEFT].adjacent = s.edges[RIGHT]
		}
		if a, ok := mapping[Coordinate{s.edges[DOWN].p1.X, s.edges[DOWN].p1.Y + 1}]; ok {
			s.edges[DOWN].adjacent = a.edges[UP]
			a.edges[UP].adjacent = s.edges[DOWN]
		}
	}

	queue := []*SquareEdge{}
	for _, s := range squares {
		for _, e := range s.edges {
			if e.adjacent == nil {
				queue = append(queue, e)
			}
		}
	}

	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]

		if e.adjacent != nil {
			continue
		}

		a := e.square.edges[(e.direction+1)%4].adjacent
		if a == nil {
			queue = append(queue, e)
			continue
		}

		a = a.square.edges[(a.direction+1)%4].adjacent
		if a == nil {
			queue = append(queue, e)
			continue
		}

		a = a.square.edges[(a.direction+1)%4]
		e.adjacent = a
		a.adjacent = e
	}

	puzzle := CubePuzzle{
		tiles:   tiles,
		mapping: mapping,
		squares: squares,
		sidelen: sidelen,
	}
	return &puzzle
}

func (puzzle *CubePuzzle) TileAt(c Coordinate) Tile {
	return puzzle.tiles[c]
}

func (puzzle *CubePuzzle) TopLeft() Coordinate {
	c := puzzle.squares[0].edges[0].p1
	for _, s := range puzzle.squares {
		if s.edges[UP].p1.Y == 0 && s.edges[UP].p1.X < c.X {
			c = s.edges[UP].p1
		}
	}
	return c
}

func (puzzle *CubePuzzle) Next(c Coordinate, d Direction) (Coordinate, Direction) {
	s := puzzle.mapping[c]

	c.X += directions[d][0]
	c.Y += directions[d][1]
	if dst, ok := puzzle.mapping[c]; ok && dst == s {
		return c, d
	}

	e := s.edges[d]
	a := e.adjacent

	var n int
	if d == LEFT || d == RIGHT {
		n = c.Y - e.p1.Y
	} else {
		n = c.X - e.p1.X
	}

	if (e.direction/2 != a.direction/2 && e.direction%2 != a.direction%2) || e.direction == a.direction {
		n = puzzle.sidelen - 1 - n
	}

	c = a.NthCoordinate(n)
	d = (a.direction + 2) % 4
	return c, d
}

type Square struct {
	id    int
	edges [4]*SquareEdge
}

func NewSquare(id int, c Coordinate, sidelen int) *Square {
	square := Square{id: id}
	var d Direction
	for d = RIGHT; d <= UP; d++ {
		var p1, p2 Coordinate
		switch d {
		case RIGHT:
			p1 = Coordinate{c.X + sidelen - 1, c.Y}
			p2 = Coordinate{c.X + sidelen - 1, c.Y + sidelen - 1}
		case DOWN:
			p1 = Coordinate{c.X, c.Y + sidelen - 1}
			p2 = Coordinate{c.X + sidelen - 1, c.Y + sidelen - 1}
		case LEFT:
			p1 = c
			p2 = Coordinate{c.X, c.Y + sidelen - 1}
		case UP:
			p1 = c
			p2 = Coordinate{c.X + sidelen - 1, c.Y}
		}
		edge := SquareEdge{
			p1:        p1,
			p2:        p2,
			square:    &square,
			direction: d,
		}
		square.edges[d] = &edge
	}
	return &square
}

type SquareEdge struct {
	p1, p2    Coordinate
	square    *Square
	direction Direction
	adjacent  *SquareEdge
}

func (edge *SquareEdge) NthCoordinate(n int) Coordinate {
	c := edge.p1
	if edge.direction == UP || edge.direction == DOWN {
		c.X += n
	} else {
		c.Y += n
	}
	return c
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

type Coordinate struct {
	X, Y int
}

type Tile bool

const (
	OPEN Tile = true
	WALL Tile = false
)
