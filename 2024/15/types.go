package main

import (
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

const (
	Empty int = iota
	Wall
	Box
	Robot
)

type Warehouse struct {
	Robot *Object
	Walls []*Object
	Boxes []*Object
	grid  map[aoc.Coordinate]*Object
}

func ParseWarehouse(s string) *Warehouse {
	var robot *Object
	walls := []*Object{}
	boxes := []*Object{}

	for y, line := range strings.Split(s, "\n") {
		for x, r := range line {
			c := aoc.NewCoordinate(x, y)

			switch r {
			case '#':
				walls = append(walls, &Object{c, 1, Wall})

			case 'O':
				boxes = append(boxes, &Object{c, 1, Box})

			case '@':
				robot = &Object{c, 1, Robot}
			}
		}
	}

	grid := make(map[aoc.Coordinate]*Object)
	populateGrid(grid, []*Object{robot})
	populateGrid(grid, walls)
	populateGrid(grid, boxes)

	w := Warehouse{robot, walls, boxes, grid}
	return &w
}

func (w *Warehouse) Clone() *Warehouse {
	robot := &Object{w.Robot.Coordinate, w.Robot.Width, Robot}
	walls := []*Object{}
	boxes := []*Object{}

	for _, wall := range w.Walls {
		walls = append(walls, &Object{wall.Coordinate, wall.Width, Wall})
	}

	for _, box := range w.Boxes {
		boxes = append(boxes, &Object{box.Coordinate, box.Width, Box})
	}

	grid := make(map[aoc.Coordinate]*Object)
	populateGrid(grid, []*Object{robot})
	populateGrid(grid, walls)
	populateGrid(grid, boxes)

	return &Warehouse{robot, walls, boxes, grid}
}

func (w *Warehouse) Move(obj *Object, d Direction) {
	if !w.CanMove(obj, d) {
		return
	}

	tc := obj.Coordinate
	switch d {
	case Left:
		tc.X--
	case Right:
		tc.X++
	case Up:
		tc.Y--
	case Down:
		tc.Y++
	}

	tgts := make(map[*Object]struct{})
	for i := 0; i < obj.Width; i++ {
		tgt := w.grid[aoc.NewCoordinate(tc.X+i, tc.Y)]
		if tgt == nil || tgt == obj {
			continue
		}

		tgts[tgt] = struct{}{}
	}

	for tgt := range tgts {
		w.Move(tgt, d)
	}

	for i := 0; i < obj.Width; i++ {
		delete(w.grid, aoc.NewCoordinate(obj.X+i, obj.Y))
	}

	for i := 0; i < obj.Width; i++ {
		w.grid[aoc.NewCoordinate(tc.X+i, tc.Y)] = obj
	}

	obj.Coordinate = tc
}

func (w *Warehouse) CanMove(obj *Object, d Direction) bool {
	if obj.Type == Wall {
		return false
	}

	tc := obj.Coordinate
	switch d {
	case Left:
		tc.X--
	case Right:
		tc.X++
	case Up:
		tc.Y--
	case Down:
		tc.Y++
	}

	tgts := make(map[*Object]struct{})

	for i := 0; i < obj.Width; i++ {
		tgt := w.grid[aoc.NewCoordinate(tc.X+i, tc.Y)]
		if tgt == nil || tgt == obj {
			continue
		}

		tgts[tgt] = struct{}{}
	}

	for tgt := range tgts {
		if !w.CanMove(tgt, d) {
			return false
		}
	}

	return true
}

func (w *Warehouse) ExpandWidth(f int) {
	w.Robot.X *= f

	for _, wall := range w.Walls {
		wall.Width *= f
		wall.X *= f
	}

	for _, box := range w.Boxes {
		box.Width *= f
		box.X *= f
	}

	grid := make(map[aoc.Coordinate]*Object)
	populateGrid(grid, []*Object{w.Robot})
	populateGrid(grid, w.Walls)
	populateGrid(grid, w.Boxes)

	w.grid = grid
}

func (w *Warehouse) String() string {
	minX := w.Robot.X
	maxX := w.Robot.X + w.Robot.Width - 1
	minY := w.Robot.Y
	maxY := w.Robot.Y

	for _, wall := range w.Walls {
		if wall.X < minX {
			minX = wall.X
		}
		if wall.X+wall.Width-1 > maxX {
			maxX = wall.X + wall.Width - 1
		}
		if wall.Y < minY {
			minY = wall.Y
		}
		if wall.Y > maxY {
			maxY = wall.Y
		}
	}

	for _, box := range w.Boxes {
		if box.X < minX {
			minX = box.X
		}
		if box.X+box.Width-1 > maxX {
			maxX = box.X + box.Width - 1
		}
		if box.Y < minY {
			minY = box.Y
		}
		if box.Y > maxY {
			maxY = box.Y
		}
	}

	output := ""

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			obj := w.grid[aoc.NewCoordinate(x, y)]
			if obj == nil {
				output += "."
				continue
			}

			switch obj.Type {
			case Wall:
				output += "#"

			case Box:
				output += "O"

			case Robot:
				output += "@"

			}
		}
		output += "\n"
	}

	return output
}

type Object struct {
	aoc.Coordinate
	Width int
	Type  int
}

func populateGrid(grid map[aoc.Coordinate]*Object, objects []*Object) {
	for _, o := range objects {
		for i := 0; i < o.Width; i++ {
			c := aoc.NewCoordinate(o.X+i, o.Y)
			grid[c] = o
		}
	}
}
